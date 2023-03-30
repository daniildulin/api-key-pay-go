package request

type NewOrder struct {
	Nonce       uint32       `json:"nonce"`
	Akey        string       `json:"akey,omitempty"`
	Bkey        string       `json:"bkey,omitempty"`
	Order       NewOrderData `json:"Order"`
	Confirm     bool         `json:"confirm"`
	PayInfo     bool         `json:"pay-info"`
	Skey        string       `json:"skey,omitempty"`
	UserOrderID string       `json:"vkey,omitempty"`
	Redirect    *Redirect    `json:"redirect,omitempty"`
}

type NewOrderData struct {
	IP         string           `json:"ip"`
	UserAgent  string           `json:"HTTP_USER_AGENT"`
	Agreement  string           `json:"agreement"`
	CouponName string           `json:"coupon_name,omitempty"`
	Props      []RequisiteShort `json:"props"`
	Psid1      int              `json:"psid1"`
	Psid2      int              `json:"psid2"`
	In         float64          `json:"in"`
	Out        float64          `json:"out"`
	Direct     int              `json:"direct"`
}

type Redirect struct {
	SuccessLink string `json:"success"`
	ErrorLink   string `json:"error"`
}
