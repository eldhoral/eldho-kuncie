package resp

type CostExplanationPage struct {
	Cost        []interface{} `json:"cost"`
	CostExplain []interface{} `json:"cost_explain"`
}

type Cost struct {
	ID           int64  `json:"id"`
	LoanOption   string `json:"loan_option"`
	IDLoanOption int    `json:"id_loan_option"`
	Interest     string `json:"interest"`
	AdminFee     string `json:"admin_fee"`
	FinePerDay   string `json:"fine_per_day"`
	Description  string `json:"description"`
	IsVisible    bool   `json:"is_visible"`
}

type CostExplain struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
