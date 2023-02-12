package repository

import (
	"testing"

	"github.com/eldhoral/eldho-kuncie/pkg/data"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func Test_GetProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "sku", "name", "price", "quantity"}).
			AddRow(int64(1), "120P90", "Google Home", 49.99, 10)

		query := `SELECT id, sku, name, price, quantity FROM product ORDER BY id ASC`

		mock.ExpectQuery("[" + query + "]").WillReturnRows(rows)
		repo := NewRepository(sqlxDB)

		result, err := repo.SelectProducts()
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func Test_CreateCheckout(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	params := data.NewParamsWrapper()
	params.Add("product_id", 1)
	params.Add("quantity", 2)
	uniquePurchaseId := uuid.NewV1().String()
	params.Add("purchase_id", uniquePurchaseId)

	t.Run("success", func(t *testing.T) {
		query := `INSERT INTO checkout
		SET product_id = :product_id, quantity = :quantity, purchase_id = :purchase_id`

		mock.ExpectExec("[" + query + "]").WillReturnResult(sqlmock.NewResult(1, 1))
		repo := NewRepository(sqlxDB)

		err := repo.CreateCheckout(params)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectQuery("INSERT").WillReturnError(err)
		repo := NewRepository(sqlxDB)

		err := repo.CreateCheckout(params)
		assert.Error(t, err)
	})
}
