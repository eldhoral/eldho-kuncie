package handler

import (
	"net/http"

	"bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/base/app"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/data"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/data/constant"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/server"
)

// CostByID for h.Route("GET", "/cost/detail/{id:[0-9]+}", h.AboutService.CostByID)
func (h HTTPHandler) CostByID(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	httpStatus, service, err := h.AboutService.GetCostByID(id)
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Get Cost Successfully", service)
}

// CostList for h.Route("GET", "/cost/list", h.AboutService.CostList)
func (h HTTPHandler) CostList(ctx *app.Context) *server.Response {
	httpStatus, service, err := h.AboutService.ListCost()
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Get Cost List Successfully", service)
}

// CreateCost for h.Route("POST", "/cost/create", h.AboutService.CreateCost)
func (h HTTPHandler) CreateCost(ctx *app.Context) *server.Response {
	formBody := ctx.GetFormBody()
	if formBody != nil {
		loanOption := formBody["loan_option"]
		if loanOption == "" {
			return h.AsMobileJson(ctx, http.StatusBadRequest, "Loan Option must be filled", constant.EmptyArray)
		}
		idLoanOption := formBody["id_loan_option"]
		if idLoanOption == "" {
			return h.AsMobileJson(ctx, http.StatusBadRequest, "ID Loan Option must be filled", constant.EmptyArray)
		}
		interest := formBody["interest"]
		if interest == "" {
			return h.AsMobileJson(ctx, http.StatusBadRequest, "Interest must be filled", constant.EmptyArray)
		}
		adminFee := formBody["admin_fee"]
		if adminFee == "" {
			return h.AsMobileJson(ctx, http.StatusBadRequest, "Admin Fee must be filled", constant.EmptyArray)
		}
		finePerDay := formBody["fine_per_day"]
		if finePerDay == "" {
			return h.AsMobileJson(ctx, http.StatusBadRequest, "Fine per Day must be filled", constant.EmptyArray)
		}
		description := formBody["description"]
		if description == "" {
			return h.AsMobileJson(ctx, http.StatusBadRequest, "Description must be filled", constant.EmptyArray)
		}
	}

	if formBody == nil {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Body form is a must", constant.EmptyArray)
	}

	loanOption := ctx.Request.FormValue("loan_option")
	idLoanOption := ctx.Request.FormValue("id_loan_option")
	interest := ctx.Request.FormValue("interest")
	adminFee := ctx.Request.FormValue("admin_fee")
	finePerDay := ctx.Request.FormValue("fine_per_day")
	description := ctx.Request.FormValue("description")

	params := data.NewParamsWrapper()
	params.Add("loan_option", loanOption)
	params.Add("id_loan_option", idLoanOption)
	params.Add("interest", interest)
	params.Add("admin_fee", adminFee)
	params.Add("fine_per_day", finePerDay)
	params.Add("description", description)

	httpStatus, service, err := h.AboutService.CreateCost(params)
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Create Cost Successfully", service)
}

// UpdateCost for h.Route("POST", "/cost/update/{id:[0-9]+}", h.AboutService.UpdateCost)
func (h HTTPHandler) UpdateCost(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	formBody := ctx.GetFormBody()
	if formBody == nil {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Body form is a must", constant.EmptyArray)
	}

	loanOption := ctx.Request.FormValue("loan_option")
	interest := ctx.Request.FormValue("interest")
	adminFee := ctx.Request.FormValue("admin_fee")
	finePerDay := ctx.Request.FormValue("fine_per_day")
	description := ctx.Request.FormValue("description")

	params := data.NewParamsWrapper()
	params.Add("loan_option", loanOption)
	params.Add("interest", interest)
	params.Add("admin_fee", adminFee)
	params.Add("fine_per_day", finePerDay)
	params.Add("description", description)

	if loanOption == "" &&
		interest == "" &&
		adminFee == "" &&
		finePerDay == "" &&
		description == "" {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Value form is a must to one of these", constant.EmptyArray)
	}

	httpStatus, err := h.AboutService.UpdateCostByID(id, params)
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Update Cost Successfully", err)
}

// DeleteCost for h.Route("POST", "/cost/delete/{id:[0-9]+}", h.AboutService.DeleteCost)
func (h HTTPHandler) DeleteCost(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	status, err := h.AboutService.DeleteCostByID(id)
	if err != nil {
		return h.AsMobileJson(ctx, status, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, http.StatusOK, "Delete Cost Successfully", err)
}

// CostExplainByID for h.Route("GET", "/cost/explain/detail/{id:[0-9]+}", h.AboutService.CostExplainByID)
func (h HTTPHandler) CostExplainByID(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	httpStatus, service, err := h.AboutService.GetCostExplainByID(id)
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Get Cost Explain Successfully", service)
}

// CostExplainList for h.Route("GET", "/cost/explain/list", h.AboutService.CostExplainList)
func (h HTTPHandler) CostExplainList(ctx *app.Context) *server.Response {
	httpStatus, service, err := h.AboutService.ListCostExplain()
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Get Cost Explain List Successfully", service)
}

// CreateCostExplain for h.Route("POST", "/cost/explain/create", h.AboutService.CreateCostExplain)
func (h HTTPHandler) CreateCostExplain(ctx *app.Context) *server.Response {
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

	params := data.NewParamsWrapper()
	params.Add("title", title)
	params.Add("description", description)

	httpStatus, service, err := h.AboutService.CreateCostExplain(params)
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Create Cost Explain Successfully", service)
}

// UpdateCostExplain for h.Route("POST", "/cost/explain/update/{id:[0-9]+}", h.AboutService.UpdateCostExplain)
func (h HTTPHandler) UpdateCostExplain(ctx *app.Context) *server.Response {
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

	if title == "" && description == "" {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Value form is a must to one of these", constant.EmptyArray)
	}

	httpStatus, err := h.AboutService.UpdateCostExplainByID(id, params)
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Update Cost Explain Successfully", err)
}

// DeleteCostExplain for h.Route("POST", "/cost/explain/delete/{id:[0-9]+}", h.AboutService.DeleteCostExplain)
func (h HTTPHandler) DeleteCostExplain(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	status, err := h.AboutService.DeleteCostExplainByID(id)
	if err != nil {
		return h.AsMobileJson(ctx, status, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, http.StatusOK, "Delete Cost Explain Successfully", err)
}
