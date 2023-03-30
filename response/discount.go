package response

type Discount struct {
	Valute  string  `json:"valute"`
	Percent float64 `json:"percent"`
	Amount  float64 `json:"amount"`
}
