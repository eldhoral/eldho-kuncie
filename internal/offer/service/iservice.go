package service

import (
	modelLanding "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/offer/domain/landing"
	modelTnc "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/offer/domain/tnc"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/offer/presenter/resp"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/data"
)

// Cost service
type Service interface {
	//CMS
	//Loan limit
	GetLoanLimitByID(id int64) (int, *modelLanding.LoanLimit, error)
	GetLoanLimit() (int, *modelLanding.LoanLimit, error)
	CreateLoanLimit(limit string) (int, *modelLanding.LoanLimit, error)
	UpdateLoanLimit(limit string) (int, error)
	DeleteLoanLimit() (httpStatus int, err error)

	//Benefit
	GetBenefitByID(id int64) (int, *modelLanding.Benefit, error)
	ListBenefit() (int, []modelLanding.Benefit, error)
	CreateBenefit(title string, description string, path string) (int, *modelLanding.Benefit, error)
	UpdateBenefitByID(id int64, params data.Params, path string) (int, error)
	DeleteBenefitByID(id int64) (httpStatus int, err error)

	//Loan method
	GetLoanMethodByID(id int64) (int, *modelLanding.LoanMethod, error)
	ListLoanMethod() (int, []modelLanding.LoanMethod, error)
	CreateLoanMethod(title string, description string) (int, *modelLanding.LoanMethod, error)
	UpdateLoanMethodByID(id int64, title string, description string) (int, error)
	DeleteLoanMethodByID(id int64) (httpStatus int, err error)

	//Tnc
	GetTncPageMobile() (int, *modelTnc.TncMobile, error)
	UpdateTncMobile(params data.Params) (int, error)

	//Page
	//Paylater Offer Page - Landing Page
	GetLandingPage() (int, *resp.LandingPage, error)
}
