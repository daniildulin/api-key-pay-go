package response

type OrderPayInfo struct {
	Alert   *string            `json:"alert"`
	SCI     SCI                `json:"SCI"`
	Extra   string             `json:"extra"`
	Info    []OrderPayInfoItem `json:"info"`
	Pd      []PaymentDetail    `json:"pd"`
	OrderID string             `json:"order_id"`
}

type OrderPayInfoItem struct {
	Value interface{} `json:"value"`
	Extra *string     `json:"extra"`
	Title string      `json:"title"`
}

type OrderPayInfoSuccessResponse struct {
	Status string       `json:"status"`
	Value  OrderPayInfo `json:"value"`
}
