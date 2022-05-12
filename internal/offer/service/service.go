package service

import (
	"database/sql"
	"errors"
	"strconv"

	modelLanding "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/offer/domain/landing"
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
func (s service) GetLoanLimitByID(id int64) (*modelLanding.LoanLimit, error) {
	return s.offerRepo.GetLoanLimitByID(id)
}

func (s service) GetLoanLimit() (*modelLanding.LoanLimit, error) {
	return s.offerRepo.GetLoanLimit()
}

func (s service) CreateLoanLimit(limit string) (*modelLanding.LoanLimit, error) {
	limitLoan, _ := strconv.Atoi(limit)
	model := &modelLanding.LoanLimit{
		Limit: float64(limitLoan),
	}
	repo, err := s.offerRepo.CreateLoanLimit(model)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (s service) UpdateLoanLimit(limit string) error {
	_, err := s.offerRepo.UpdateLoanLimit(limit)
	if err == sql.ErrNoRows {
		return errors.New("Loan limit ID not found")
	}
	if err != nil {
		return err
	}

	return nil
}

func (s service) DeleteLoanLimit() error {
	err := s.offerRepo.DeleteLoanLimit()
	if err == sql.ErrNoRows {
		return errors.New("Loan limit not found")
	}
	if err != nil {
		return err
	}

	return nil
}

//Benefit
func (s service) GetBenefitByID(id int64) (*modelLanding.Benefit, error) {
	return s.offerRepo.GetBenefitByID(id)
}

func (s service) ListBenefit() ([]modelLanding.Benefit, error) {
	return s.offerRepo.ListBenefit()
}

func (s service) CreateBenefit(title string, description string, path string) (*modelLanding.Benefit, error) {
	model := &modelLanding.Benefit{
		Title:       title,
		Description: description,
		Image:       path,
	}
	repo, err := s.offerRepo.CreateBenefit(model)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (s service) UpdateBenefitByID(id int64, params data.Params, path string) error {
	_, err := s.offerRepo.UpdateBenefitByID(id, params, path)
	if err == sql.ErrNoRows {
		return errors.New("Benefit ID not found")
	}
	if err != nil {
		return err
	}

	return nil
}

func (s service) DeleteBenefitByID(id int64) error {
	err := s.offerRepo.DeleteBenefitByID(id)
	if err == sql.ErrNoRows {
		return errors.New("Benefit ID not found")
	}
	if err != nil {
		return err
	}
	return nil
}

//Loan method
func (s service) GetLoanMethodByID(id int64) (*modelLanding.LoanMethod, error) {
	return s.offerRepo.GetLoanMethodByID(id)
}

func (s service) ListLoanMethod() ([]modelLanding.LoanMethod, error) {
	return s.offerRepo.ListLoanMethod()
}

func (s service) CreateLoanMethod(title string, description string) (*modelLanding.LoanMethod, error) {
	model := &modelLanding.LoanMethod{
		Title:       title,
		Description: description,
	}
	repo, err := s.offerRepo.CreateLoanMethod(model)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (s service) UpdateLoanMethodByID(id int64, title string, description string) error {
	_, err := s.offerRepo.UpdateLoanMethodByID(id, title, description)
	if err == sql.ErrNoRows {
		return errors.New("Loan method ID not found")
	}
	if err != nil {
		return err
	}

	return nil
}

func (s service) DeleteLoanMethodByID(id int64) error {
	err := s.offerRepo.DeleteLoanMethodByID(id)
	if err == sql.ErrNoRows {
		return errors.New("Loan method ID not found")
	}
	if err != nil {
		return err
	}
	return nil
}

//Paylater Offer Page - Landing Page
func (s service) GetLandingPage() (*resp.LandingPage, error) {
	repoLoanLimit, err := s.offerRepo.GetLoanLimit()
	if err != nil {
		return nil, err
	}
	respLoanLimit := resp.LoanLimit{
		ID:    repoLoanLimit.ID,
		Limit: repoLoanLimit.Limit,
	}

	respLandingPage := resp.LandingPage{
		LoanLimit: &respLoanLimit,
	}

	repoBenefit, err := s.offerRepo.ListBenefit()
	if err != nil {
		return nil, err
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

	repoLoanMethod, err := s.offerRepo.ListLoanMethod()
	if err != nil {
		return nil, err
	}
	for _, dataLoanMethod := range repoLoanMethod {
		respLoanMethod := resp.LoanMethod{
			ID:          dataLoanMethod.ID,
			Title:       dataLoanMethod.Title,
			Description: dataLoanMethod.Description,
		}
		respLandingPage.LoanMethod = append(respLandingPage.LoanMethod, respLoanMethod)
	}

	return &respLandingPage, nil
}
