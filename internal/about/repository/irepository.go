package repository

import (
	modelCost "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/about/domain/cost"
	modelFaq "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/about/domain/faq"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/data"
)

type Repository interface {
	// Cost
	GetCostByID(id int64) (*modelCost.Cost, error)
	ListCost() ([]modelCost.Cost, error)
	ListCostByIDLoanOption(idLoanOption int64) ([]modelCost.Cost, error)
	UpdateCostByID(id int64, params data.Params) (int64, error)

	// Cost Explanation
	GetCostExplanationByID(id int64) (*modelCost.CostExplanation, error)
	ListCostExplanation() ([]modelCost.CostExplanation, error)
	CreateCostExplanation(ce *modelCost.CostExplanation) (*modelCost.CostExplanation, error)
	UpdateCostExplanationByID(id int64, params data.Params) (int64, error)
	DeleteCostExplanationByID(id int64) (count int64, err error)

	// FAQ
	GetFaqID(id int64) (*modelFaq.Faq, error)
	GetFaqIDOrder(idOrder int64) (*modelFaq.Faq, error)
	ListFaq() ([]modelFaq.Faq, error)
	AutoIncrementIDOrder(idOrder int64) (int64, error)
	AutoDecrementIDOrder(idOrder int64) (int64, error)
	DecrementIDOrderByDecrementNumber(decrementNumber int64, idOrder int64) (int64, error)
	CreateFaq(ce *modelFaq.Faq) (*modelFaq.Faq, error)
	UpdateFaqByID(id int64, params data.Params) (int64, error)
	DeleteFaqByID(id int64) (count int64, err error)

	// FAQ Title
	GetFaqTitleID(id int64) (*modelFaq.FaqTitle, error)
	GetFaqTitleIDOrder(idOrder int64) (*modelFaq.FaqTitle, error)
	GetFaqTitleLastIDOrder(idOrder int64) (*modelFaq.FAQTitleResponseMAX, error)
	GetFaqTitleFirstIDOrder(idOrder int64) (*modelFaq.FAQTitleResponseMIN, error)
	ListFaqTitle() ([]modelFaq.FaqTitle, error)
	ListFaqTitleByIDFaq(idFaq int64) ([]modelFaq.FaqTitle, error)
	CreateFaqTitle(ce *modelFaq.FaqTitle) (*modelFaq.FaqTitle, error)
	UpdateFaqTitleByID(id int64, params data.Params) (int64, error)
	DeleteFaqTitleByID(id int64) (count int64, err error)
	DeleteFaqTitleByIDFAQ(idFaq int64) (count int64, err error)
}
