package resp

type TncPage struct {
	MainTitle string     `json:"main_title"`
	Title     []TncTitle `json:"title"`
}

type TncTitle struct {
	Title    string        `json:"data_title"`
	Subtitle []TncSubtitle `json:"subtitle"`
}
type TncSubtitle struct {
	Subtitle *string      `json:"data_subtitle"`
	Explain  []TncExplain `json:"description"`
}

type TncExplain struct {
	Explain string `json:"data_description"`
}
