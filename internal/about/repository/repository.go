package repository

import (
	"net/http"
	"time"

	modelCost "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/about/domain/cost"
	modelFaq "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/about/domain/faq"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/data"
	generator "bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/query"

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
	err := r.db.Select(&cost, "SELECT * FROM tbl_cost WHERE is_visible = 1")
	return cost, err
}
func (r repo) ListCostByIDLoanOption(idLoanOption int64) ([]modelCost.Cost, error) {
	cost := []modelCost.Cost{}
	err := r.db.Select(&cost, "SELECT * FROM tbl_cost where id_loan_option = ?", idLoanOption)
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
	isVisible := params.GetValue("is_visible")

	query := `UPDATE tbl_cost SET loan_option = ?, interest = ?, admin_fee = ?,
		fine_per_day = ?, description = ?, is_visible = ?
		WHERE id = ?`
	result, err := r.db.Exec(query,
		loanOption, interest, adminFee,
		finePerDay, description, isVisible, id)
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

//Faq
func (r repo) GetFaqID(id int64) (*modelFaq.Faq, error) {
	faq := &modelFaq.Faq{}
	err := r.db.Get(faq, "SELECT * FROM tbl_faq WHERE id = ? ORDER BY id_order ASC", id)
	return faq, err
}
func (r repo) GetFaqIDOrder(idOrder int64) (*modelFaq.Faq, error) {
	faq := &modelFaq.Faq{}
	err := r.db.Get(faq, "SELECT * FROM tbl_faq WHERE id_order = ? ORDER BY id_order ASC", idOrder)
	return faq, err
}
func (r repo) ListFaq() ([]modelFaq.Faq, error) {
	faq := []modelFaq.Faq{}
	err := r.db.Select(&faq, "SELECT * FROM tbl_faq ORDER BY id_order ASC")
	return faq, err
}
func (r repo) CreateFaq(f *modelFaq.Faq) (*modelFaq.Faq, error) {
	arg := map[string]interface{}{
		"title":    f.Title,
		"id_order": f.IDOrder,
	}

	query := `INSERT INTO tbl_faq
		SET title = :title, id_order = :id_order`

	cost, err := r.db.NamedExec(query, arg)
	if err != nil {
		return nil, err
	}

	lastID, _ := cost.LastInsertId()

	return &modelFaq.Faq{ID: lastID,
		Title:       f.Title,
		IDOrder:     f.IDOrder,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now()}, nil
}
func (r repo) UpdateFaqByID(id int64, params data.Params) (int64, error) {
	arg := map[string]interface{}{
		"id_order": params.GetString("id_order"),
		"title":    params.GetString("title"),
		"id":       id,
	}
	column := []string{"id_order", "title"}
	columns := generator.DynamicUpdateStatement(column, params)

	query := "UPDATE tbl_faq SET " + columns +
		"WHERE id = :id"
	result, err := r.db.NamedExec(query, arg)
	if err != nil {
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r repo) DeleteFaqByID(id int64) (httpStatus int, err error) {
	query := `DELETE FROM tbl_faq WHERE id = ?`
	_, err = r.db.Exec(query, id)
	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, nil
}

//Faq Title
func (r repo) GetFaqTitleID(id int64) (*modelFaq.FaqTitle, error) {
	faq := &modelFaq.FaqTitle{}
	err := r.db.Get(faq, "SELECT * FROM tbl_faq_title WHERE id = ? ORDER BY id_order ASC", id)
	return faq, err
}
func (r repo) GetFaqTitleIDOrder(idOrder int64) (*modelFaq.FaqTitle, error) {
	faq := &modelFaq.FaqTitle{}
	err := r.db.Get(faq, "SELECT * FROM tbl_faq_title WHERE id_order = ? ORDER BY id_order ASC", idOrder)
	return faq, err
}
func (r repo) ListFaqTitle() ([]modelFaq.FaqTitle, error) {
	faq := []modelFaq.FaqTitle{}
	err := r.db.Select(&faq, "SELECT * FROM tbl_faq_title ORDER BY id_order ASC")
	return faq, err
}
func (r repo) AutoIncrementIDOrder(idOrder int64) (int64, error) {
	query := "UPDATE tbl_faq_title SET id_order = id_order + 1 WHERE id_order >= ? ORDER BY id_order ASC"
	result, err := r.db.Exec(query, idOrder)
	if err != nil {
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r repo) AutoDecrementIDOrder(idOrder int64) (int64, error) {
	query := "UPDATE tbl_faq_title SET id_order = id_order - 1 WHERE id_order >= ? ORDER BY id_order ASC"
	result, err := r.db.Exec(query, idOrder)
	if err != nil {
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r repo) ListFaqTitleByIDFaq(idFaq int64) ([]modelFaq.FaqTitle, error) {
	faq := []modelFaq.FaqTitle{}
	err := r.db.Select(&faq, "SELECT * FROM tbl_faq_title where id_faq = ?", idFaq)
	return faq, err
}
func (r repo) CreateFaqTitle(ft *modelFaq.FaqTitle) (*modelFaq.FaqTitle, error) {
	// update the ordering number of id order
	_, err := r.AutoIncrementIDOrder(ft.IDOrder)

	arg := map[string]interface{}{
		"id_faq":      ft.IDFaq,
		"title":       ft.Title,
		"description": ft.Description,
		"id_order":    ft.IDOrder,
	}
	query := `INSERT INTO tbl_faq_title
		SET id_faq = :id_faq, title = :title, description = :description, id_order = :id_order`
	cost, err := r.db.NamedExec(query, arg)
	if err != nil {
		return nil, err
	}

	lastID, _ := cost.LastInsertId()

	return &modelFaq.FaqTitle{ID: lastID,
		IDFaq:       ft.IDFaq,
		Title:       ft.Title,
		Description: ft.Description,
		IDOrder:     ft.IDOrder,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now()}, nil
}
func (r repo) UpdateFaqTitleByID(id int64, params data.Params) (int64, error) {
	arg := map[string]interface{}{
		"title":       params.GetString("title"),
		"id_order":    params.GetString("id_order"),
		"id_faq":      params.GetString("id_faq"),
		"description": params.GetString("description"),
		"id":          id,
	}
	column := []string{"title", "id_order", "id_faq", "description"}
	columns := generator.DynamicUpdateStatement(column, params)

	query := "UPDATE tbl_faq_title SET " + columns +
		"WHERE id = :id"
	result, err := r.db.NamedExec(query, arg)
	if err != nil {
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r repo) DeleteFaqTitleByID(id int64) (httpStatus int, err error) {
	getIdOrder, err := r.GetFaqTitleID(id)
	// update the ordering number of id order
	_, err = r.AutoDecrementIDOrder(getIdOrder.IDOrder)

	query := `DELETE FROM tbl_faq_title WHERE id = ?`
	_, err = r.db.Exec(query, id)
	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, nil
}
