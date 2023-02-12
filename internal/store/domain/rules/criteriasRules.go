package rules

type CriteriasRules struct {
	Id       int64 `json:"id" db:"id"`
	Criteria int64 `json:"criteria" db:"criteria"`
	Reward   int64 `json:"reward" db:"reward"`
}
