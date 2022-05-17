package service

import (
	modelLanding "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/offer/domain/landing"
	modelTnc "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/offer/domain/tnc"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/offer/presenter/resp"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/data"
)

type Service interface {
	//CMS
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

	//Tnc
	GetTncByID(id int64) (*modelTnc.Tnc, error)
	ListTnc() ([]modelTnc.Tnc, error)
	CreateTnc(title string) (*modelTnc.Tnc, error)
	UpdateTncByID(id int64, params data.Params) error
	DeleteTncByID(id int64) error

	//Tnc title
	GetTncTitleByID(id int64) (*modelTnc.TncTitle, error)
	ListTncTitle() ([]modelTnc.TncTitle, error)
	CreateTncTitle(idTnc int64, title string) (*modelTnc.TncTitle, error)
	UpdateTncTitleByID(id int64, params data.Params) error
	DeleteTncTitleByID(id int64) error

	//Tnc subtitle
	GetTncSubtitleByID(id int64) (*modelTnc.TncSubtitle, error)
	ListTncSubtitle() ([]modelTnc.TncSubtitle, error)
	CreateTncSubtitle(idTncTitle int64, subtitle string) (*modelTnc.TncSubtitle, error)
	UpdateTncSubtitleByID(id int64, params data.Params) error
	DeleteTncSubtitleByID(id int64) error

	//Tnc explain
	GetTncExplainByID(id int64) (*modelTnc.TncExplain, error)
	ListTncExplain() ([]modelTnc.TncExplain, error)
	CreateTncExplain(idTnc int64, idTncTitle int64, idTncSubtitle *int64, description string) (*modelTnc.TncExplain, error)
	UpdateTncExplainByID(id int64, params data.Params) error
	DeleteTncExplainByID(id int64) error

	//Page
	//Paylater Offer Page - Landing Page
	GetLandingPage() (*resp.LandingPage, error)

	//Paylater Offer Page - Tnc Page
	GetTncPage() ([]*resp.TncPage, error)
}
