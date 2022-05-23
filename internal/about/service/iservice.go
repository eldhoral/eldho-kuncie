package service

import (
	modelCost "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/about/domain/cost"
	modelFaq "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/about/domain/faq"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/about/presenter/resp"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/data"
)

type Service interface {
	//CMS
	//Cost
	GetCostByID(id int64) (int, *modelCost.Cost, error)
	ListCost() (int, []modelCost.Cost, error)
	CreateCost(params data.Params) (int, *modelCost.Cost, error)
	UpdateCostByID(id int64, params data.Params) (int, error)
	DeleteCostByID(id int64) (httpStatus int, err error)

	//Cost Explain
	GetCostExplainByID(id int64) (int, *modelCost.CostExplanation, error)
	ListCostExplain() (int, []modelCost.CostExplanation, error)
	CreateCostExplain(params data.Params) (int, *modelCost.CostExplanation, error)
	UpdateCostExplainByID(id int64, params data.Params) (int, error)
	DeleteCostExplainByID(id int64) (httpStatus int, err error)

	//Faq
	GetFaqByID(id int64) (int, *modelFaq.Faq, error)
	ListFaq() (int, []modelFaq.Faq, error)
	CreateFaq(params data.Params) (int, *modelFaq.Faq, error)
	UpdateFaqByID(id int64, params data.Params) (int, error)
	DeleteFaqByID(id int64) (httpStatus int, err error)

	//Faq Title
	GetFaqTitleByID(id int64) (int, *modelFaq.FaqTitle, error)
	ListFaqTitle() (int, []modelFaq.FaqTitle, error)
	ListFaqTitleByIDFaq(idFaq int64) (int, []modelFaq.FaqTitle, error)
	CreateFaqTitle(params data.Params) (int, *modelFaq.FaqTitle, error)
	UpdateFaqTitleByID(id int64, params data.Params) (int, error)
	DeleteFaqTitleByID(id int64) (httpStatus int, err error)

	//Page
	//About Paylater - Penjelasan Biaya
	GetCostExplanationPage(params data.Params) (int, *resp.CostExplanationPage, error)

	//About FAQ
	GetFaqPage() (int, *resp.FAQPage, error)
}
