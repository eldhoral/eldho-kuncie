package handler

import (
	"net/http"

	"bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/base/app"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/data"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/data/constant"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/server"
)

// LoanLimitDetail for h.Route("GET", "/loanlimit/detail", h.OfferService.GetLoanLimit)
func (h HTTPHandler) LoanLimitDetail(ctx *app.Context) *server.Response {
	service, err := h.OfferService.GetLoanLimit()
	if err != nil {
		return h.AsMobileJson(ctx, http.StatusForbidden, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, http.StatusOK, "Get Loan Limit Successfully", service)
}

// LoanLimitDetailByID for h.Route("GET", "/loanlimit/detail/{id:[0-9]+}", h.OfferService.GetLoanLimit)
func (h HTTPHandler) LoanLimitDetailByID(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	service, err := h.OfferService.GetLoanLimitByID(id)
	if err != nil {
		return h.AsMobileJson(ctx, http.StatusForbidden, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, http.StatusOK, "Get Loan Limit Successfully", service)
}

// CreateLoanLimit for h.Route("POST", "/loanlimit/create", h.OfferService.CreateLoanLimit)
func (h HTTPHandler) CreateLoanLimit(ctx *app.Context) *server.Response {
	formBody := ctx.GetFormBody()

	if formBody != nil {
		limit := formBody["limit"]
		if limit == "" {
			return h.AsMobileJson(ctx, http.StatusForbidden, "Limit must be filled", constant.EmptyArray)
		}
	}
	service, err := h.OfferService.CreateLoanLimit(ctx.Request.FormValue("limit"))
	if err != nil {
		return h.AsMobileJson(ctx, http.StatusForbidden, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, http.StatusOK, "Create Loan Limit Successfully", service)
}

// UpdateLoanLimit for h.Route("POST", "/loanlimit/update", h.OfferService.UpdateLoanLimit)
func (h HTTPHandler) UpdateLoanLimit(ctx *app.Context) *server.Response {
	limit := ctx.Request.FormValue("limit")
	if limit == "" {
		return h.AsMobileJson(ctx, http.StatusForbidden, "Limit must be filled", constant.EmptyArray)
	}

	err := h.OfferService.UpdateLoanLimit(limit)
	if err != nil {
		return h.AsMobileJson(ctx, http.StatusForbidden, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, http.StatusOK, "Create Loan Limit Successfully", err)
}

// DeleteLoanLimit for h.Route("POST", "/loanlimit/delete", h.OfferService.DeleteLoanLimit)
func (h HTTPHandler) DeleteLoanLimit(ctx *app.Context) *server.Response {
	err := h.OfferService.DeleteLoanLimit()
	if err != nil {
		return h.AsMobileJson(ctx, http.StatusForbidden, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, http.StatusOK, "Delete Loan Limit Successfully", err)
}

// BenefitByID for h.Route("GET", "/benefit/detail/{id:[0-9]+}", h.OfferService.BenefitByID)
func (h HTTPHandler) BenefitByID(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	service, err := h.OfferService.GetBenefitByID(id)
	if err != nil {
		return h.AsMobileJson(ctx, http.StatusForbidden, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, http.StatusOK, "Get Benefit Successfully", service)
}

// BenefitList for h.Route("GET", "/benefit/list", h.OfferService.BenefitList)
func (h HTTPHandler) BenefitList(ctx *app.Context) *server.Response {
	service, err := h.OfferService.ListBenefit()
	if err != nil {
		return h.AsMobileJson(ctx, http.StatusForbidden, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, http.StatusOK, "Get Benefit List Successfully", service)
}

// CreateBenefit for h.Route("POST", "/benefit/create", h.OfferService.CreateBenefit)
func (h HTTPHandler) CreateBenefit(ctx *app.Context) *server.Response {
	formBody := ctx.GetFormBody()
	if formBody != nil {
		title := formBody["title"]
		if title == "" {
			return h.AsMobileJson(ctx, http.StatusForbidden, "Title must be filled", constant.EmptyArray)
		}
		description := formBody["description"]
		if description == "" {
			return h.AsMobileJson(ctx, http.StatusForbidden, "Description must be filled", constant.EmptyArray)
		}
	}

	if formBody == nil {
		return h.AsMobileJson(ctx, http.StatusForbidden, "Body form is a must", constant.EmptyArray)
	}

	title := ctx.Request.FormValue("title")
	description := ctx.Request.FormValue("description")
	image, err := ctx.GetUploadFile("image")
	if err != nil {
		return h.AsMobileJson(ctx, http.StatusForbidden, err.Error(), nil)
	}

	if image == nil {
		return h.AsMobileJson(ctx, http.StatusForbidden, "Image must be filled", constant.EmptyArray)
	}

	service, err := h.OfferService.CreateBenefit(title, description, image.Path)
	if err != nil {
		return h.AsMobileJson(ctx, http.StatusForbidden, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, http.StatusOK, "Create Benefit Successfully", service)
}

// UpdateBenefit for h.Route("POST", "/benefit/update{id:[0-9]+}", h.OfferService.UpdateBenefit)
func (h HTTPHandler) UpdateBenefit(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	formBody := ctx.GetFormBody()
	if formBody == nil {
		return h.AsMobileJson(ctx, http.StatusForbidden, "Body form is a must", constant.EmptyArray)
	}

	title := ctx.Request.FormValue("title")
	description := ctx.Request.FormValue("description")
	params := data.NewParamsWrapper()
	params.Add("title", title)
	params.Add("description", description)
	asset, _, _ := ctx.Request.FormFile("image")
	if asset != nil {
		image, _ := ctx.GetUploadFile("image")

		err := h.OfferService.UpdateBenefitByID(id, params, image.Path)
		if err != nil {
			return h.AsMobileJson(ctx, http.StatusForbidden, err.Error(), nil)
		}

		return h.AsMobileJson(ctx, http.StatusOK, "Update Benefit Successfully", err)
	}

	if title == "" && description == "" && asset == nil {
		return h.AsMobileJson(ctx, http.StatusForbidden, "Value form is a must to one of these", constant.EmptyArray)
	}

	err := h.OfferService.UpdateBenefitByID(id, params, "")
	if err != nil {
		return h.AsMobileJson(ctx, http.StatusForbidden, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, http.StatusOK, "Update Benefit Successfully", err)
}

// DeleteBenefit for h.Route("POST", "/benefit/delete{id:[0-9]+}", h.OfferService.DeleteBenefit)
func (h HTTPHandler) DeleteBenefit(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	err := h.OfferService.DeleteBenefitByID(id)
	if err != nil {
		return h.AsMobileJson(ctx, http.StatusForbidden, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, http.StatusOK, "Delete Benefit Successfully", err)
}

// LoanMethodByID for h.Route("GET", "/loanmethod/detail/{id:[0-9]+}", h.OfferService.LoanMethodByID)
func (h HTTPHandler) LoanMethodByID(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	service, err := h.OfferService.GetLoanMethodByID(id)
	if err != nil {
		return h.AsMobileJson(ctx, http.StatusForbidden, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, http.StatusOK, "Get Loan Method Successfully", service)
}

// LoanMethodList for h.Route("GET", "/loanmethod/list", h.OfferService.LoanMethodList)
func (h HTTPHandler) LoanMethodList(ctx *app.Context) *server.Response {
	service, err := h.OfferService.ListLoanMethod()
	if err != nil {
		return h.AsMobileJson(ctx, http.StatusForbidden, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, http.StatusOK, "Get Loan Method List Successfully", service)
}

// CreateLoanMethod for h.Route("POST", "/loanmethod/create", h.OfferService.CreateLoanMethod)
func (h HTTPHandler) CreateLoanMethod(ctx *app.Context) *server.Response {
	formBody := ctx.GetFormBody()
	if formBody != nil {
		title := formBody["title"]
		if title == "" {
			return h.AsMobileJson(ctx, http.StatusForbidden, "Title must be filled", constant.EmptyArray)
		}
		description := formBody["description"]
		if description == "" {
			return h.AsMobileJson(ctx, http.StatusForbidden, "Description must be filled", constant.EmptyArray)
		}
	}

	if formBody == nil {
		return h.AsMobileJson(ctx, http.StatusForbidden, "Body form is a must", constant.EmptyArray)
	}

	title := ctx.Request.FormValue("title")
	description := ctx.Request.FormValue("description")

	service, err := h.OfferService.CreateLoanMethod(title, description)
	if err != nil {
		return h.AsMobileJson(ctx, http.StatusForbidden, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, http.StatusOK, "Create Loan Method Successfully", service)
}

// UpdateLoanMethod for h.Route("POST", "/loanmethod/update{id:[0-9]+}", h.OfferService.UpdateLoanMethod)
func (h HTTPHandler) UpdateLoanMethod(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	formBody := ctx.GetFormBody()
	if formBody == nil {
		return h.AsMobileJson(ctx, http.StatusForbidden, "Body form is a must", constant.EmptyArray)
	}

	title := ctx.Request.FormValue("title")
	description := ctx.Request.FormValue("description")

	if title == "" && description == "" {
		return h.AsMobileJson(ctx, http.StatusForbidden, "Value form is a must to one of these", constant.EmptyArray)
	}

	err := h.OfferService.UpdateLoanMethodByID(id, title, description)
	if err != nil {
		return h.AsMobileJson(ctx, http.StatusForbidden, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, http.StatusOK, "Update Loan Method Successfully", err)
}

// DeleteLoanMethod for h.Route("POST", "/loanmethod/delete{id:[0-9]+}", h.OfferService.DeleteLoanMethod)
func (h HTTPHandler) DeleteLoanMethod(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	err := h.OfferService.DeleteLoanMethodByID(id)
	if err != nil {
		return h.AsMobileJson(ctx, http.StatusForbidden, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, http.StatusOK, "Delete Loan Method Successfully", err)
}

// LandingPage for h.Route("GET", "/landingpage", h.OfferService.GetLandingPage)
func (h HTTPHandler) LandingPage(ctx *app.Context) *server.Response {
	service, err := h.OfferService.GetLandingPage()
	if err != nil {
		return h.AsMobileJson(ctx, http.StatusForbidden, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, http.StatusOK, "Get Landing Page Successfully", service)
}
