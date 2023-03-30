package response

type ExchangeRateCalculateSuccessResponse struct {
	Status string                `json:"status"`
	Value  ExchangeRateCalculate `json:"value"`
}

type ExchangeRateCalculate struct {
	Alert     *string   `json:"alert"`
	Warnings  string    `json:"warnings"`
	Discount  Discount  `json:"discount"`
	Direction Direction `json:"direction"`
}
