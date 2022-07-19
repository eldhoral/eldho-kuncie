package service

import (
	"database/sql"
	"errors"
	"net/http"
	"os"

	modelLanding "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/offer/domain/landing"
	modelTnc "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/offer/domain/tnc"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/offer/presenter/resp"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/offer/repository"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/data"
)

// NewService creates new user service
func NewService(repo repository.Repository) Service {
	return &service{
		offerRepo: repo,
	}
}

type service struct {
	offerRepo repository.Repository
}

//Loan limit
func (s service) GetLoanLimit() (int, *modelLanding.LoanLimit, error) {
	repo, err := s.offerRepo.GetLoanLimit()
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("Loan limit not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}
func (s service) UpdateLoanLimit(limit string) (int, error) {
	_, err := s.offerRepo.UpdateLoanLimit(limit)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, errors.New("Loan limit ID not found")
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

//Benefit
func (s service) GetBenefitByID(id int64) (int, *modelLanding.Benefit, error) {
	repo, err := s.offerRepo.GetBenefitByID(id)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("ID Benefit not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}
func (s service) ListBenefit() (int, []modelLanding.Benefit, error) {
	repo, err := s.offerRepo.ListBenefit()
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("Any ID Benefit not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	benefit := []modelLanding.Benefit{}
	for _, dataBenefit := range repo {
		dataBenefit.Image = os.Getenv("BASE_URL_IMAGE") + dataBenefit.Image

		benefit = append(benefit, dataBenefit)
	}

	return http.StatusOK, benefit, nil
}
func (s service) CreateBenefit(title string, description string, path string) (int, *modelLanding.Benefit, error) {
	model := &modelLanding.Benefit{
		Title:       title,
		Description: description,
		Image:       path,
	}
	repo, err := s.offerRepo.CreateBenefit(model)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, repo, nil
}
func (s service) UpdateBenefitByID(id int64, params data.Params, path string) (int, error) {
	_, err := s.offerRepo.UpdateBenefitByID(id, params, path)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, errors.New("ID Benefit not found")
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
func (s service) DeleteBenefitByID(id int64) (int, error) {
	status, err := s.offerRepo.DeleteBenefitByID(id)
	if err == sql.ErrNoRows {
		return status, errors.New("Benefit ID not found")
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

//Loan method
func (s service) GetLoanMethodByID(id int64) (int, *modelLanding.LoanMethod, error) {
	repo, err := s.offerRepo.GetLoanMethodByID(id)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("ID Loan method not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}
func (s service) ListLoanMethod() (int, []modelLanding.LoanMethod, error) {
	repo, err := s.offerRepo.ListLoanMethod()
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("Any ID Loan method not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}
func (s service) CreateLoanMethod(title string, description string) (int, *modelLanding.LoanMethod, error) {
	model := &modelLanding.LoanMethod{
		Title:       title,
		Description: description,
	}
	repo, err := s.offerRepo.CreateLoanMethod(model)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, repo, nil
}
func (s service) UpdateLoanMethodByID(id int64, title string, description string) (int, error) {
	_, err := s.offerRepo.UpdateLoanMethodByID(id, title, description)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, errors.New("Loan method ID not found")
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
func (s service) DeleteLoanMethodByID(id int64) (int, error) {
	status, err := s.offerRepo.DeleteLoanMethodByID(id)
	if err == sql.ErrNoRows {
		return status, errors.New("Loan method ID not found")
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

//Tnc
func (s service) UpdateTncMobile(params data.Params) (int, error) {
	_, err := s.offerRepo.UpdateTncMobile(params)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, errors.New("Tnc Row not found")
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
func (s service) GetTncPageMobile() (int, *modelTnc.TncMobile, error) {
	repo, err := s.offerRepo.GetTncMobile()
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("Tnc Mobile not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}
func (s service) GetTncPageMobileByID(id int64) (int, *modelTnc.TncMobile, error) {
	repo, err := s.offerRepo.GetTncMobileByID(id)
	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, errors.New("Tnc Mobile ID not found")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, repo, nil
}

//Paylater Offer Page - Landing Page
func (s service) GetLandingPage() (int, *resp.LandingPage, error) {
	httpStatus, repoLoanLimit, err := s.GetLoanLimit()
	if err != nil {
		return httpStatus, nil, err
	}
	respLoanLimit := resp.LoanLimit{
		ID:    repoLoanLimit.ID,
		Limit: repoLoanLimit.Limit,
	}

	respLandingPage := resp.LandingPage{
		LoanLimit: &respLoanLimit,
	}

	httpStatus, repoBenefit, err := s.ListBenefit()
	if err != nil {
		return httpStatus, nil, err
	}
	for _, dataBenefit := range repoBenefit {
		respBenefit := resp.Benefit{
			ID:          dataBenefit.ID,
			Title:       dataBenefit.Title,
			Description: dataBenefit.Description,
			Image:       dataBenefit.Image,
		}
		respLandingPage.Benefit = append(respLandingPage.Benefit, respBenefit)
	}

	httpStatus, repoLoanMethod, err := s.ListLoanMethod()
	if err != nil {
		return httpStatus, nil, err
	}
	for _, dataLoanMethod := range repoLoanMethod {
		respLoanMethod := resp.LoanMethod{
			ID:          dataLoanMethod.ID,
			Title:       dataLoanMethod.Title,
			Description: dataLoanMethod.Description,
		}
		respLandingPage.LoanMethod = append(respLandingPage.LoanMethod, respLoanMethod)
	}

	return http.StatusOK, &respLandingPage, nil
}
