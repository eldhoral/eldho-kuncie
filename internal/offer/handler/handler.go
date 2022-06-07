package handler

import (
	"net/http"

	"bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/base/app"

	"bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/base/handler"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/offer/service"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/server"
)

//HTTPHandler handles company API methods
type HTTPHandler struct {
	App          *handler.BaseHTTPHandler
	OfferService service.Service
}

//NewHTTPHandler creates new http handler
func NewHTTPHandler(base *handler.BaseHTTPHandler, ofService service.Service) *HTTPHandler {
	return &HTTPHandler{App: base, OfferService: ofService}
}

// Handler Basic Method ======================================================================================================

func (h HTTPHandler) Test(ctx *app.Context) *server.Response {
	return h.AsMobileJson(ctx, http.StatusOK, "Success", "Test")
}

// AsWebResponse will set httpStatus based on status
func (h HTTPHandler) AsWebResponse(ctx *app.Context, status int, message string, data interface{}) *server.Response {
	if data == nil {
		data = []int{}
	}
	return h.App.AsJson(ctx, status, message, data)
}

// AsMobileJson always return httpStatus: 200, but Status field: 500,400,200...
func (h HTTPHandler) AsMobileJson(ctx *app.Context, status int, message string, data interface{}) *server.Response {
	return h.App.AsMobileStatusOK(ctx, status, message, data)
}

// Backward compatibility with Yii2, not handle exception
func (h HTTPHandler) ThrowBadRequestException(ctx *app.Context, message string) *server.Response {
	return h.App.ThrowExceptionJson(ctx, http.StatusBadRequest, 0, "Bad Request", message)
}

// MobileBadRequest. For mobile, httpStatus:200, but Status field: http.MobileBadRequest
func (h HTTPHandler) MobileBadRequest(ctx *app.Context, err error) *server.Response {
	return h.App.AsJson(ctx, http.StatusBadRequest, err.Error(), nil)
}

// MobileForbiddenRequest. For mobile, httpStatus:200, but Status field: http.StatusForbidden
func (h HTTPHandler) MobileForbiddenRequest(ctx *app.Context, err error) *server.Response {
	return h.App.AsJson(ctx, http.StatusForbidden, err.Error(), nil)
}

// AsInternalError will set httpStatus to 500. Different with AsMobileJson always 200 code
func (h HTTPHandler) AsInternalError(ctx *app.Context, err error, message string) *server.Response {
	return h.App.AsJson(ctx, http.StatusInternalServerError, message, nil)
}
