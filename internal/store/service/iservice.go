package service

import (
	store "github.com/eldhoral/eldho-kuncie/internal/store/domain/product"
	"github.com/eldhoral/eldho-kuncie/internal/store/domain/rules"
	"github.com/eldhoral/eldho-kuncie/pkg/data"
)

type Service interface {
	ListProduct() (httpStatus int, result *[]store.Product, err error)
	CreateCheckout(params data.Params) (httpStatus int, err error)
	AddProductToCheckout(params data.Params) (httpStatus int, err error)
	PurchaseProduct(params data.Params) (httpStatus int, purchasement store.ProductTotalPurchase, err error)
	DiscountRules(productPurchasetotal []store.ProductPurchase, criteriaRules *rules.CriteriasRules) (discountMapping []store.ProductPurchase)
	ListCheckout() (httpStatus int, result *[]store.Checkout, err error)
	CheckoutDetail(params data.Params) (httpStatus int, purchasement store.ProductTotalPurchase, err error)
}
