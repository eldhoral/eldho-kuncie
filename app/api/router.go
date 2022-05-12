package api

import (
	"fmt"
	"net/http"

	"bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/base/handler"
)

func (h *HttpServe) setupRouter() {
	// StrictSlash will treat /projects/ to be same as /projects
	h.v1 = h.router.PathPrefix("/api/v1/").Subrouter().StrictSlash(true)
	// h.static = h.router.PathPrefix("/").Subrouter().StrictSlash(true)

	// Loan Limit
	h.Route("GET", "/loanlimit/detail", h.offer.LoanLimitDetail)
	h.Route("GET", "/loanlimit/detail/{id:[0-9]+}", h.offer.LoanLimitDetailByID)
	h.Route("POST", "/loanlimit/create", h.offer.CreateLoanLimit)
	h.Route("POST", "/loanlimit/update/{id:[0-9]+}", h.offer.UpdateLoanLimit)
	h.Route("DELETE", "/loanlimit/delete", h.offer.DeleteLoanLimit)

	// Benefit
	h.Route("GET", "/benefit/detail/{id:[0-9]+}", h.offer.BenefitByID)
	h.Route("GET", "/benefit/list", h.offer.BenefitList)
	h.Route("POST", "/benefit/create", h.offer.CreateBenefit)
	h.Route("POST", "/benefit/update/{id:[0-9]+}", h.offer.UpdateBenefit)
	h.Route("DELETE", "/benefit/delete/{id:[0-9]+}", h.offer.DeleteBenefit)

	// Loan Method
	h.Route("GET", "/loanmethod/detail/{id:[0-9]+}", h.offer.LoanMethodByID)
	h.Route("GET", "/loanmethod/list", h.offer.LoanMethodList)
	h.Route("POST", "/loanmethod/create", h.offer.CreateLoanMethod)
	h.Route("POST", "/loanmethod/update/{id:[0-9]+}", h.offer.UpdateLoanMethod)
	h.Route("DELETE", "/loanmethod/delete/{id:[0-9]+}", h.offer.DeleteLoanMethod)

	// Landing Page
	h.Route("GET", "/landingpage", h.offer.LandingPage)

	// Serve static image
	h.static = h.router.PathPrefix("/assets/upload/image/").Handler(http.StripPrefix("/assets/upload/image/", http.FileServer(http.Dir("./assets/upload/image"))))

	// // Health Check
	// h.router.HandleFunc("/health-check", h.base.HealthCheck).Methods(http.MethodGet)
	// h.router.HandleFunc("/status-check", h.base.CheckStatus).Methods(http.MethodGet)
	// h.router.HandleFunc("/consumer-check", h.base.CheckConsumer).Methods(http.MethodGet)

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