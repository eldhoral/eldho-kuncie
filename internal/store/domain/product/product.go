package store

import "github.com/shopspring/decimal"

type Product struct {
	Id       int64           `json:"id" db:"id"`
	Sku      string          `json:"sku" db:"sku"`
	Name     string          `json:"name" db:"name"`
	Price    decimal.Decimal `json:"price" db:"price"`
	Quantity int             `json:"quantity" db:"quantity"`
}

type ProductQuantity struct {
	Id       int64 `json:"id" db:"id"`
	Quantity int   `json:"quantity" db:"quantity"`
}

type ProductCheckout struct {
	Id         int64  `json:"id" db:"id"`
	ProductId  int64  `json:"product_id" db:"product_id"`
	PurchaseId string `json:"purchase_id" db:"purchase_id"`
	Quantity   int    `json:"quantity" db:"quantity"`
}

type ProductPurchase struct {
	Id         int64           `json:"id" db:"id"`
	ProductId  int64           `json:"product_id" db:"product_id"`
	PurchaseId string          `json:"purchase_id" db:"purchase_id"`
	Quantity   int             `json:"quantity" db:"quantity"`
	TotalPrice decimal.Decimal `json:"total_price" db:"total_price"`
}

type ProductTotalPurchase struct {
	ProductPurchase    []ProductPurchase `json:"product_purchase"`
	TotalPurchasePrice string            `json:"total_purchase_price"`
}

type Checkout struct {
	PurchaseId string `json:"purchase_id" db:"purchase_id"`
}
