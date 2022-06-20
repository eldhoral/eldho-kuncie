package repository

import (
	modelLanding "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/offer/domain/landing"
	modelTnc "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/offer/domain/tnc"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/data"
)

type Repository interface {
	//Loan limit
	GetLoanLimit() (*modelLanding.LoanLimit, error)
	UpdateLoanLimit(limit string) (int64, error)

	//Benefit
	GetBenefitByID(id int64) (*modelLanding.Benefit, error)
	ListBenefit() ([]modelLanding.Benefit, error)
	CreateBenefit(b *modelLanding.Benefit) (*modelLanding.Benefit, error)
	UpdateBenefitByID(id int64, params data.Params, path string) (int64, error)
	DeleteBenefitByID(id int64) (httpStatus int, err error)

	//Loan method
	GetLoanMethodByID(id int64) (*modelLanding.LoanMethod, error)
	ListLoanMethod() ([]modelLanding.LoanMethod, error)
	CreateLoanMethod(lm *modelLanding.LoanMethod) (*modelLanding.LoanMethod, error)
	UpdateLoanMethodByID(id int64, title string, description string) (int64, error)
	DeleteLoanMethodByID(id int64) (httpStatus int, err error)

	//Tnc
	GetTncMobile() (*modelTnc.TncMobile, error)
	UpdateTncMobile(params data.Params) (int64, error)
}
