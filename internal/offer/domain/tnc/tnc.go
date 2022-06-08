package offer

import (
	"time"
)

type TncMobile struct {
	ID          int64     `json:"id" db:"id"`
	Description string    `json:"description" db:"description"`
	CreatedDate time.Time `json:"created_date" db:"created_date"`
	UpdatedDate time.Time `json:"updated_date" db:"updated_date"`
}
