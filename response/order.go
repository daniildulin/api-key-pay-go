package response

type Order struct {
	Vkey        interface{}   `json:"vkey"`
	StatusTitle string        `json:"status_title"`
	Created     string        `json:"created"`
	UpdateAt    string        `json:"updateAt"`
	Expired     string        `json:"expired"`
	ID          string        `json:"id"`
	Reason      string        `json:"reason"`
	Agreement   string        `json:"agreement"`
	Status      string        `json:"status"`
	Rate        []float64     `json:"rate"`
	Extra       []interface{} `json:"extra"`
	Props       []Requisite   `json:"props"`
	Psid1       int           `json:"psid1"`
	Out         float64       `json:"out"`
	In          float64       `json:"in"`
	Psid2       int           `json:"psid2"`
	Direct      uint8         `json:"direct"`
}

type OrderSuccessResponse struct {
	Status    string `json:"status"`
	Reason    string `json:"reason,omitempty"`
	Msg       string `json:"msg,omitempty"`
	ErrorName string `json:"errname,omitempty"`
	Value     Order  `json:"value"`
	ErrorCode int    `json:"errcode,omitempty"`
}
