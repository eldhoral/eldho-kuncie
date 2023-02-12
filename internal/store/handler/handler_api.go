package handler

import (
	"github.com/eldhoral/eldho-kuncie/internal/base/app"
	"github.com/eldhoral/eldho-kuncie/pkg/data"
	"github.com/eldhoral/eldho-kuncie/pkg/server"
)

func (h HTTPHandler) ListProducthandler(ctx *app.Context) *server.Response {
	httpStatus, service, err := h.StoreService.ListProduct()
	if err != nil {
		return h.AsWebResponse(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsWebResponse(ctx, httpStatus, "Get List Product Successfully", service)
}

func (h HTTPHandler) CreateCheckoutHandler(ctx *app.Context) *server.Response {
	jsonBody := ctx.GetJsonBody()

	params := data.NewParamsWrapper()
	params.Add("product_id", jsonBody["product_id"])
	params.Add("quantity", jsonBody["quantity"])

	httpStatus, err := h.StoreService.CreateCheckout(params)
	if err != nil {
		return h.AsWebResponse(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsWebResponse(ctx, httpStatus, "Create Checkout Successfully", nil)
}

func (h HTTPHandler) AddProductToCheckoutHandler(ctx *app.Context) *server.Response {
	jsonBody := ctx.GetJsonBody()

	params := data.NewParamsWrapper()
	params.Add("purchase_id", jsonBody["purchase_id"])
	params.Add("product_id", jsonBody["product_id"])
	params.Add("quantity", jsonBody["quantity"])

	httpStatus, err := h.StoreService.AddProductToCheckout(params)
	if err != nil {
		return h.AsWebResponse(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsWebResponse(ctx, httpStatus, "Add product to checkout successfully", err)
}

func (h HTTPHandler) PurchaseProduct(ctx *app.Context) *server.Response {
	jsonBody := ctx.GetJsonBody()

	params := data.NewParamsWrapper()
	params.Add("purchase_id", jsonBody["purchase_id"])

	httpStatus, service, err := h.StoreService.PurchaseProduct(params)
	if err != nil {
		return h.AsWebResponse(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsWebResponse(ctx, httpStatus, "Purchase Successfully", service)
}

func (h HTTPHandler) ListCheckoutHandler(ctx *app.Context) *server.Response {
	httpStatus, service, err := h.StoreService.ListCheckout()
	if err != nil {
		return h.AsWebResponse(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsWebResponse(ctx, httpStatus, "Get List Checkout Successfully", service)
}

func (h HTTPHandler) GetCheckoutDetail(ctx *app.Context) *server.Response {
	jsonBody := ctx.GetJsonBody()

	params := data.NewParamsWrapper()
	params.Add("purchase_id", jsonBody["purchase_id"])

	httpStatus, service, err := h.StoreService.CheckoutDetail(params)
	if err != nil {
		return h.AsWebResponse(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsWebResponse(ctx, httpStatus, "Get Checkout Detail Successfully", service)
}
