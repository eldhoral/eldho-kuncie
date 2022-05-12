package service

import (
	modelLanding "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/offer/domain/landing"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/offer/presenter/resp"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/data"
)

type Service interface {
	//Loan limit
	GetLoanLimitByID(id int64) (*modelLanding.LoanLimit, error)
	GetLoanLimit() (*modelLanding.LoanLimit, error)
	CreateLoanLimit(limit string) (*modelLanding.LoanLimit, error)
	UpdateLoanLimit(limit string) error
	DeleteLoanLimit() error

	//Benefit
	GetBenefitByID(id int64) (*modelLanding.Benefit, error)
	ListBenefit() ([]modelLanding.Benefit, error)
	CreateBenefit(title string, description string, path string) (*modelLanding.Benefit, error)
	UpdateBenefitByID(id int64, params data.Params, path string) error
	DeleteBenefitByID(id int64) error

	//Loan method
	GetLoanMethodByID(id int64) (*modelLanding.LoanMethod, error)
	ListLoanMethod() ([]modelLanding.LoanMethod, error)
	CreateLoanMethod(title string, description string) (*modelLanding.LoanMethod, error)
	UpdateLoanMethodByID(id int64, title string, description string) error
	DeleteLoanMethodByID(id int64) error

	//Paylater Offer Page - Landing Page
	GetLandingPage() (*resp.LandingPage, error)
}
