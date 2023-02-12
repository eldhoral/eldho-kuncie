package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	"github.com/eldhoral/eldho-kuncie/internal/base/app"
	storeSer "github.com/eldhoral/eldho-kuncie/internal/store/service"
	"github.com/eldhoral/eldho-kuncie/pkg/httpclient"
	"github.com/eldhoral/eldho-kuncie/pkg/server"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type HandlerFn func(*app.Context) *server.Response

type BaseHTTPHandler struct {
	Handlers   interface{}
	DB         *sqlx.DB
	HTTPClient httpclient.Client
	Params     map[string]string

	StoreService storeSer.Service
}

func NewBaseHTTPHandler(
	db *sqlx.DB,
	httpClient httpclient.Client,
	params map[string]string,

	storeService storeSer.Service,

) *BaseHTTPHandler {

	return &BaseHTTPHandler{DB: db, HTTPClient: httpClient, Params: params,
		StoreService: storeService,
	}
}

// AsJson to response custom message: 200, 201 with message (Mobile use 500 error)
func (h BaseHTTPHandler) AsJson(ctx *app.Context, status int, message string, data interface{}) *server.Response {
	return &server.Response{
		Status:       status,
		Message:      message,
		Data:         data,
		Version:      os.Getenv("APP_VERSION"),
		ResponseType: server.DefaultResponseType,
	}
}

// AsJsonWithLog for custom log
func (h BaseHTTPHandler) AsJsonWithLog(status int, message string, data interface{}, log *server.LogMessage) *server.Response {
	return &server.Response{Status: status, Message: message, Data: data, Version: os.Getenv("APP_VERSION"), Log: log}
}

func (h BaseHTTPHandler) IsStaging() bool {
	return h.Params["APP_ENV"] == "development"
}

func (h BaseHTTPHandler) IsProd() bool {
	return h.Params["APP_ENV"] == "production"
}

func (h BaseHTTPHandler) GetParam(key string) string {
	return h.Params[key]
}

// RunAction entry point to handle route.
func (h BaseHTTPHandler) RunAction(fn HandlerFn) http.HandlerFunc {
	return h.CapturePanic(h.Execute(fn))
}

// Execute SpecificHandler.Method(ctx *app.Context)
// Auth(r) -> Handler(ctx) -> Send Log ULMS (if any) -> JsonResponse
func (f BaseHTTPHandler) Execute(handler HandlerFn) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		// 1. Authentication
		ctx, _ := f.Authentication(rw, r)

		// 2. Capture handler error to avoid infinite loop SendFlock
		defer func() {
			if err0 := recover(); err0 != nil {
				fmt.Errorf(" :boom: :boom: :boom: Panic: %v", err0)

				WriteJSON(rw, http.StatusInternalServerError,
					"Request is halted unexpectedly, please contact the administrator.")
			}
		}()

		// 3. Process route action, and return *server.Response
		resp := handler(ctx)
		httpStatus := resp.GetStatus()

		// Except mobile with server.MobileStatusOKType always response httpStatus 200
		if httpStatus >= http.StatusInternalServerError && resp.ResponseType != server.MobileStatusOKType {
			// Send more clue for other internal app can debug.
			WriteJSON(rw, httpStatus, server.Yii2HTTPExceptionResponse{
				Name:    "Server Error",
				Message: resp.Message,
				Code:    0,
				Status:  httpStatus,
			})
			return
		}

		if httpStatus >= 300 {
			span, _ := tracer.StartSpanFromContext(ctx.Context(), "response", tracer.ResourceName(ctx.Request.RequestURI))
			defer span.Finish(tracer.WithError(fmt.Errorf("%v %v", resp.Message, resp.Data)))
		}

		if f.IsStaging() {
			fmt.Printf("INFO: %s   - code: %d\n\n", ctx.GetRequestInfo(), httpStatus)
		}

		// 5. Response JSON data for web route, and mobile route
		if ctx.IsWebRoute() {
			if resp.ResponseType == server.Yii2ExceptionResponseType {
				WriteJSON(rw, httpStatus, resp.GetYii2Exception())
			} else if resp.ResponseType == server.StreamResponseType {
				WriteStream(rw, resp.Message, resp.GetStream())
			} else {
				WriteJSON(rw, httpStatus, server.WebResponse{
					Status:  httpStatus,
					Message: resp.Message,
					Data:    resp.Data})
			}
			return
		} else {
			if resp.ResponseType == server.Yii2ExceptionResponseType {
				// CodeResponseRequest: happen due to Throw Exception without handle. Or in catch().
				// Often has Code field
				WriteJSON(rw, httpStatus, resp.GetYii2Exception())
				return
			} else if resp.ResponseType == server.MobileSetStatusCodeType {
				WriteJSON(rw, httpStatus, server.MobileResponse{
					Status:  httpStatus,
					Message: resp.Message,
					Data:    resp.Data,
					Version: resp.Version})
				return

			} else if resp.ResponseType == server.StreamResponseType {
				WriteStream(rw, resp.Message, resp.GetStream())
			} else {
				// MobileResponse status always http.StatusOK but Status field 200,403,400,500,...
				WriteJSON(rw, http.StatusOK, server.MobileResponse{
					Status:  httpStatus,
					Message: resp.Message,
					Data:    resp.Data,
					Version: resp.Version})
				return
			}

			// Can extend more type of response here
		}
	}
}

