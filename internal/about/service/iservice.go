package service

import (
	modelCost "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/about/domain/cost"
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

	//Page
	//About Paylater - Penjelasan Biaya
	GetCostExplanationPage(params data.Params) (int, *resp.CostExplanationPage, error)
}
