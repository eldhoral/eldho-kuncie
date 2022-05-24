package handler

import (
	"net/http"
	"strconv"

	"bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/base/app"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/data"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/data/constant"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/server"
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
		if limit == "" {
			return h.AsMobileJson(ctx, http.StatusBadRequest, "Limit must be filled", constant.EmptyArray)
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
	if limit == "" {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Limit must be filled", constant.EmptyArray)
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

	httpStatus, service, err := h.OfferService.CreateBenefit(title, description, image.Path)
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

		httpStatus, err := h.OfferService.UpdateBenefitByID(id, params, image.Path)
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

// TncByID for h.Route("GET", "/tnc/detail/{id:[0-9]+}", h.OfferService.TncByID)
func (h HTTPHandler) TncByID(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	httpStatus, service, err := h.OfferService.GetTncByID(id)
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Get Tnc Successfully", service)
}

// TncList for h.Route("GET", "/tnc/list", h.OfferService.ListTnc)
func (h HTTPHandler) TncList(ctx *app.Context) *server.Response {
	httpStatus, service, err := h.OfferService.ListTnc()
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Get Tnc List Successfully", service)
}

// CreateTnc for h.Route("POST", "/tnc/create", h.OfferService.CreateTnc)
func (h HTTPHandler) CreateTnc(ctx *app.Context) *server.Response {
	formBody := ctx.GetFormBody()
	if formBody != nil {
		title := formBody["title"]
		if title == "" {
			return h.AsMobileJson(ctx, http.StatusBadRequest, "Title must be filled", constant.EmptyArray)
		}
	}

	if formBody == nil {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Body form is a must", constant.EmptyArray)
	}

	title := ctx.Request.FormValue("title")

	httpStatus, service, err := h.OfferService.CreateTnc(title)
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Create Tnc Successfully", service)
}

// UpdateTnc for h.Route("POST", "/tnc/update/{id:[0-9]+}", h.OfferService.UpdateTnc)
func (h HTTPHandler) UpdateTnc(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	formBody := ctx.GetFormBody()
	if formBody == nil {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Body form is a must", constant.EmptyArray)
	}

	title := ctx.Request.FormValue("title")
	params := data.NewParamsWrapper()
	params.Add("title", title)

	httpStatus, err := h.OfferService.UpdateTncByID(id, params)
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Update Tnc Successfully", err)
}

// DeleteTnc for h.Route("POST", "/tnc/delete/{id:[0-9]+}", h.OfferService.DeleteTnc)
func (h HTTPHandler) DeleteTnc(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	status, err := h.OfferService.DeleteTncByID(id)
	if err != nil {
		return h.AsMobileJson(ctx, status, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, http.StatusOK, "Delete Tnc Successfully", err)
}

// TncTitleByID for h.Route("GET", "/tnc/title/detail/{id:[0-9]+}", h.OfferService.TncTitleByID)
func (h HTTPHandler) TncTitleByID(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	httpStatus, service, err := h.OfferService.GetTncTitleByID(id)
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Get Tnc Title Successfully", service)
}

// TncTitleList for h.Route("GET", "/tnc/title/list", h.OfferService.TncTitleList)
func (h HTTPHandler) TncTitleList(ctx *app.Context) *server.Response {
	httpStatus, service, err := h.OfferService.ListTncTitle()
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Get Tnc Title List Successfully", service)
}

// CreateTncTitle for h.Route("POST", "/tnc/title/create", h.OfferService.CreateTncTitle)
func (h HTTPHandler) CreateTncTitle(ctx *app.Context) *server.Response {
	formBody := ctx.GetFormBody()
	if formBody != nil {
		idTnc := formBody["id_tnc"]
		if idTnc == "" {
			return h.AsMobileJson(ctx, http.StatusBadRequest, "ID Tnc must be filled", constant.EmptyArray)
		}
		title := formBody["title"]
		if title == "" {
			return h.AsMobileJson(ctx, http.StatusBadRequest, "Title must be filled", constant.EmptyArray)
		}
	}

	if formBody == nil {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Body form is a must", constant.EmptyArray)
	}

	idTnc, err := strconv.Atoi(ctx.Request.FormValue("id_tnc"))
	if err != nil {
		return h.AsMobileJson(ctx, http.StatusInternalServerError, err.Error(), nil)
	}
	title := ctx.Request.FormValue("title")
	httpStatus, service, err := h.OfferService.CreateTncTitle(int64(idTnc), title)
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Create Tnc Title Successfully", service)
}

// UpdateTncTitle for h.Route("POST", "/tnc/title/update/{id:[0-9]+}", h.OfferService.UpdateTncTitle)
func (h HTTPHandler) UpdateTncTitle(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	formBody := ctx.GetFormBody()
	if formBody == nil {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Body form is a must", constant.EmptyArray)
	}

	idTnc := ctx.Request.FormValue("id_tnc")
	title := ctx.Request.FormValue("title")
	params := data.NewParamsWrapper()
	params.Add("id_tnc", idTnc)
	params.Add("title", title)

	httpStatus, err := h.OfferService.UpdateTncTitleByID(id, params)
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Update Tnc Title Successfully", err)
}

// DeleteTncTitle for h.Route("POST", "/tnc/title/delete/{id:[0-9]+}", h.OfferService.DeleteTncTitle)
func (h HTTPHandler) DeleteTncTitle(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	status, err := h.OfferService.DeleteTncTitleByID(id)
	if err != nil {
		return h.AsMobileJson(ctx, status, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, http.StatusOK, "Delete Tnc Title Successfully", err)
}

// TncSubtitleByID for h.Route("GET", "/tnc/subtitle/detail/{id:[0-9]+}", h.OfferService.TncSubtitleByID)
func (h HTTPHandler) TncSubtitleByID(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	httpStatus, service, err := h.OfferService.GetTncSubtitleByID(id)
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Get Tnc Subtitle Successfully", service)
}

// TncSubtitleList for h.Route("GET", "/tnc/subtitle/list", h.OfferService.TncSubtitleList)
func (h HTTPHandler) TncSubtitleList(ctx *app.Context) *server.Response {
	httpStatus, service, err := h.OfferService.ListTncSubtitle()
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Get Tnc Subtitle List Successfully", service)
}

// CreateTncSubtitle for h.Route("POST", "/tnc/subtitle/create", h.OfferService.CreateTncSubtitle)
func (h HTTPHandler) CreateTncSubtitle(ctx *app.Context) *server.Response {
	formBody := ctx.GetFormBody()
	if formBody != nil {
		idTncTitle := formBody["id_tnc_title"]
		if idTncTitle == "" {
			return h.AsMobileJson(ctx, http.StatusBadRequest, "ID Tnc Title must be filled", constant.EmptyArray)
		}
		subtitle := formBody["subtitle"]
		if subtitle == "" {
			return h.AsMobileJson(ctx, http.StatusBadRequest, "Subtitle must be filled", constant.EmptyArray)
		}
	}

	if formBody == nil {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Body form is a must", constant.EmptyArray)
	}

	idTncTitle, err := strconv.Atoi(ctx.Request.FormValue("id_tnc_title"))
	if err != nil {
		return h.AsMobileJson(ctx, http.StatusInternalServerError, err.Error(), nil)
	}
	subtitle := ctx.Request.FormValue("subtitle")
	httpStatus, service, err := h.OfferService.CreateTncSubtitle(int64(idTncTitle), subtitle)
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Create Tnc Title Successfully", service)
}

// UpdateTncTitle for h.Route("POST", "/tnc/subtitle/update/{id:[0-9]+}", h.OfferService.UpdateTncTitle)
func (h HTTPHandler) UpdateTncSubtitle(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	formBody := ctx.GetFormBody()
	if formBody == nil {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Body form is a must", constant.EmptyArray)
	}

	idTncTitle := ctx.Request.FormValue("id_tnc_title")
	subtitle := ctx.Request.FormValue("subtitle")
	params := data.NewParamsWrapper()
	params.Add("id_tnc_title", idTncTitle)
	params.Add("subtitle", subtitle)

	httpStatus, err := h.OfferService.UpdateTncSubtitleByID(id, params)
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Update Tnc Subtitle Successfully", err)
}

// DeleteTncTitle for h.Route("POST", "/tnc/subtitle/delete/{id:[0-9]+}", h.OfferService.DeleteTncTitle)
func (h HTTPHandler) DeleteTncSubtitle(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	status, err := h.OfferService.DeleteTncSubtitleByID(id)
	if err != nil {
		return h.AsMobileJson(ctx, status, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, http.StatusOK, "Delete Tnc Subtitle Successfully", err)
}

// TncExplainByID for h.Route("GET", "/tnc/explain/detail/{id:[0-9]+}", h.OfferService.TncExplainByID)
func (h HTTPHandler) TncExplainByID(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	httpStatus, service, err := h.OfferService.GetTncExplainByID(id)
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Get Tnc Explain Successfully", service)
}

// TncExplainList for h.Route("GET", "/tnc/explain/list", h.OfferService.TncExplainList)
func (h HTTPHandler) TncExplainList(ctx *app.Context) *server.Response {
	httpStatus, service, err := h.OfferService.ListTncExplain()
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Get Tnc Explain List Successfully", service)
}

// CreateTncExplain for h.Route("POST", "/tnc/explain/create", h.OfferService.CreateTncExplain)
func (h HTTPHandler) CreateTncExplain(ctx *app.Context) *server.Response {
	formBody := ctx.GetFormBody()
	if formBody != nil {
		idTnc := formBody["id_tnc"]
		if idTnc == "" {
			return h.AsMobileJson(ctx, http.StatusBadRequest, "ID Tnc must be filled", constant.EmptyArray)
		}
		idTncTitle := formBody["id_tnc_title"]
		if idTncTitle == "" {
			return h.AsMobileJson(ctx, http.StatusBadRequest, "ID Tnc Title must be filled", constant.EmptyArray)
		}
		description := formBody["description"]
		if description == "" {
			return h.AsMobileJson(ctx, http.StatusBadRequest, "Description must be filled", constant.EmptyArray)
		}
	}

	if formBody == nil {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Body form is a must", constant.EmptyArray)
	}

	idTnc, err := strconv.Atoi(ctx.Request.FormValue("id_tnc"))
	idTncTitle, err := strconv.Atoi(ctx.Request.FormValue("id_tnc_title"))
	idTncSubtitle, _ := strconv.Atoi(ctx.Request.FormValue("id_tnc_subtitle"))
	if err != nil {
		return h.AsMobileJson(ctx, http.StatusInternalServerError, err.Error(), nil)
	}
	idTncSubtitleConverted := int64(idTncSubtitle)
	description := ctx.Request.FormValue("description")
	httpStatus, service, err := h.OfferService.CreateTncExplain(int64(idTnc), int64(idTncTitle), &idTncSubtitleConverted, description)
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Create Tnc Explain Successfully", service)
}

// UpdateTncExplain for h.Route("POST", "/tnc/explain/update/{id:[0-9]+}", h.OfferService.UpdateTncExplain)
func (h HTTPHandler) UpdateTncExplain(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	formBody := ctx.GetFormBody()
	if formBody == nil {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Body form is a must", constant.EmptyArray)
	}

	description := ctx.Request.FormValue("description")
	params := data.NewParamsWrapper()
	params.Add("description", description)

	httpStatus, err := h.OfferService.UpdateTncExplainByID(id, params)
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Update Tnc Explain Successfully", err)
}

// DeleteTncExplain for h.Route("POST", "/tnc/explain/delete/{id:[0-9]+}", h.OfferService.DeleteTncExplain)
func (h HTTPHandler) DeleteTncExplain(ctx *app.Context) *server.Response {
	id := ctx.GetVarInt64("id")
	if ctx.HasError() {
		return h.AsMobileJson(ctx, http.StatusBadRequest, "Missing required parameters: id", nil)
	}

	status, err := h.OfferService.DeleteTncExplainByID(id)
	if err != nil {
		return h.AsMobileJson(ctx, status, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, http.StatusOK, "Delete Tnc Explain Successfully", err)
}

// LandingPage for h.Route("GET", "/landingpage", h.OfferService.GetLandingPage)
func (h HTTPHandler) LandingPage(ctx *app.Context) *server.Response {
	httpStatus, service, err := h.OfferService.GetLandingPage()
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Get Landing Page Successfully", service)
}

// TncPage for h.Route("GET", "/tncpage", h.OfferService.TncPage)
func (h HTTPHandler) TncPage(ctx *app.Context) *server.Response {
	httpStatus, service, err := h.OfferService.GetTncPage()
	if err != nil {
		return h.AsMobileJson(ctx, httpStatus, err.Error(), nil)
	}

	return h.AsMobileJson(ctx, httpStatus, "Get TNC Page Successfully", service)
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
