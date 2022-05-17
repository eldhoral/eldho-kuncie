package repository

import (
	modelLanding "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/offer/domain/landing"
	modelTnc "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/offer/domain/tnc"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/data"
)

type Repository interface {
	//Loan limit
	GetLoanLimitByID(id int64) (*modelLanding.LoanLimit, error)
	GetLoanLimit() (*modelLanding.LoanLimit, error)
	CreateLoanLimit(ll *modelLanding.LoanLimit) (*modelLanding.LoanLimit, error)
	UpdateLoanLimit(limit string) (int64, error)
	DeleteLoanLimit() error

	//Benefit
	GetBenefitByID(id int64) (*modelLanding.Benefit, error)
	ListBenefit() ([]modelLanding.Benefit, error)
	CreateBenefit(b *modelLanding.Benefit) (*modelLanding.Benefit, error)
	UpdateBenefitByID(id int64, params data.Params, path string) (int64, error)
	DeleteBenefitByID(id int64) error

	//Loan method
	GetLoanMethodByID(id int64) (*modelLanding.LoanMethod, error)
	ListLoanMethod() ([]modelLanding.LoanMethod, error)
	CreateLoanMethod(lm *modelLanding.LoanMethod) (*modelLanding.LoanMethod, error)
	UpdateLoanMethodByID(id int64, title string, description string) (int64, error)
	DeleteLoanMethodByID(id int64) error

	//Tnc
	GetTncByID(id int64) (*modelTnc.Tnc, error)
	ListTnc() ([]modelTnc.Tnc, error)
	CreateTnc(t *modelTnc.Tnc) (*modelTnc.Tnc, error)
	UpdateTncByID(id int64, params data.Params) (int64, error)
	DeleteTncByID(id int64) error

	//Tnc title
	GetTncTitleByID(id int64) (*modelTnc.TncTitle, error)
	ListTncTitle() ([]modelTnc.TncTitle, error)
	ListTncTitleByID(idTnc int64) ([]modelTnc.TncTitle, error)
	CreateTncTitle(t *modelTnc.TncTitle) (*modelTnc.TncTitle, error)
	UpdateTncTitleByID(id int64, params data.Params) (int64, error)
	DeleteTncTitleByID(id int64) error

	//Tnc subtitle
	GetTncSubtitleByID(id int64) (*modelTnc.TncSubtitle, error)
	ListTncSubtitle() ([]modelTnc.TncSubtitle, error)
	ListTncSubtitleByID(idTncTitle int64) ([]*modelTnc.TncSubtitle, error)
	CreateTncSubtitle(t *modelTnc.TncSubtitle) (*modelTnc.TncSubtitle, error)
	UpdateTncSubtitleByID(id int64, params data.Params) (int64, error)
	DeleteTncSubtitleByID(id int64) error

	//Tnc explain
	GetTncExplainByID(id int64) (*modelTnc.TncExplain, error)
	ListTncExplain() ([]modelTnc.TncExplain, error)
	ListTncExplainByID(idTnc int64, idTncTitle int64) ([]modelTnc.TncExplain, error)
	ListTncExplainByIDWithSubtitle(idTnc int64, idTncTitle int64, idTncSubtitle int64) ([]*modelTnc.TncExplain, error)
	CreateTncExplain(t *modelTnc.TncExplain) (*modelTnc.TncExplain, error)
	UpdateTncExplainByID(id int64, params data.Params) (int64, error)
	DeleteTncExplainByID(id int64) error
}
