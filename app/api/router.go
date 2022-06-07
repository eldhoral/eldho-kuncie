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

	// Tnc
	h.Route("GET", "/tnc/detail/{id:[0-9]+}", h.offer.TncByID)
	h.Route("GET", "/tnc/list", h.offer.TncList)
	h.Route("POST", "/tnc/create", h.offer.CreateTnc)
	h.Route("POST", "/tnc/update/{id:[0-9]+}", h.offer.UpdateTnc)
	h.Route("POST", "/tnc/mobile/update", h.offer.UpdateTncMobile)
	h.Route("DELETE", "/tnc/delete/{id:[0-9]+}", h.offer.DeleteTnc)

	// Tnc Title
	h.Route("GET", "/tnc/title/detail/{id:[0-9]+}", h.offer.TncTitleByID)
	h.Route("GET", "/tnc/title/list", h.offer.TncTitleList)
	h.Route("POST", "/tnc/title/create", h.offer.CreateTncTitle)
	h.Route("POST", "/tnc/title/update/{id:[0-9]+}", h.offer.UpdateTncTitle)
	h.Route("POST", "/tnc/mobile/update", h.offer.UpdateTncMobile)
	h.Route("DELETE", "/tnc/title/delete/{id:[0-9]+}", h.offer.DeleteTncTitle)

	// Tnc Subtitle
	h.Route("GET", "/tnc/subtitle/detail/{id:[0-9]+}", h.offer.TncSubtitleByID)
	h.Route("GET", "/tnc/subtitle/list", h.offer.TncSubtitleList)
	h.Route("POST", "/tnc/subtitle/create", h.offer.CreateTncSubtitle)
	h.Route("POST", "/tnc/subtitle/update/{id:[0-9]+}", h.offer.UpdateTncSubtitle)
	h.Route("DELETE", "/tnc/subtitle/delete/{id:[0-9]+}", h.offer.DeleteTncSubtitle)

	// Tnc Explain
	h.Route("GET", "/tnc/explain/detail/{id:[0-9]+}", h.offer.TncExplainByID)
	h.Route("GET", "/tnc/explain/list", h.offer.TncExplainList)
	h.Route("POST", "/tnc/explain/create", h.offer.CreateTncExplain)
	h.Route("POST", "/tnc/explain/update/{id:[0-9]+}", h.offer.UpdateTncExplain)
	h.Route("DELETE", "/tnc/explain/delete/{id:[0-9]+}", h.offer.DeleteTncExplain)

	// Cost
	h.Route("GET", "/cost/detail/{id:[0-9]+}", h.about.CostByID)
	h.Route("GET", "/cost/list", h.about.CostList)
	h.Route("POST", "/cost/create", h.about.CreateCost)
	h.Route("POST", "/cost/update/{id:[0-9]+}", h.about.UpdateCost)
	h.Route("DELETE", "/cost/delete/{id:[0-9]+}", h.about.DeleteCost)

	// Cost Explain
	h.Route("GET", "/cost/explain/detail/{id:[0-9]+}", h.about.CostExplainByID)
	h.Route("GET", "/cost/explain/list", h.about.CostExplainList)
	h.Route("POST", "/cost/explain/create", h.about.CreateCostExplain)
	h.Route("POST", "/cost/explain/update/{id:[0-9]+}", h.about.UpdateCostExplain)
	h.Route("DELETE", "/cost/explain/delete/{id:[0-9]+}", h.about.DeleteCostExplain)

	// FAQ
	h.Route("GET", "/faq/detail/{id:[0-9]+}", h.about.FAQByID)
	h.Route("GET", "/faq/list", h.about.FAQList)
	h.Route("POST", "/faq/create", h.about.CreateFAQ)
	h.Route("POST", "/faq/update/{id:[0-9]+}", h.about.UpdateFAQ)
	h.Route("DELETE", "/faq/delete/{id:[0-9]+}", h.about.DeleteFAQ)

	// FAQ Title
	h.Route("GET", "/faq/title/detail/{id:[0-9]+}", h.about.FAQTitleByID)
	h.Route("GET", "/faq/title/list", h.about.FAQTitleList)
	h.Route("POST", "/faq/title/create", h.about.CreateFAQTitle)
	h.Route("POST", "/faq/title/update/{id:[0-9]+}", h.about.UpdateFAQTitle)
	h.Route("DELETE", "/faq/title/delete/{id:[0-9]+}", h.about.DeleteFAQTitle)

	// Landing Page
	h.Route("GET", "/landingpage", h.offer.LandingPage)

	// Tnc Page
	h.Route("GET", "/tncpage", h.offer.TncPage)

	// Tnc Mobile
	h.Route("GET", "/tncpage/mobile", h.offer.TncPageMobile)

	// Cost Explanation Page
	h.Route("GET", "/costexplanationpage", h.about.CostExplainationPage)

	// Benefit List Page
	h.Route("GET", "/benefitpage", h.offer.BenefitListPage)

	// FAQ
	h.Route("GET", "/faqpage", h.about.FAQPage)

	// Serve static image localhost
	h.static = h.router.PathPrefix("/api/v1/").Handler(http.StripPrefix("/api/v1/", http.FileServer(http.Dir("./assets/upload/image"))))

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
