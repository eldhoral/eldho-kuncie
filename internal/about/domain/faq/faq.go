package offer

import "time"

type Faq struct {
	ID          int64     `json:"id" db:"id"`
	IDOrder     int64     `json:"id_order" db:"id_order"`
	Title       string    `json:"title" db:"title"`
	CreatedDate time.Time `json:"created_date" db:"created_date"`
	UpdatedDate time.Time `json:"updated_date" db:"updated_date"`
}

type FaqTitle struct {
	ID          int64     `json:"id" db:"id"`
	IDOrder     int64     `json:"id_order" db:"id_order"`
	IDFaq       int64     `json:"id_faq" db:"id_faq"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	CreatedDate time.Time `json:"created_date" db:"created_date"`
	UpdatedDate time.Time `json:"updated_date" db:"updated_date"`
}
