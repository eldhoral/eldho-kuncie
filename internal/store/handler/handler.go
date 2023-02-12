package handler

import (
	"net/http"

	"github.com/eldhoral/eldho-kuncie/internal/base/app"

	"github.com/eldhoral/eldho-kuncie/internal/base/handler"
	"github.com/eldhoral/eldho-kuncie/internal/store/service"
	"github.com/eldhoral/eldho-kuncie/pkg/server"
)

// HTTPHandler handles company API methods
type HTTPHandler struct {
	App          *handler.BaseHTTPHandler
	StoreService service.Service
}

// NewHTTPHandler creates new http handler
func NewHTTPHandler(base *handler.BaseHTTPHandler, storeService service.Service) *HTTPHandler {
	return &HTTPHandler{App: base, StoreService: storeService}
}

// Handler Basic Method ======================================================================================================

// AsWebResponse will set httpStatus based on status
func (h HTTPHandler) AsWebResponse(ctx *app.Context, status int, message string, data interface{}) *server.Response {
	if data == nil {
		data = []int{}
	}
	return h.App.AsJson(ctx, status, message, data)
}

// AsInternalError will set httpStatus to 500. Different with AsMobileJson always 200 code
func (h HTTPHandler) AsInternalError(ctx *app.Context, err error, message string) *server.Response {
	return h.App.AsJson(ctx, http.StatusInternalServerError, message, nil)
}
