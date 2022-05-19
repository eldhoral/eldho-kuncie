package repository

import (
	"net/http"
	"time"

	modelCost "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/about/domain/cost"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/data"

	"github.com/jmoiron/sqlx"
)

// NewRepository creates new repository
func NewRepository(db *sqlx.DB) Repository {
	return &repo{db: db}
}

type repo struct {
	db *sqlx.DB
}

//Cost
func (r repo) GetCostByID(id int64) (*modelCost.Cost, error) {
	cost := &modelCost.Cost{}
	err := r.db.Get(cost, "SELECT * FROM tbl_cost WHERE id = ?", id)
	return cost, err
}
func (r repo) ListCost() ([]modelCost.Cost, error) {
	cost := []modelCost.Cost{}
	err := r.db.Select(&cost, "SELECT * FROM tbl_cost")
	return cost, err
}
func (r repo) CreateCost(c *modelCost.Cost) (*modelCost.Cost, error) {
	arg := map[string]interface{}{
		"loan_option":    c.LoanOption,
		"id_loan_option": c.IDLoanOption,
		"interest":       c.Interest,
		"admin_fee":      c.AdminFee,
		"fine_per_day":   c.FinePerDay,
		"description":    c.Description,
	}

	query := `INSERT INTO tbl_cost
		SET loan_option = :loan_option, id_loan_option = :id_loan_option, interest = :interest,
		admin_fee = :admin_fee, fine_per_day = :fine_per_day, description = :description`

	cost, err := r.db.NamedExec(query, arg)
	if err != nil {
		return nil, err
	}

	lastID, _ := cost.LastInsertId()

	return &modelCost.Cost{ID: lastID,
		LoanOption:   c.LoanOption,
		IDLoanOption: c.IDLoanOption,
		Interest:     c.Interest,
		AdminFee:     c.AdminFee,
		FinePerDay:   c.FinePerDay,
		Description:  c.Description,
		CreatedDate:  time.Now(),
		UpdatedDate:  time.Now()}, nil
}
func (r repo) UpdateCostByID(id int64, params data.Params) (int64, error) {
	loanOption := params.GetValue("loan_option")
	interest := params.GetValue("interest")
	adminFee := params.GetValue("admin_fee")
	finePerDay := params.GetValue("fine_per_day")
	description := params.GetValue("description")

	query := `UPDATE tbl_cost SET loan_option = ?, interest = ?, admin_fee = ?,
		fine_per_day = ?, description = ?
		WHERE id = ?`
	result, err := r.db.Exec(query,
		loanOption, interest, adminFee,
		finePerDay, description, id)
	if err != nil {
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r repo) DeleteCostByID(id int64) (httpStatus int, err error) {
	query := `DELETE FROM tbl_cost WHERE id = ?`
	_, err = r.db.Exec(query, id)
	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, nil
}

//Cost Explain
func (r repo) GetCostExplanationByID(id int64) (*modelCost.CostExplanation, error) {
	cost := &modelCost.CostExplanation{}
	err := r.db.Get(cost, "SELECT * FROM tbl_cost_explain WHERE id = ?", id)
	return cost, err
}
func (r repo) ListCostExplanation() ([]modelCost.CostExplanation, error) {
	cost := []modelCost.CostExplanation{}
	err := r.db.Select(&cost, "SELECT * FROM tbl_cost_explain")
	return cost, err
}
func (r repo) CreateCostExplanation(ce *modelCost.CostExplanation) (*modelCost.CostExplanation, error) {
	arg := map[string]interface{}{
		"title":       ce.Title,
		"description": ce.Description,
	}

	query := `INSERT INTO tbl_cost_explain
		SET title = :title, description = :description`

	cost, err := r.db.NamedExec(query, arg)
	if err != nil {
		return nil, err
	}

	lastID, _ := cost.LastInsertId()

	return &modelCost.CostExplanation{ID: lastID,
		Title:       ce.Title,
		Description: ce.Description,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now()}, nil
}
func (r repo) UpdateCostExplanationByID(id int64, params data.Params) (int64, error) {
	title := params.GetValue("title")
	description := params.GetValue("description")

	query := `UPDATE tbl_cost_explain SET title = ?, description = ?
		WHERE id = ?`
	result, err := r.db.Exec(query,
		title, description, id)
	if err != nil {
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r repo) DeleteCostExplanationByID(id int64) (httpStatus int, err error) {
	query := `DELETE FROM tbl_cost_explain WHERE id = ?`
	_, err = r.db.Exec(query, id)
	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, nil
}