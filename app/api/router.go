package api

import (
	"fmt"
	"net/http"

	"github.com/eldhoral/eldho-kuncie/internal/base/handler"
)

func (h *HttpServe) setupRouter() {
	// StrictSlash will treat /projects/ to be same as /projects
	h.v1 = h.router.PathPrefix("/api/v1/").Subrouter().StrictSlash(true)

	h.Route("GET", "/product/list", h.store.ListProducthandler)
	h.Route("POST", "/product/purchase", h.store.PurchaseProduct)
	h.Route("POST", "/checkout/create", h.store.CreateCheckoutHandler)
	h.Route("POST", "/checkout/add", h.store.AddProductToCheckoutHandler)
	h.Route("GET", "/checkout/list", h.store.ListCheckoutHandler)
	h.Route("POST", "/checkout/detail", h.store.GetCheckoutDetail)

	// assign method not allowed handler
	h.v1.MethodNotAllowedHandler = h.base.MethodNotAllowedHandler()
}

func (h *HttpServe) Route(method string, path string, f handler.HandlerFn) {
	if method != http.MethodGet &&
		method != http.MethodPost &&
		method != http.MethodDelete &&
		method != http.MethodPut {
		panic(fmt.Sprintf(":%s method not allow", method))
	}

	h.v1.HandleFunc(path, h.base.RunAction(f)).Methods(method)
}
