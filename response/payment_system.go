package response

type PaymentSystem struct {
	Name      string  `json:"name"`
	Valute    string  `json:"valute"`
	Psid      int     `json:"psid"`
	Reserve   float64 `json:"reserve"`
	WithCodes int     `json:"with_codes"`
}

type PaymentSystemSuccessResponse struct {
	Status string          `json:"status"`
	Value  []PaymentSystem `json:"value"`
}
