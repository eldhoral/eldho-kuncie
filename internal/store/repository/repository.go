package repository

import (
	"log"

	store "github.com/eldhoral/eldho-kuncie/internal/store/domain/product"
	"github.com/eldhoral/eldho-kuncie/internal/store/domain/rules"
	"github.com/eldhoral/eldho-kuncie/pkg/data"
	"github.com/eldhoral/eldho-kuncie/pkg/helper/slicehelper"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

// NewRepository creates new repository
func NewRepository(db *sqlx.DB) Repository {
	return &repo{db: db}
}

type repo struct {
	db *sqlx.DB
}

func (r repo) SelectProducts() (result *[]store.Product, err error) {
	product := []store.Product{}
	err = r.db.Select(&product, "SELECT id, sku, name, price, quantity FROM product ORDER BY id ASC")
	return &product, err
}

func (r repo) GetProduct(productId int64) (result *store.Product, err error) {
	product := store.Product{}
	err = r.db.Get(&product, "SELECT id, sku, name, price, quantity FROM product WHERE id = ? ORDER BY id ASC", productId)
	return &product, err
}

func (r repo) CheckQuantityByProductIds(productIds []int64) (result *[]store.ProductQuantity, err error) {
	productQuantity := []store.ProductQuantity{}
	query, args, err := sqlx.In("SELECT id, quantity FROM product WHERE id IN (?) ORDER BY id ASC", productIds)
	if err != nil {
		log.Fatal(err)
	}
	query = r.db.Rebind(query)
	err = r.db.Select(&productQuantity, query, args...)
	return &productQuantity, err
}

func (r repo) UpdateProductByCheckout(params data.Params) (rowsAffected int64, err error) {
	arg := map[string]interface{}{
		"product_id":  params.GetValue("product_id"),
		"purchase_id": params.GetValue("purchase_id"),
		"quantity":    params.GetValue("quantity"),
	}
	query := `UPDATE checkout
	SET quantity = :quantity WHERE purchase_id = :purchase_id AND product_id = :product_id`

	result, err := r.db.NamedExec(query, arg)
	if err != nil {
		return
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return
}

func (r repo) PurchaseProductByCheckout(params data.Params) (rowsAffected int64, err error) {
	// start transaction
	tx, err := r.db.Beginx()
	defer func() {
		if err == nil {
			err = tx.Commit()
		} else {
			logrus.Info("Error when deleting temporary data, rollback. Error : " + err.Error())
			err = tx.Rollback()
		}
	}()

	arg := map[string]interface{}{
		"product_id": params.GetValue("product_id"),
		"quantity":   params.GetValue("quantity"),
	}
	query := `UPDATE product
	SET quantity = :quantity WHERE id = :product_id`

	result, err := r.db.NamedExec(query, arg)
	if err != nil {
		return
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	logrus.Info("Purchased product_id :" + params.GetString("product_id") + ", continue deleting purchase_id :" + params.GetString("purchase_id"))
	query = `DELETE FROM checkout where purchase_id = ? AND product_id = ?`
	result, err = tx.Exec(query, params.GetString("purchase_id"), params.GetInt64("product_id"))

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return
}

func (r repo) GetCheckoutByPurchaseid(purchaseId string) (result *[]store.ProductCheckout, err error) {
	productCheckout := []store.ProductCheckout{}
	err = r.db.Select(&productCheckout, `SELECT id, product_id, purchase_id, quantity FROM checkout WHERE purchase_id = ?`, purchaseId)
	return &productCheckout, err
}

func (r repo) CreateCheckout(params data.Params) (err error) {
	arg := map[string]interface{}{
		"product_id":  params.GetValue("product_id"),
		"purchase_id": params.GetValue("purchase_id"),
		"quantity":    params.GetValue("quantity"),
	}
	query := `INSERT INTO checkout
	SET product_id = :product_id, quantity = :quantity, purchase_id = :purchase_id`

	_, err = r.db.NamedExec(query, arg)
	if err != nil {
		return
	}

	return
}

func (r repo) AddNewProductToCheckout(params data.Params) (err error) {
	arg := map[string]interface{}{
		"product_id":  params.GetValue("product_id"),
		"purchase_id": params.GetValue("purchase_id"),
		"quantity":    params.GetValue("quantity"),
	}

	query := `INSERT INTO checkout
	SET product_id = :product_id, quantity = :quantity, purchase_id = :purchase_id`

	_, err = r.db.NamedExec(query, arg)
	if err != nil {
		return
	}
	return
}

func (r repo) SelectCheckout() (result []string, err error) {
	checkout := []store.Checkout{}
	err = r.db.Select(&checkout, "SELECT purchase_id FROM checkout ORDER BY id ASC")
	for _, dataCheckout := range checkout {
		result = append(result, dataCheckout.PurchaseId)
	}
	return slicehelper.UniqueSlices(result), err
}

func (r repo) UpdateQuantityProductInCheckout(params data.Params) (rowsAffected int64, err error) {
	arg := map[string]interface{}{
		"product_id":  params.GetValue("product_id"),
		"purchase_id": params.GetValue("purchase_id"),
		"quantity":    params.GetValue("quantity"),
	}
	query := `UPDATE checkout
	SET product_id = :product_id, quantity = :quantity WHERE purchase_id = ?`

	result, err := r.db.NamedExec(query, arg)
	if err != nil {
		return
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return
}

func (r repo) CheckExistingDiscountRules(productIds []int64) (result *[]rules.DiscountRules, err error) {
	discountRules := []rules.DiscountRules{}
	query, args, err := sqlx.In("SELECT id, product_id, rules FROM discount_rules WHERE product_id IN (?) ORDER BY id ASC", productIds)
	if err != nil {
		log.Fatal(err)
	}
	query = r.db.Rebind(query)
	err = r.db.Select(&discountRules, query, args...)
	return &discountRules, err
}

func (r repo) GetCriteriasRules(rulesId int64) (result *rules.CriteriasRules, err error) {
	criteriasRules := rules.CriteriasRules{}
	err = r.db.Get(&criteriasRules, "SELECT id, criteria, reward FROM criterias_rules WHERE id = ? ORDER BY id ASC", rulesId)
	return &criteriasRules, err
}
