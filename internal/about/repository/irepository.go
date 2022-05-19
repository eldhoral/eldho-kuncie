package repository

import (
	modelCost "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/about/domain/cost"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/data"
)

type Repository interface {
	// Cost
	GetCostByID(id int64) (*modelCost.Cost, error)
	ListCost() ([]modelCost.Cost, error)
	CreateCost(c *modelCost.Cost) (*modelCost.Cost, error)
	UpdateCostByID(id int64, params data.Params) (int64, error)
	DeleteCostByID(id int64) (httpStatus int, err error)

	// Cost Explanation
	GetCostExplanationByID(id int64) (*modelCost.CostExplanation, error)
	ListCostExplanation() ([]modelCost.CostExplanation, error)
	CreateCostExplanation(ce *modelCost.CostExplanation) (*modelCost.CostExplanation, error)
	UpdateCostExplanationByID(id int64, params data.Params) (int64, error)
	DeleteCostExplanationByID(id int64) (httpStatus int, err error)
}
