package rules

type DiscountRules struct {
	Id        int64 `json:"id" db:"id"`
	ProductId int64 `json:"product_id" db:"product_id"`
	Rules     int64 `json:"rules" db:"rules"`
}
