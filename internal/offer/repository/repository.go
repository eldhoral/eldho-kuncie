package repository

import (
	"net/http"
	"strconv"

	modelLanding "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/offer/domain/landing"
	modelTnc "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/offer/domain/tnc"
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
func (r repo) DeleteLoanLimit() (httpStatus int, err error) {
	query := `DELETE FROM tbl_loan_limit ORDER BY id ASC LIMIT 1`
	_, err = r.db.Exec(query)
	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, nil
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
func (r repo) DeleteBenefitByID(id int64) (httpStatus int, err error) {
	query := `DELETE FROM tbl_benefit WHERE id = ?`
	_, err = r.db.Exec(query, id)
	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, nil
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
func (r repo) DeleteLoanMethodByID(id int64) (httpStatus int, err error) {
	query := `DELETE FROM tbl_loan_method WHERE id = ?`
	_, err = r.db.Exec(query, id)
	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, nil
}

//Tnc
func (r repo) GetTncByID(id int64) (*modelTnc.Tnc, error) {
	tnc := &modelTnc.Tnc{}
	err := r.db.Get(tnc, "SELECT * FROM tbl_tnc WHERE id = ?", id)
	return tnc, err
}
func (r repo) ListTnc() ([]modelTnc.Tnc, error) {
	tnc := []modelTnc.Tnc{}
	err := r.db.Select(&tnc, "SELECT * FROM tbl_tnc")
	return tnc, err
}
func (r repo) CreateTnc(t *modelTnc.Tnc) (*modelTnc.Tnc, error) {
	arg := map[string]interface{}{
		"title": t.Title,
	}

	query := `INSERT INTO tbl_tnc
		SET title = :title`

	tnc, err := r.db.NamedExec(query, arg)
	if err != nil {
		return nil, err
	}

	lastID, _ := tnc.LastInsertId()

	return &modelTnc.Tnc{ID: lastID,
		Title: t.Title}, nil
}
func (r repo) UpdateTncByID(id int64, params data.Params) (int64, error) {
	title := params.GetValue("title")
	query := `UPDATE tbl_tnc SET title = ? WHERE id = ?`
	result, err := r.db.Exec(query, title, id)
	if err != nil {
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r repo) DeleteTncByID(id int64) (httpStatus int, err error) {
	query := `DELETE FROM tbl_tnc WHERE id = ?`
	_, err = r.db.Exec(query, id)
	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, nil
}

//Tnc title
func (r repo) GetTncTitleByID(id int64) (*modelTnc.TncTitle, error) {
	tnc := &modelTnc.TncTitle{}
	err := r.db.Get(tnc, "SELECT * FROM tbl_tnc_title WHERE id = ?", id)
	return tnc, err
}
func (r repo) ListTncTitle() ([]modelTnc.TncTitle, error) {
	tnc := []modelTnc.TncTitle{}
	err := r.db.Select(&tnc, "SELECT * FROM tbl_tnc_title")
	return tnc, err
}
func (r repo) ListTncTitleByID(idTnc int64) ([]modelTnc.TncTitle, error) {
	tnc := []modelTnc.TncTitle{}
	err := r.db.Select(&tnc, "SELECT * FROM tbl_tnc_title WHERE id_tnc = ? ORDER BY id ASC", idTnc)
	return tnc, err
}
func (r repo) CreateTncTitle(t *modelTnc.TncTitle) (*modelTnc.TncTitle, error) {
	arg := map[string]interface{}{
		"id_tnc": t.IDTnc,
		"title":  t.Title,
	}

	query := `INSERT INTO tbl_tnc_title
		SET id_tnc = :id_tnc, title = :title`

	tnc, err := r.db.NamedExec(query, arg)
	if err != nil {
		return nil, err
	}

	lastID, _ := tnc.LastInsertId()

	return &modelTnc.TncTitle{ID: lastID,
		IDTnc: t.IDTnc,
		Title: t.Title}, nil
}
func (r repo) UpdateTncTitleByID(id int64, params data.Params) (int64, error) {
	idTnc := params.GetInt64("id_tnc")
	title := params.GetValue("title")
	query := `UPDATE tbl_tnc_title SET id_tnc = ?, title = ? WHERE id = ?`
	result, err := r.db.Exec(query, idTnc, title, id)
	if err != nil {
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r repo) DeleteTncTitleByID(id int64) (httpStatus int, err error) {
	query := `DELETE FROM tbl_tnc_title WHERE id = ?`
	_, err = r.db.Exec(query, id)
	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, nil
}

//Tnc subtitle
func (r repo) GetTncSubtitleByID(id int64) (*modelTnc.TncSubtitle, error) {
	tnc := &modelTnc.TncSubtitle{}
	err := r.db.Get(tnc, "SELECT * FROM tbl_tnc_subtitle WHERE id = ?", id)
	return tnc, err
}
func (r repo) ListTncSubtitle() ([]modelTnc.TncSubtitle, error) {
	tnc := []modelTnc.TncSubtitle{}
	err := r.db.Select(&tnc, "SELECT * FROM tbl_tnc_subtitle")
	return tnc, err
}
func (r repo) ListTncSubtitleByID(idTncTitle int64) ([]*modelTnc.TncSubtitle, error) {
	tnc := []*modelTnc.TncSubtitle{}
	err := r.db.Select(&tnc, "SELECT * FROM tbl_tnc_subtitle WHERE id_tnc_title = ? ORDER BY id ASC", idTncTitle)
	return tnc, err
}
func (r repo) CreateTncSubtitle(t *modelTnc.TncSubtitle) (*modelTnc.TncSubtitle, error) {
	arg := map[string]interface{}{
		"id_tnc_title": t.IDTncTitle,
		"subtitle":     t.Subtitle,
	}

	query := `INSERT INTO tbl_tnc_subtitle
		SET id_tnc_title = :id_tnc_title, subtitle = :subtitle`

	tnc, err := r.db.NamedExec(query, arg)
	if err != nil {
		return nil, err
	}

	lastID, _ := tnc.LastInsertId()

	return &modelTnc.TncSubtitle{ID: lastID,
		IDTncTitle: t.IDTncTitle,
		Subtitle:   t.Subtitle}, nil
}
func (r repo) UpdateTncSubtitleByID(id int64, params data.Params) (int64, error) {
	idTncTitle := params.GetInt64("id_tnc_title")
	subtitle := params.GetValue("subtitle")
	query := `UPDATE tbl_tnc_subtitle SET id_tnc_title = ?, subtitle = ? WHERE id = ?`
	result, err := r.db.Exec(query, idTncTitle, subtitle, id)
	if err != nil {
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r repo) DeleteTncSubtitleByID(id int64) (httpStatus int, err error) {
	query := `DELETE FROM tbl_tnc_subtitle WHERE id = ?`
	_, err = r.db.Exec(query, id)
	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, nil
}

//Tnc explain
func (r repo) GetTncExplainByID(id int64) (*modelTnc.TncExplain, error) {
	tnc := &modelTnc.TncExplain{}
	err := r.db.Get(tnc, "SELECT * FROM tbl_tnc_explain WHERE id = ?", id)
	return tnc, err
}
func (r repo) ListTncExplain() ([]modelTnc.TncExplain, error) {
	tnc := []modelTnc.TncExplain{}
	err := r.db.Select(&tnc, "SELECT * FROM tbl_tnc_explain")
	return tnc, err
}
func (r repo) ListTncExplainByID(idTnc int64, idTncTitle int64) ([]modelTnc.TncExplain, error) {
	tnc := []modelTnc.TncExplain{}
	err := r.db.Select(&tnc, "SELECT * FROM tbl_tnc_explain WHERE id_tnc = ? AND id_tnc_title = ?", idTnc, idTncTitle)
	return tnc, err
}
func (r repo) ListTncExplainByIDWithSubtitle(idTnc int64, idTncTitle int64, idTncSubtitle int64) ([]*modelTnc.TncExplain, error) {
	tnc := []*modelTnc.TncExplain{}
	err := r.db.Select(&tnc, "SELECT * FROM tbl_tnc_explain WHERE id_tnc = ? AND id_tnc_title = ? AND id_tnc_subtitle = ?", idTnc, idTncTitle, idTncSubtitle)
	return tnc, err
}
func (r repo) CreateTncExplain(t *modelTnc.TncExplain) (*modelTnc.TncExplain, error) {
	arg := map[string]interface{}{
		"id_tnc":          t.IDTnc,
		"id_tnc_title":    t.IDTncTitle,
		"id_tnc_subtitle": t.IDTncSubtitle,
		"description":     t.Description,
	}

	query := `INSERT INTO tbl_tnc_explain
		SET id_tnc = :id_tnc, id_tnc_title = :id_tnc_title, id_tnc_subtitle = :id_tnc_subtitle, description = :description`

	tnc, err := r.db.NamedExec(query, arg)
	if err != nil {
		return nil, err
	}

	lastID, _ := tnc.LastInsertId()

	return &modelTnc.TncExplain{ID: lastID,
		IDTnc:         t.IDTnc,
		IDTncTitle:    t.IDTncTitle,
		IDTncSubtitle: t.IDTncSubtitle,
		Description:   t.Description}, nil
}
func (r repo) UpdateTncExplainByID(id int64, params data.Params) (int64, error) {
	description := params.GetValue("description")
	query := `UPDATE tbl_tnc_explain 
	SET description = ? WHERE id = ?`
	result, err := r.db.Exec(query, description, id)
	if err != nil {
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r repo) DeleteTncExplainByID(id int64) (httpStatus int, err error) {
	query := `DELETE FROM tbl_tnc_explain WHERE id = ?`
	_, err = r.db.Exec(query, id)
	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, nil
}
