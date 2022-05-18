package offer

import (
	"time"
)

type Tnc struct {
	ID          int64     `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	CreatedDate time.Time `json:"created_date" db:"created_date"`
	UpdatedDate time.Time `json:"updated_date" db:"updated_date"`
}

type TncTitle struct {
	ID          int64     `json:"id" db:"id"`
	IDTnc       int64     `json:"id_tnc" db:"id_tnc"`
	Title       string    `json:"title" db:"title"`
	CreatedDate time.Time `json:"created_date" db:"created_date"`
	UpdatedDate time.Time `json:"updated_date" db:"updated_date"`
}

type TncSubtitle struct {
	ID          int64     `json:"id" db:"id"`
	IDTncTitle  int64     `json:"id_tnc_title" db:"id_tnc_title"`
	Subtitle    *string   `json:"subtitle" db:"subtitle"`
	CreatedDate time.Time `json:"created_date" db:"created_date"`
	UpdatedDate time.Time `json:"updated_date" db:"updated_date"`
}

type TncExplain struct {
	ID            int64     `json:"id" db:"id"`
	IDTnc         int64     `json:"id_tnc" db:"id_tnc"`
	IDTncTitle    int64     `json:"id_tnc_title" db:"id_tnc_title"`
	IDTncSubtitle *int64    `json:"id_tnc_subtitle" db:"id_tnc_subtitle"`
	Description   string    `json:"description" db:"description"`
	CreatedDate   time.Time `json:"created_date" db:"created_date"`
	UpdatedDate   time.Time `json:"updated_date" db:"updated_date"`
}
