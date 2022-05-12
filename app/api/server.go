package api

import (
	"fmt"
	"net/http"
	"os"

	ofModule "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/offer/handler"

	"bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/base/handler"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/server"

	"github.com/gorilla/mux"
	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
)

// HttpServe is a http server implementation
type HttpServe struct {
	router *muxtrace.Router

	base  *handler.BaseHTTPHandler
	offer *ofModule.HTTPHandler

	v1     *mux.Router
	static *mux.Route
}

//Run runs the HTTP server application
func (h *HttpServe) Run() error {
	h.setupRouter()
	h.base.Handlers = h

	return http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("HTTP_SERVER_PORT")), h.router)
}

//New creates new API server application
func New(appName string,
	base *handler.BaseHTTPHandler,
	offer *ofModule.HTTPHandler,
) server.App {
	return &HttpServe{
		base:   base,
		offer:  offer,
		router: muxtrace.NewRouter(muxtrace.WithServiceName(appName)),
	}
}
