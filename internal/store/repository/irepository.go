package repository

import (
	store "github.com/eldhoral/eldho-kuncie/internal/store/domain/product"
	"github.com/eldhoral/eldho-kuncie/internal/store/domain/rules"
	"github.com/eldhoral/eldho-kuncie/pkg/data"
)

type Repository interface {
	SelectProducts() (result *[]store.Product, err error)
	GetProduct(productId int64) (result *store.Product, err error)
	CheckQuantityByProductIds(productIds []int64) (result *[]store.ProductQuantity, err error)

	GetCheckoutByPurchaseid(purchaseId string) (result *[]store.ProductCheckout, err error)
	CreateCheckout(params data.Params) (err error)
	AddNewProductToCheckout(params data.Params) (err error)
	UpdateQuantityProductInCheckout(params data.Params) (rowsAffected int64, err error)
	SelectCheckout() (result []string, err error)
	PurchaseProductByCheckout(params data.Params) (rowsAffected int64, err error)

	CheckExistingDiscountRules(productIds []int64) (result *[]rules.DiscountRules, err error)
	GetCriteriasRules(rulesId int64) (result *rules.CriteriasRules, err error)

	UpdateProductByCheckout(params data.Params) (rowsAffected int64, err error)
}
