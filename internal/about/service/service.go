package service

import (
	"database/sql"
	"errors"
	"net/http"

	modelCost "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/about/domain/cost"
	modelFaq "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/about/domain/faq"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/about/presenter/resp"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/about/repository"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/data"
)

// NewService creates new user service
func NewService(repo repository.Repository) Service {
	return &service{
		aboutRepo: repo,
	}
}

type service struct {
	aboutRepo repository.Repository
}

//Cost
func (s service) GetCostByID(id int64) (int, *modelCost.Cost, error) {
	repo, err := s.aboutRepo.GetCostByID(id)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("ID Cost not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}
func (s service) ListCost() (int, []modelCost.Cost, error) {
	repo, err := s.aboutRepo.ListCost()
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("Any ID Cost not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}
func (s service) UpdateCostByID(id int64, params data.Params) (int, error) {
	count, err := s.aboutRepo.UpdateCostByID(id, params)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if count == 0 {
		return http.StatusNotFound, errors.New("0 rows affected when updating Cost")
	}

	return http.StatusOK, nil
}

//Cost Explain
func (s service) GetCostExplainByID(id int64) (int, *modelCost.CostExplanation, error) {
	repo, err := s.aboutRepo.GetCostExplanationByID(id)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("ID Cost Explain not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}
func (s service) ListCostExplain() (int, []modelCost.CostExplanation, error) {
	repo, err := s.aboutRepo.ListCostExplanation()
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("Any ID Cost Explain not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}
func (s service) CreateCostExplain(params data.Params) (int, *modelCost.CostExplanation, error) {
	title := params.GetString("title")
	description := params.GetString("description")
	model := &modelCost.CostExplanation{
		Title:       title,
		Description: description,
	}
	repo, err := s.aboutRepo.CreateCostExplanation(model)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, repo, nil
}
func (s service) UpdateCostExplainByID(id int64, params data.Params) (int, error) {
	count, err := s.aboutRepo.UpdateCostExplanationByID(id, params)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if count == 0 {
		return http.StatusNotFound, errors.New("0 rows affected when updating Cost Explain")
	}

	return http.StatusOK, nil
}
func (s service) DeleteCostExplainByID(id int64) (httpStatus int, err error) {
	count, err := s.aboutRepo.DeleteCostExplanationByID(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if count == 0 {
		return http.StatusNotFound, errors.New("0 rows affected when updating FAQ")
	}
	return http.StatusOK, nil
}

//Faq
func (s service) GetFaqByID(id int64) (int, *modelFaq.Faq, error) {
	repo, err := s.aboutRepo.GetFaqID(id)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("ID FAQ not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}
func (s service) ListFaq() (int, []modelFaq.Faq, error) {
	repo, err := s.aboutRepo.ListFaq()
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("Any ID FAQ not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}
func (s service) CreateFaq(params data.Params) (int, *modelFaq.Faq, error) {
	title := params.GetString("title")
	idOrder := params.GetInt64("id_order")
	model := &modelFaq.Faq{
		Title:   title,
		IDOrder: idOrder,
	}
	checkIDOrderIfExist, _ := s.aboutRepo.GetFaqIDOrder(idOrder)
	if checkIDOrderIfExist.IDOrder == idOrder {
		return http.StatusFound, nil, errors.New("ID Order is found, please change id_order or delete the first one")
	}
	repo, err := s.aboutRepo.CreateFaq(model)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, repo, nil
}
func (s service) UpdateFaqByID(id int64, params data.Params) (int, error) {
	idOrder := params.GetInt64("id_order")
	if idOrder != 0 {
		checkIDOrderIfExist, err := s.aboutRepo.GetFaqIDOrder(idOrder)
		if checkIDOrderIfExist.IDOrder == idOrder {
			return http.StatusFound, errors.New("ID Order is found, please change id_order or delete the first one")
		}
		if err != nil {
			return http.StatusInternalServerError, err
		}
	}

	count, err := s.aboutRepo.UpdateFaqByID(id, params)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if count == 0 {
		return http.StatusNotFound, errors.New("0 rows affected when updating FAQ")
	}

	return http.StatusOK, nil
}
func (s service) DeleteFaqByID(id int64) (httpStatus int, err error) {
	count, err := s.aboutRepo.DeleteFaqByID(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if count == 0 {
		return http.StatusNotFound, errors.New("ID FAQ is not found")
	}
	modelLast, err := s.aboutRepo.GetFaqTitleLastIDOrder(id)
	if modelLast.IDOrder == 0 {
		return http.StatusNotFound, nil
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}
	modelFirst, err := s.aboutRepo.GetFaqTitleFirstIDOrder(id)
	if modelFirst.IDOrder == 0 {
		// return statusOK because there is no ID FAQ > id
		// no need to continue logic below
		return http.StatusOK, nil
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}

	decrementNumber := (modelFirst.IDOrder - modelLast.IDOrder) - int64(1)
	count, err = s.aboutRepo.DecrementIDOrderByDecrementNumber(decrementNumber, modelFirst.IDOrder)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if count == 0 {
		return http.StatusNotFound, errors.New("Error when decrementing id order")
	}
	return http.StatusOK, nil
}

//Faq Title
func (s service) GetFaqTitleByID(id int64) (int, *modelFaq.FaqTitle, error) {
	repo, err := s.aboutRepo.GetFaqTitleID(id)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("ID FAQ title not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}
func (s service) ListFaqTitle() (int, []modelFaq.FaqTitle, error) {
	repo, err := s.aboutRepo.ListFaqTitle()
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("Any ID FAQ title not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}
func (s service) ListFaqTitleByIDFaq(idFaq int64) (int, []modelFaq.FaqTitle, error) {
	repo, err := s.aboutRepo.ListFaqTitleByIDFaq(idFaq)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("ID FAQ title not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}
func (s service) CreateFaqTitle(params data.Params) (int, *modelFaq.FaqTitle, error) {
	idFaq := params.GetInt64("id_faq")
	title := params.GetString("title")
	description := params.GetString("description")
	idOrder := params.GetInt64("id_order")
	model := &modelFaq.FaqTitle{
		IDFaq:       idFaq,
		Title:       title,
		Description: description,
		IDOrder:     idOrder,
	}
	checkIdFAQIfExist, err := s.aboutRepo.GetFaqID(model.IDFaq)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("ID FAQ is not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, errors.New("There is error related to server")
	}
	if checkIdFAQIfExist.ID == 0 {
		return http.StatusNotFound, nil, errors.New("ID FAQ is not found")
	}
	repo, err := s.aboutRepo.CreateFaqTitle(model)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusFound, repo, nil
}
func (s service) UpdateFaqTitleByID(id int64, params data.Params) (int, error) {
	repo, err := s.aboutRepo.GetFaqTitleID(id)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, errors.New("ID FAQ Title is not found")
	}
	if err != nil {
		return http.StatusInternalServerError, errors.New("ID FAQ Title is not found")
	}

	if repo.IDOrder != params.GetInt64("id_order") {
		getFaqTitle, err := s.aboutRepo.GetFaqTitleIDOrder(params.GetInt64("id_order"))
		if err == sql.ErrNoRows {
			return http.StatusNotFound, errors.New("ID Order FAQ Title is not found")
		}
		if err != nil {
			return http.StatusInternalServerError, errors.New("ID Order FAQ Title is not found")
		}

		count, err := s.aboutRepo.UpdateFaqTitleIDOrderByID(getFaqTitle.ID, repo.IDOrder)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		if count == 0 {
			return http.StatusNotFound, errors.New("ID FAQ Title is not found")
		}
	}

	count, err := s.aboutRepo.UpdateFaqTitleByID(id, params)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if count == 0 {
		return http.StatusNotFound, errors.New("ID FAQ Title is not found")
	}

	return http.StatusOK, nil
}
func (s service) DeleteFaqTitleByID(id int64) (httpStatus int, err error) {
	count, err := s.aboutRepo.DeleteFaqTitleByID(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if count == 0 {
		return http.StatusNotFound, errors.New("ID FAQ is not found")
	}
	return http.StatusOK, nil
}

// Cost Explanation Page
func (s service) GetCostExplanationPage() (int, *resp.CostExplanationPage, error) {
	statusHttp, service, err := s.ListCost()
	if err == sql.ErrNoRows {
		return statusHttp, nil, errors.New("Cost ID Loan Option not found")
	}
	if err != nil {
		return statusHttp, nil, err
	}
	respCostExplanationPage := resp.CostExplanationPage{}
	for _, dataService := range service {
		respCost := resp.Cost{}
		respCost.ID = dataService.ID
		respCost.LoanOption = dataService.LoanOption
		respCost.IDLoanOption = dataService.IDLoanOption
		respCost.Interest = dataService.Interest
		respCost.AdminFee = dataService.AdminFee
		respCost.FinePerDay = dataService.FinePerDay
		respCost.Description = dataService.Description
		respCost.IsVisible = dataService.IsVisible

		respCostExplanationPage.Cost = append(respCostExplanationPage.Cost, respCost)
	}

	statusHttpExplain, serviceExplain, err := s.ListCostExplain()
	if err == sql.ErrNoRows {
		return statusHttpExplain, nil, errors.New("Any Cost Explanation ID not found")
	}
	if err != nil {
		return statusHttpExplain, nil, err
	}

	for _, dataServiceExplain := range serviceExplain {
		respCostexplain := resp.CostExplain{}
		respCostexplain.ID = dataServiceExplain.ID
		respCostexplain.Title = dataServiceExplain.Title
		respCostexplain.Description = dataServiceExplain.Description

		respCostExplanationPage.CostExplain = append(respCostExplanationPage.CostExplain, respCostexplain)
	}
	return http.StatusOK, &respCostExplanationPage, nil
}

// FAQ Page
func (s service) GetFaqPage() (int, *resp.FAQPage, error) {
	statusHttp, service, err := s.ListFaq()
	if err == sql.ErrNoRows {
		return statusHttp, nil, errors.New("ID FAQ not found")
	}
	if err != nil {
		return statusHttp, nil, err
	}

	respFaqPage := resp.FAQPage{}
	for _, dataService := range service {
		respFaq := resp.FAQ{}
		respFaq.Id = dataService.ID
		respFaq.Title = dataService.Title
		respFaq.IdOrder = dataService.IDOrder

		statusHttp, serviceTitle, err := s.ListFaqTitleByIDFaq(dataService.ID)
		if err == sql.ErrNoRows {
			return statusHttp, nil, errors.New("ID FAQ not found")
		}
		if err != nil {
			return statusHttp, nil, err
		}
		for _, dataServiceTitle := range serviceTitle {
			respFaqTitle := resp.FAQTitle{}
			respFaqTitle.ID = dataServiceTitle.ID
			respFaqTitle.IDFaq = dataServiceTitle.IDFaq
			respFaqTitle.Title = dataServiceTitle.Title
			respFaqTitle.Description = dataServiceTitle.Description
			respFaqTitle.IdOrder = dataServiceTitle.IDOrder

			respFaq.FAQTitle = append(respFaq.FAQTitle, respFaqTitle)
		}
		respFaqPage.FAQ = append(respFaqPage.FAQ, respFaq)
	}

	return http.StatusOK, &respFaqPage, nil
}
