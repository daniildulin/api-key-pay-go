package entity

type ExchangeRateCalculateOptions struct {
	Skey         string  `json:"skey"`
	CouponName   string  `json:"coupon_name"`
	AmountValute string  `json:"amount_valute"`
	Nonce        uint32  `json:"nonce"`
	Psid1        int     `json:"psid1"`
	Psid2        int     `json:"psid2"`
	Amount       float64 `json:"amount"`
	Direct       int     `json:"direct"`
}
