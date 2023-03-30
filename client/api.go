package client

import (
	"github.com/daniildulin/api-key-pay-go/entity"
	"github.com/daniildulin/api-key-pay-go/request"
	"github.com/daniildulin/api-key-pay-go/response"
)

type API interface {
	Directions(nonce uint32) ([]response.Direction, error)
	PaymentSystems(nonce uint32) ([]response.PaymentSystem, error)
	ExchangeRateCalculate(options entity.ExchangeRateCalculateOptions) (*response.ExchangeRateCalculate, error)
	OrderPropsGet(nonce uint32, psID1, psID2 int) ([]response.Requisite, error)
	OrderCreate(nonce uint32, order *request.NewOrder) (*response.Order, error)
	OrderValidate(nonce uint32, orderID uint64) (*response.Order, error)
	OrderPayInfo(nonce uint32, orderID uint64, paymentDetail bool) (*response.OrderPayInfo, error)
	OrderConfirm(nonce uint32, orderID uint64) (*response.Order, error)
	OrderCancel(nonce uint32, orderID uint64) error
}
