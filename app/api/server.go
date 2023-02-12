package api

import (
	"fmt"
	"net/http"
	"os"

	storModule "github.com/eldhoral/eldho-kuncie/internal/store/handler"

	"github.com/eldhoral/eldho-kuncie/internal/base/handler"
	"github.com/eldhoral/eldho-kuncie/pkg/server"

	"github.com/gorilla/mux"
	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
)

// HttpServe is a http server implementation
type HttpServe struct {
	router *muxtrace.Router

	base  *handler.BaseHTTPHandler
	store *storModule.HTTPHandler

	v1     *mux.Router
	static *mux.Route
}

// Run runs the HTTP server application
func (h *HttpServe) Run() error {
	h.setupRouter()
	h.base.Handlers = h

	return http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("HTTP_SERVER_PORT")), h.router)
}

// New creates new API server application
func New(appName string,
	base *handler.BaseHTTPHandler,
	store *storModule.HTTPHandler,
) server.App {
	return &HttpServe{
		base:   base,
		store:  store,
		router: muxtrace.NewRouter(muxtrace.WithServiceName(appName)),
	}
}
