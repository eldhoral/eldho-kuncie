package service

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	store "github.com/eldhoral/eldho-kuncie/internal/store/domain/product"
	"github.com/eldhoral/eldho-kuncie/internal/store/domain/rules"
	"github.com/eldhoral/eldho-kuncie/internal/store/repository"
	"github.com/eldhoral/eldho-kuncie/pkg/data"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
)

// NewService creates new user service
func NewService(repo repository.Repository) Service {
	return &service{
		storeRepo: repo,
	}
}

type service struct {
	storeRepo repository.Repository
}

func (s service) ListProduct() (httpStatus int, result *[]store.Product, err error) {
	repo, err := s.storeRepo.SelectProducts()
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("No product was found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}

func (s service) CreateCheckout(params data.Params) (httpStatus int, err error) {
	// Check if product quantity in db is enough for the checkput's product quantity
	productIds := []int64{params.GetInt64("product_id")}
	productQuantity, err := s.storeRepo.CheckQuantityByProductIds(productIds)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, errors.New("No product was found")
	}
	if err != nil {
		fmt.Println(err)
		return http.StatusInternalServerError, errors.New("Error checking product quantity")
	}

	for _, dataProduct := range *productQuantity {
		if dataProduct.Quantity < params.GetInt("quantity") {
			return http.StatusInternalServerError, errors.New("Product quantity is less then checkput's product quantity")
		}
	}

	// Generate new unique UUID based on timestamp and mac address
	// everytime new checkout is want to be created
	uniquePurchaseId := uuid.NewV1().String()
	params.Add("purchase_id", uniquePurchaseId)

	err = s.storeRepo.CreateCheckout(params)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (s service) AddProductToCheckout(params data.Params) (httpStatus int, err error) {
	// Check if product quantity in db is enough for the checkput's product quantity
	productIds := []int64{params.GetInt64("product_id")}
	productQuantity, err := s.storeRepo.CheckQuantityByProductIds(productIds)

	for _, dataProduct := range *productQuantity {
		if dataProduct.Quantity < params.GetInt("quantity") {
			return http.StatusInternalServerError, errors.New("Product quantity is less then checkput's product quantity")
		}
	}

	// Check if the checkout is existing
	product, err := s.storeRepo.GetCheckoutByPurchaseid(params.GetString("purchase_id"))
	if err == sql.ErrNoRows {
		return http.StatusNotFound, errors.New("No checkout was found")
	}
	if err != nil {
		return http.StatusInternalServerError, errors.New("Error when finding checkout")
	}

	// Check if product is existing in the checkout
	for _, dataProduct := range *product {
		// If exist then only update te quantity of the product itself
		if dataProduct.ProductId == params.GetInt64("product_id") {
			addedQuantity := params.GetInt64("quantity") + int64(dataProduct.Quantity)
			params.Add("quantity", addedQuantity)
			_, err := s.storeRepo.UpdateProductByCheckout(params)
			if err != nil {
				return http.StatusInternalServerError, errors.New("Error when adding new product to checkout")
			}

			return http.StatusCreated, nil

		}
	}

	// If not exist, then add new product and the quantity of the product
	err = s.storeRepo.AddNewProductToCheckout(params)
	if err != nil {
		return http.StatusInternalServerError, errors.New("Error when adding new product to checkout")
	}

	return http.StatusCreated, nil
}

func (s service) PurchaseProduct(params data.Params) (httpStatus int, purchasement store.ProductTotalPurchase, err error) {
	// Check if the checkout is existing
	productCheckout, err := s.storeRepo.GetCheckoutByPurchaseid(params.GetString("purchase_id"))
	if err == sql.ErrNoRows {
		return http.StatusNotFound, purchasement, errors.New("No checkout was found")
	}
	if err != nil {
		return http.StatusInternalServerError, purchasement, err
	}

	productIds := []int64{}
	productPurchasetotal := make([]store.ProductPurchase, 0)

	for _, dataProduct := range *productCheckout {
		productIds = append(productIds, dataProduct.ProductId)
	}

	// Check if the product quantity is enought to purchase
	productQuantity, err := s.storeRepo.CheckQuantityByProductIds(productIds)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, purchasement, errors.New("No product was found")
	}
	if err != nil {
		return http.StatusInternalServerError, purchasement, err
	}

	// Check if the product quantity is enought to purchase
	for i, dataProduct := range *productCheckout {
		// return err if product quantity is less then checkout quantity
		if dataProduct.Quantity > (*productQuantity)[i].Quantity {
			return http.StatusInternalServerError, purchasement, errors.New("Quantity for product_id : " + cast.ToString(dataProduct.ProductId) + " is less then quantity in checkout")
		}

		// Check if the product is existing
		product, err := s.storeRepo.GetProduct(dataProduct.ProductId)
		if err == sql.ErrNoRows {
			return http.StatusNotFound, purchasement, errors.New("No product was found with product_id : " + cast.ToString(dataProduct.ProductId))
		}
		if err != nil {
			return http.StatusInternalServerError, purchasement, err
		}

		productIds = append(productIds, dataProduct.ProductId)
		totalPrice := product.Price.Mul(decimal.NewFromInt(int64(dataProduct.Quantity)))
		productPurchase := store.ProductPurchase{
			Id:         dataProduct.Id,
			ProductId:  dataProduct.ProductId,
			Quantity:   dataProduct.Quantity,
			PurchaseId: dataProduct.PurchaseId,
			TotalPrice: totalPrice,
		}

		productPurchasetotal = append(productPurchasetotal, productPurchase)

		params.Add("product_id", dataProduct.Id)
		params.Add("quantity", product.Quantity-dataProduct.Quantity)
		rowsAffected, err := s.storeRepo.PurchaseProductByCheckout(params)
		if rowsAffected == 0 {
			return http.StatusNotFound, purchasement, errors.New("No changes was made")
		}
		if err != nil {
			return http.StatusInternalServerError, purchasement, err
		}
	}

	// Check if there is discount rules
	discountRules, err := s.storeRepo.CheckExistingDiscountRules(productIds)
	for _, dataDiscountRules := range *discountRules {
		// Get criteria rules for each product
		criteriasRules, _ := s.storeRepo.GetCriteriasRules(dataDiscountRules.Rules)
		// Mapping discount product
		productPurchasetotal = s.DiscountRules(productPurchasetotal, criteriasRules)

	}

	totalPrice := decimal.NewFromInt(int64(0))
	for _, dataPurchase := range productPurchasetotal {
		totalPrice = totalPrice.Add(dataPurchase.TotalPrice)
	}

	purchasement = store.ProductTotalPurchase{
		ProductPurchase:    productPurchasetotal,
		TotalPurchasePrice: "$" + totalPrice.String(),
	}

	return http.StatusCreated, purchasement, nil
}

func (s service) DiscountRules(productPurchasetotal []store.ProductPurchase, criteriaRules *rules.CriteriasRules) (discountMapping []store.ProductPurchase) {
	findLastIdPurchase := len(productPurchasetotal)
	for i, dataproductPurchase := range productPurchasetotal {
		if criteriaRules.Id == 1 {
			// buy 1, get 1 selected product
			if dataproductPurchase.Quantity >= int(criteriaRules.Criteria) {
				productPurchasetotal = append(productPurchasetotal, store.ProductPurchase{
					Id:         int64(findLastIdPurchase + 1),
					ProductId:  criteriaRules.Reward,
					PurchaseId: dataproductPurchase.PurchaseId,
					Quantity:   dataproductPurchase.Quantity,
					TotalPrice: decimal.NewFromInt(0),
				})
			}

		}

		if criteriaRules.Id == 2 {
			// buy 3, pay 2
			if dataproductPurchase.Quantity == int(criteriaRules.Criteria) {
				totalPrice := productPurchasetotal[i].TotalPrice.Div(decimal.NewFromInt(int64(criteriaRules.Criteria)))
				productPurchasetotal[i].TotalPrice = totalPrice.Mul(decimal.NewFromInt(int64(criteriaRules.Reward)))
			}
		}

		if criteriaRules.Id == 3 {
			// buy >3, get 10% discount on price
			if dataproductPurchase.Quantity > int(criteriaRules.Criteria) {
				totalPrice := productPurchasetotal[i].TotalPrice.Div(decimal.NewFromInt(int64(dataproductPurchase.Quantity)))
				totalPrice = totalPrice.Mul(decimal.NewFromFloat(float64(criteriaRules.Reward) / 100))
				totalPrice = totalPrice.Mul(decimal.NewFromInt(int64(dataproductPurchase.Quantity)))
				productPurchasetotal[i].TotalPrice = totalPrice
			}
		}
	}

	return productPurchasetotal
}

func (s service) CheckoutDetail(params data.Params) (httpStatus int, purchasement store.ProductTotalPurchase, err error) {
	// Check if the checkout is existing
	productCheckout, err := s.storeRepo.GetCheckoutByPurchaseid(params.GetString("purchase_id"))
	if err == sql.ErrNoRows {
		return http.StatusNotFound, purchasement, errors.New("No checkout was found")
	}
	if err != nil {
		return http.StatusInternalServerError, purchasement, err
	}

	productIds := []int64{}
	productPurchasetotal := make([]store.ProductPurchase, 0)

	for _, dataProduct := range *productCheckout {
		productIds = append(productIds, dataProduct.ProductId)
	}

	// Check if the product quantity is enought to purchase
	productQuantity, err := s.storeRepo.CheckQuantityByProductIds(productIds)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, purchasement, errors.New("No product was found")
	}
	if err != nil {
		return http.StatusInternalServerError, purchasement, err
	}

	// Check if the product quantity is enought to purchase
	for i, dataProduct := range *productCheckout {
		// return err if product quantity is less then checkout quantity
		if dataProduct.Quantity > (*productQuantity)[i].Quantity {
			return http.StatusInternalServerError, purchasement, errors.New("Quantity for product_id : " + cast.ToString(dataProduct.ProductId) + " is less then quantity in checkout")
		}

		// Check if the product is existing
		product, err := s.storeRepo.GetProduct(dataProduct.ProductId)
		if err == sql.ErrNoRows {
			return http.StatusNotFound, purchasement, errors.New("No product was found with product_id : " + cast.ToString(dataProduct.ProductId))
		}
		if err != nil {
			return http.StatusInternalServerError, purchasement, err
		}

		productIds = append(productIds, dataProduct.ProductId)
		totalPrice := product.Price.Mul(decimal.NewFromInt(int64(dataProduct.Quantity)))
		productPurchase := store.ProductPurchase{
			Id:         dataProduct.Id,
			ProductId:  dataProduct.ProductId,
			Quantity:   dataProduct.Quantity,
			PurchaseId: dataProduct.PurchaseId,
			TotalPrice: totalPrice,
		}

		productPurchasetotal = append(productPurchasetotal, productPurchase)
	}

	// Check if there is discount rules
	discountRules, err := s.storeRepo.CheckExistingDiscountRules(productIds)
	for _, dataDiscountRules := range *discountRules {
		// Get criteria rules for each product
		criteriasRules, _ := s.storeRepo.GetCriteriasRules(dataDiscountRules.Rules)
		// Mapping discount product
		productPurchasetotal = s.DiscountRules(productPurchasetotal, criteriasRules)

	}

	totalPrice := decimal.NewFromInt(int64(0))
	for _, dataPurchase := range productPurchasetotal {
		totalPrice = totalPrice.Add(dataPurchase.TotalPrice)
	}

	purchasement = store.ProductTotalPurchase{
		ProductPurchase:    productPurchasetotal,
		TotalPurchasePrice: "$" + totalPrice.String(),
	}

	return http.StatusCreated, purchasement, nil
}

func (s service) ListCheckout() (httpStatus int, result *[]store.Checkout, err error) {
	repo, err := s.storeRepo.SelectCheckout()
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("No checkout was found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	checkout := make([]store.Checkout, len(repo))

	for i, dataCheckout := range repo {
		checkout[i].PurchaseId = dataCheckout
	}
	return http.StatusOK, &checkout, nil
}
