package service

import (
	"database/sql"
	"errors"
	"net/http"

	modelCost "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/about/domain/cost"
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
func (s service) CreateCost(params data.Params) (int, *modelCost.Cost, error) {
	loanOption := params.GetString("loan_option")
	idLoanOption := params.GetInt("id_loan_option")
	interest := params.GetString("interest")
	adminFee := params.GetString("admin_fee")
	finePerDay := params.GetString("fine_per_day")
	description := params.GetString("description")
	model := &modelCost.Cost{
		LoanOption:   loanOption,
		IDLoanOption: idLoanOption,
		Interest:     interest,
		AdminFee:     adminFee,
		FinePerDay:   finePerDay,
		Description:  description,
	}
	repo, err := s.aboutRepo.CreateCost(model)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, repo, nil
}
func (s service) UpdateCostByID(id int64, params data.Params) (int, error) {
	_, err := s.aboutRepo.UpdateCostByID(id, params)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, errors.New("ID Cost not found")
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
func (s service) DeleteCostByID(id int64) (httpStatus int, err error) {
	status, err := s.aboutRepo.DeleteCostByID(id)
	if err == sql.ErrNoRows {
		return status, errors.New("Cost ID not found")
	}
	if err != nil {
		return http.StatusInternalServerError, err
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
	_, err := s.aboutRepo.UpdateCostExplanationByID(id, params)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, errors.New("ID Cost Explain not found")
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
func (s service) DeleteCostExplainByID(id int64) (httpStatus int, err error) {
	status, err := s.aboutRepo.DeleteCostExplanationByID(id)
	if err == sql.ErrNoRows {
		return status, errors.New("Cost ID Explain not found")
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
