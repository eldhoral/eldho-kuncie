package handler

import (
	"net/http"
	"regexp"

	"bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/base/app"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/data"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/data/constant"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/server"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// LoanLimitDetail for h.Route("GET", "/loanlimit/detail", h.OfferService.GetLoanLimit)
func (h HTTPHandler) LoanLimitDetail(ctx *app.Context) *server.Response {
	httpStatus, service, err := h.OfferService.GetLoanLimit()
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Get Loan Limit Successfully", service)
}

// LoanLimitDetailByID for h.Route("GET", "/loanlimit/detail/{id:[0-9]+}", h.OfferService.GetLoanLimit)
func (h HTTPHandler) LoanLimitDetailByID(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	httpStatus, service, err := h.OfferService.GetLoanLimitByID(id)
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Get Loan Limit Successfully", service)
}

// CreateLoanLimit for h.Route("POST", "/loanlimit/create", h.OfferService.CreateLoanLimit)
func (h HTTPHandler) CreateLoanLimit(ctx *app.Context) *server.Response {
	formBody := ctx.GetFormBody()

	if formBody != nil {
		limit := formBody["limit"]
		err := validation.Validate(limit, validation.Required, validation.Length(0, 9),
			is.Digit, validation.Match(regexp.MustCompile(`^[0-9]*$`)))
		if err != nil {
			return h.AsMobileJson(ctx, http.StatusBadRequest, err.Error(), nil)
		}
	}
	httpStatus, service, err := h.OfferService.CreateLoanLimit(ctx.Request.FormValue("limit"))
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Create Loan Limit Successfully", service)
}

// UpdateLoanLimit for h.Route("POST", "/loanlimit/update", h.OfferService.UpdateLoanLimit)
func (h HTTPHandler) UpdateLoanLimit(ctx *app.Context) *server.Response {
	limit := ctx.Request.FormValue("limit")

	err := validation.Validate(limit, validation.Required, validation.Length(0, 9),
		is.Digit, validation.Match(regexp.MustCompile(`^[0-9]*$`)))
	if err != nil {
		return h.AsMobileJson(ctx, http.StatusBadRequest, err.Error(), nil)
	}

	httpStatus, err := h.OfferService.UpdateLoanLimit(limit)
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Create Loan Limit Successfully", err)
}

// DeleteLoanLimit for h.Route("POST", "/loanlimit/delete", h.OfferService.DeleteLoanLimit)
func (h HTTPHandler) DeleteLoanLimit(ctx *app.Context) *server.Response {
	status, err := h.OfferService.DeleteLoanLimit()
	if err != nil {
		return h.AsMobileJson(ctx, status, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, http.StatusOK, "Delete Loan Limit Successfully", err)
}

// BenefitByID for h.Route("GET", "/benefit/detail/{id:[0-9]+}", h.OfferService.BenefitByID)
func (h HTTPHandler) BenefitByID(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	httpStatus, service, err := h.OfferService.GetBenefitByID(id)
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Get Benefit Successfully", service)
}

// BenefitList for h.Route("GET", "/benefit/list", h.OfferService.BenefitList)
func (h HTTPHandler) BenefitList(ctx *app.Context) *server.Response {
	httpStatus, service, err := h.OfferService.ListBenefit()
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Get Benefit List Successfully", service)
}

// CreateBenefit for h.Route("POST", "/benefit/create", h.OfferService.CreateBenefit)
func (h HTTPHandler) CreateBenefit(ctx *app.Context) *server.Response {
	formBody := ctx.GetFormBody()
	if formBody != nil {
		title := formBody["title"]
		if title == "" {
			return h.AsMobileJson(ctx, http.StatusBadRequest, "Title must be filled", constant.EmptyArray)
		}
		description := formBody["description"]
		if description == "" {
			return h.AsMobileJson(ctx, http.StatusBadRequest, "Description must be filled", constant.EmptyArray)
		}
	}

	if formBody == nil {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Body form is a must", constant.EmptyArray)
	}

	title := ctx.Request.FormValue("title")
	description := ctx.Request.FormValue("description")
	image, err := ctx.GetUploadFile("image")
	if err != nil {
		return h.AsMobileJson(ctx, http.StatusInternalServerError, err.Error(), nil)
	}

	if image == nil {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Image must be filled", constant.EmptyArray)
	}

	httpStatus, service, err := h.OfferService.CreateBenefit(title, description, image.Name)
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Create Benefit Successfully", service)
}

// UpdateBenefit for h.Route("POST", "/benefit/update{id:[0-9]+}", h.OfferService.UpdateBenefit)
func (h HTTPHandler) UpdateBenefit(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	formBody := ctx.GetFormBody()
	if formBody == nil {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Body form is a must", constant.EmptyArray)
	}

	title := ctx.Request.FormValue("title")
	description := ctx.Request.FormValue("description")
	params := data.NewParamsWrapper()
	params.Add("title", title)
	params.Add("description", description)
	asset, _, _ := ctx.Request.FormFile("image")
	if asset != nil {
		image, err := ctx.GetUploadFile("image")
		if err != nil {
			return h.AsMobileJson(ctx, http.StatusInternalServerError, err.Error(), nil)
		}

		httpStatus, err := h.OfferService.UpdateBenefitByID(id, params, image.Name)
		if err != nil {
			return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
		}

		return h.AsMobileJson(ctx, httpStatus, "Update Benefit Successfully", err)
	}

	if title == "" && description == "" && asset == nil {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Value form is a must to one of these", constant.EmptyArray)
	}

	httpStatus, err := h.OfferService.UpdateBenefitByID(id, params, "")
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Update Benefit Successfully", err)
}

// DeleteBenefit for h.Route("POST", "/benefit/delete{id:[0-9]+}", h.OfferService.DeleteBenefit)
func (h HTTPHandler) DeleteBenefit(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	status, err := h.OfferService.DeleteBenefitByID(id)
	if err != nil {
		return h.AsMobileJson(ctx, status, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, http.StatusOK, "Delete Benefit Successfully", err)
}

// LoanMethodByID for h.Route("GET", "/loanmethod/detail/{id:[0-9]+}", h.OfferService.LoanMethodByID)
func (h HTTPHandler) LoanMethodByID(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	httpStatus, service, err := h.OfferService.GetLoanMethodByID(id)
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Get Loan Method Successfully", service)
}

// LoanMethodList for h.Route("GET", "/loanmethod/list", h.OfferService.LoanMethodList)
func (h HTTPHandler) LoanMethodList(ctx *app.Context) *server.Response {
	httpStatus, service, err := h.OfferService.ListLoanMethod()
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Get Loan Method List Successfully", service)
}

// CreateLoanMethod for h.Route("POST", "/loanmethod/create", h.OfferService.CreateLoanMethod)
func (h HTTPHandler) CreateLoanMethod(ctx *app.Context) *server.Response {
	formBody := ctx.GetFormBody()
	if formBody != nil {
		title := formBody["title"]
		if title == "" {
			return h.AsMobileJson(ctx, http.StatusBadRequest, "Title must be filled", constant.EmptyArray)
		}
		description := formBody["description"]
		if description == "" {
			return h.AsMobileJson(ctx, http.StatusBadRequest, "Description must be filled", constant.EmptyArray)
		}
	}

	if formBody == nil {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Body form is a must", constant.EmptyArray)
	}

	title := ctx.Request.FormValue("title")
	description := ctx.Request.FormValue("description")

	httpStatus, service, err := h.OfferService.CreateLoanMethod(title, description)
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Create Loan Method Successfully", service)
}

// UpdateLoanMethod for h.Route("POST", "/loanmethod/update{id:[0-9]+}", h.OfferService.UpdateLoanMethod)
func (h HTTPHandler) UpdateLoanMethod(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	formBody := ctx.GetFormBody()
	if formBody == nil {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Body form is a must", constant.EmptyArray)
	}

	title := ctx.Request.FormValue("title")
	description := ctx.Request.FormValue("description")

	if title == "" && description == "" {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Value form is a must to one of these", constant.EmptyArray)
	}

	httpStatus, err := h.OfferService.UpdateLoanMethodByID(id, title, description)
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Update Loan Method Successfully", err)
}

// DeleteLoanMethod for h.Route("POST", "/loanmethod/delete{id:[0-9]+}", h.OfferService.DeleteLoanMethod)
func (h HTTPHandler) DeleteLoanMethod(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	status, err := h.OfferService.DeleteLoanMethodByID(id)
	if err != nil {
		return h.AsMobileJson(ctx, status, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, http.StatusOK, "Delete Loan Method Successfully", err)
}

// UpdateTncMobile for h.Route("POST", "/tnc/mobile/update", h.OfferService.UpdateTncMobile)
func (h HTTPHandler) UpdateTncMobile(ctx *app.Context) *server.Response {
	formBody := ctx.GetFormBody()
	if formBody == nil {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Body form is a must", constant.EmptyArray)
	}

	description := ctx.Request.FormValue("description")
	params := data.NewParamsWrapper()
	params.Add("description", description)

	httpStatus, err := h.OfferService.UpdateTncMobile(params)
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Update Tnc Mobile Successfully", err)
}

// LandingPage for h.Route("GET", "/landingpage", h.OfferService.GetLandingPage)
func (h HTTPHandler) LandingPage(ctx *app.Context) *server.Response {
	httpStatus, service, err := h.OfferService.GetLandingPage()
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Get Landing Page Successfully", service)
}

// TncPageMobile for h.Route("GET", "/tncpage/mobile", h.OfferService.TncPageMobile)
func (h HTTPHandler) TncPageMobile(ctx *app.Context) *server.Response {
	httpStatus, service, err := h.OfferService.GetTncPageMobile()
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Get TNC Page Mobile Successfully", service)
}

// BenefitListPage for h.Route("GET", "/benefitpage", h.AboutService.BenefitListPage)
func (h HTTPHandler) BenefitListPage(ctx *app.Context) *server.Response {
	httpStatus, service, err := h.OfferService.ListBenefit()
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Get List Benefit Page Successfully", service)
}
