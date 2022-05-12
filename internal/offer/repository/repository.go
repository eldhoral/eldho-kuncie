package repository

import (
	"strconv"

	modelLanding "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/offer/domain/landing"
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

//Loan limit
func (r repo) GetLoanLimitByID(id int64) (*modelLanding.LoanLimit, error) {
	loanLimit := &modelLanding.LoanLimit{}
	err := r.db.Get(loanLimit, "SELECT * FROM tbl_loan_limit WHERE id = ?", id)
	return loanLimit, err
}

func (r repo) GetLoanLimit() (*modelLanding.LoanLimit, error) {
	loanLimit := &modelLanding.LoanLimit{}
	err := r.db.Get(loanLimit, "SELECT * FROM tbl_loan_limit ORDER BY id ASC LIMIT 1")
	return loanLimit, err
}

func (r repo) CreateLoanLimit(ll *modelLanding.LoanLimit) (*modelLanding.LoanLimit, error) {
	arg := map[string]interface{}{
		"limit": ll.Limit,
	}

	query := `INSERT INTO tbl_loan_limit
		SET loan_limit = :limit`

	loanLimit, err := r.db.NamedExec(query, arg)
	if err != nil {
		return nil, err
	}

	lastID, _ := loanLimit.LastInsertId()

	return &modelLanding.LoanLimit{ID: lastID,
		Limit: ll.Limit}, nil
}

func (r repo) UpdateLoanLimit(limit string) (int64, error) {
	limitLoan, _ := strconv.Atoi(limit)
	query := `UPDATE tbl_loan_limit SET loan_limit = ?
		ORDER BY id ASC LIMIT 1`
	result, err := r.db.Exec(query, float64(limitLoan))
	if err != nil {
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r repo) DeleteLoanLimit() error {
	query := `DELETE FROM tbl_loan_limit ORDER BY id ASC LIMIT 1`
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

//Benefit
func (r repo) GetBenefitByID(id int64) (*modelLanding.Benefit, error) {
	benefit := &modelLanding.Benefit{}
	err := r.db.Get(benefit, "SELECT * FROM tbl_benefit WHERE id = ?", id)
	return benefit, err
}

func (r repo) ListBenefit() ([]modelLanding.Benefit, error) {
	benefit := []modelLanding.Benefit{}
	err := r.db.Select(&benefit, "SELECT * FROM tbl_benefit")
	return benefit, err
}

func (r repo) CreateBenefit(b *modelLanding.Benefit) (*modelLanding.Benefit, error) {
	arg := map[string]interface{}{
		"title":       b.Title,
		"description": b.Description,
		"image":       b.Image,
	}

	query := `INSERT INTO tbl_benefit
		SET title = :title, description = :description, image = :image`

	benefit, err := r.db.NamedExec(query, arg)
	if err != nil {
		return nil, err
	}

	lastID, _ := benefit.LastInsertId()

	return &modelLanding.Benefit{ID: lastID,
		Title:       b.Title,
		Description: b.Description,
		Image:       b.Image}, nil
}

func (r repo) UpdateBenefitByID(id int64, params data.Params, path string) (int64, error) {
	title := params.GetValue("title")
	description := params.GetValue("description")
	if path == "" {
		query := `UPDATE tbl_benefit SET title = ?, description = ?
		WHERE id = ?`
		result, err := r.db.Exec(query, title, description, id)
		if err != nil {
			return 0, err
		}
		count, err := result.RowsAffected()
		if err != nil {
			return 0, err
		}
		return count, nil
	}
	query := `UPDATE tbl_benefit SET title = ?, description = ?, image = ?
		WHERE id = ?`
	result, err := r.db.Exec(query, title, description, path, id)
	if err != nil {
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r repo) DeleteBenefitByID(id int64) error {
	query := `DELETE FROM tbl_benefit WHERE id = ?`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

//Loan method
func (r repo) GetLoanMethodByID(id int64) (*modelLanding.LoanMethod, error) {
	loanMethod := &modelLanding.LoanMethod{}
	err := r.db.Get(loanMethod, "SELECT * FROM tbl_loan_method WHERE id = ?", id)
	return loanMethod, err
}

func (r repo) ListLoanMethod() ([]modelLanding.LoanMethod, error) {
	loanMethod := []modelLanding.LoanMethod{}
	err := r.db.Select(&loanMethod, "SELECT * FROM tbl_loan_method")
	return loanMethod, err
}

func (r repo) CreateLoanMethod(lm *modelLanding.LoanMethod) (*modelLanding.LoanMethod, error) {
	arg := map[string]interface{}{
		"title":       lm.Title,
		"description": lm.Description,
	}

	query := `INSERT INTO tbl_loan_method
		SET title = :title, description = :description`

	benefit, err := r.db.NamedExec(query, arg)
	if err != nil {
		return nil, err
	}

	lastID, _ := benefit.LastInsertId()

	return &modelLanding.LoanMethod{ID: lastID,
		Title:       lm.Title,
		Description: lm.Description}, nil
}

func (r repo) UpdateLoanMethodByID(id int64, title string, description string) (int64, error) {
	query := `UPDATE tbl_loan_method SET title = ?, description = ?
		WHERE id = ?`
	result, err := r.db.Exec(query, title, description, id)
	if err != nil {
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r repo) DeleteLoanMethodByID(id int64) error {
	query := `DELETE FROM tbl_loan_method WHERE id = ?`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
