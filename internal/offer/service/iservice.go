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
	GetTncByID(id int64) (int, *modelTnc.Tnc, error)
	GetTncPageMobile() (int, *modelTnc.TncMobile, error)
	ListTnc() (int, []modelTnc.Tnc, error)
	CreateTnc(title string) (int, *modelTnc.Tnc, error)
	UpdateTncByID(id int64, params data.Params) (int, error)
	DeleteTncByID(id int64) (httpStatus int, err error)

	//Tnc title
	GetTncTitleByID(id int64) (int, *modelTnc.TncTitle, error)
	ListTncTitle() (int, []modelTnc.TncTitle, error)
	CreateTncTitle(idTnc int64, title string) (int, *modelTnc.TncTitle, error)
	UpdateTncTitleByID(id int64, params data.Params) (int, error)
	DeleteTncTitleByID(id int64) (httpStatus int, err error)

	//Tnc subtitle
	GetTncSubtitleByID(id int64) (int, *modelTnc.TncSubtitle, error)
	ListTncSubtitle() (int, []modelTnc.TncSubtitle, error)
	CreateTncSubtitle(idTncTitle int64, subtitle string) (int, *modelTnc.TncSubtitle, error)
	UpdateTncSubtitleByID(id int64, params data.Params) (int, error)
	DeleteTncSubtitleByID(id int64) (httpStatus int, err error)

	//Tnc explain
	GetTncExplainByID(id int64) (int, *modelTnc.TncExplain, error)
	ListTncExplain() (int, []modelTnc.TncExplain, error)
	CreateTncExplain(idTnc int64, idTncTitle int64, idTncSubtitle *int64, description string) (int, *modelTnc.TncExplain, error)
	UpdateTncExplainByID(id int64, params data.Params) (int, error)
	DeleteTncExplainByID(id int64) (httpStatus int, err error)

	//Page
	//Paylater Offer Page - Landing Page
	GetLandingPage() (int, *resp.LandingPage, error)

	//Paylater Offer Page - Tnc Page
	GetTncPage() (int, []*resp.TncPage, error)
}
