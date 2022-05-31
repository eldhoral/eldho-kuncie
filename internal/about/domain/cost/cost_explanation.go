package offer

import "time"

type Cost struct {
	ID           int64     `json:"id" db:"id"`
	LoanOption   string    `json:"loan_option" db:"loan_option"`
	IDLoanOption int       `json:"id_loan_option" db:"id_loan_option"`
	Interest     string    `json:"interest" db:"interest"`
	AdminFee     string    `json:"admin_fee" db:"admin_fee"`
	FinePerDay   string    `json:"fine_per_day" db:"fine_per_day"`
	Description  string    `json:"description" db:"description"`
	IsVisible    bool      `json:"is_visible" db:"is_visible"`
	CreatedDate  time.Time `json:"created_date" db:"created_date"`
	UpdatedDate  time.Time `json:"updated_date" db:"updated_date"`
}

type CostExplanation struct {
	ID          int64     `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	CreatedDate time.Time `json:"created_date" db:"created_date"`
	UpdatedDate time.Time `json:"updated_date" db:"updated_date"`
}
