package resp

type LandingPage struct {
	LoanLimit  *LoanLimit    `json:"loan_limit"`
	Benefit    []interface{} `json:"benefit"`
	LoanMethod []interface{} `json:"loan_method"`
}

type LoanLimit struct {
	ID    int64 `json:"id"`
	Limit int64 `json:"limit"`
}

type Benefit struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type LoanMethod struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