func (h BaseHTTPHandler) Authentication(rw http.ResponseWriter, r *http.Request) (*app.Context, error) {
	return app.NewContext(rw, r, "", h.IsStaging()), nil
}

// CapturePanic Last layer to capture panic which might halt the whole application.
func (h BaseHTTPHandler) CapturePanic(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {

				WriteJSON(rw, http.StatusInternalServerError,
					"Request is halted unexpectedly, please contact the administrator.")
			}
		}()
		next(rw, r)
	}
}

func WriteJSON(rw http.ResponseWriter, httpCode int, data interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	if httpCode != 0 {
		rw.WriteHeader(httpCode)
	}
	if data != nil {
		_ = json.NewEncoder(rw).Encode(data)
	}
}
func WriteStream(rw http.ResponseWriter, fileName string, b bytes.Buffer) {
	rw.Header().Set("File-Name", fileName)
	rw.Header().Set("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	rw.Header().Set("Content-Type", "application/octet-stream")
	rw.Header().Set("Content-Length", fmt.Sprintf("%d", b.Len()))
	rw.Header().Set("Content-Transfer-Encoding", "binary")
	rw.Header().Set("Expires", "0")

	rw.WriteHeader(http.StatusOK)

	_, err := rw.Write(b.Bytes())
	if err != nil {
		logrus.Errorln(fmt.Errorf("Stream %s %v", fileName, err))
		return
	}
}

// common missing id response, can be used by any endpoint route, i.e. h.Route("GET", "/project/detail/", h.base.NotFoundHandler)
func (h BaseHTTPHandler) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusBadRequest, server.Yii2HTTPExceptionResponse{
		Name:    "Bad Request",
		Message: "Missing required parameters: id",
		Code:    0,
		Status:  http.StatusBadRequest,
	})
}

// common method not allowed response (405)
func (h BaseHTTPHandler) MethodNotAllowedHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		WriteJSON(rw, http.StatusMethodNotAllowed, server.Yii2HTTPExceptionResponse{
			Name: "Method Not Allowed",
			// hardcode the allowed method list for now
			Message: fmt.Sprintf("Method Not Allowed. This URL can only handle the following request methods: %s", "GET"),
			Code:    0,
			Status:  http.StatusMethodNotAllowed,
		})
	})
}

func (h BaseHTTPHandler) NotFoundRoute(ctx *app.Context) *server.Response {
	return h.AsJson(ctx, http.StatusNotFound, "page not found", nil)
}

func (h BaseHTTPHandler) MissingIDParameter(ctx *app.Context) *server.Response {
	return h.AsJson(ctx, http.StatusBadRequest, "page not found", nil)
}
