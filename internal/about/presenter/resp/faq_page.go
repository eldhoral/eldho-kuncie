package resp

type FAQPage struct {
	FAQ []interface{} `json:"data"`
}

type FAQ struct {
	Id       int64         `json:"id"`
	IdOrder  int64         `json:"id_order"`
	Title    string        `json:"main_title"`
	FAQTitle []interface{} `json:"data"`
}

type FAQTitle struct {
	ID          int64  `json:"id"`
	IdOrder     int64  `json:"id_order"`
	IDFaq       int64  `json:"id_faq"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
