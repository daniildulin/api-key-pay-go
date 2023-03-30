package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/daniildulin/api-key-pay-go/entity"
	"github.com/daniildulin/api-key-pay-go/request"
	"github.com/daniildulin/api-key-pay-go/response"
	"net/http"
	"net/url"
	"time"
)

const apiKeyPayHost = "https://apikeypay.com/api"

type ApiKeyPayClient struct {
	client *http.Client
	aKey   string
	bKey   string
	sKey   string
}

func New(AKey, BKey, SKey string) *ApiKeyPayClient {
	return &ApiKeyPayClient{
		aKey: AKey,
		bKey: BKey,
		sKey: SKey,
		client: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *ApiKeyPayClient) Directions(ctx context.Context, nonce uint32) ([]response.Direction, error) {
	endpoint := "/1/directions"
	params := map[string]string{
		"nonce": fmt.Sprintf("%d", nonce),
		"akey":  c.aKey,
		"bkey":  c.bKey,
	}
	var result []response.Direction

	err := c.sendGetRequestWithQueryParams(ctx, endpoint, params, &result)

	return result, err
}

func (c *ApiKeyPayClient) DirectionsAvailable(ctx context.Context, nonce uint32, psID int, direct uint8) ([]response.Direction, error) {
	endpoint := "/1/directions-available"
	exchangeDirection := direct
	if exchangeDirection > 1 {
		exchangeDirection = 0
	}
	params := map[string]string{
		"nonce":  fmt.Sprintf("%d", nonce),
		"akey":   c.aKey,
		"bkey":   c.bKey,
		"psid":   fmt.Sprintf("%d", psID),
		"direct": fmt.Sprintf("%d", exchangeDirection),
	}

	var result []response.Direction

	err := c.sendGetRequestWithQueryParams(ctx, endpoint, params, &result)

	return result, err
}

func (c *ApiKeyPayClient) PaymentSystems(ctx context.Context, nonce uint32) ([]response.PaymentSystem, error) {
	endpoint := "/1/payment-systems"
	params := map[string]string{
		"nonce": fmt.Sprintf("%d", nonce),
		"akey":  c.aKey,
		"bkey":  c.bKey,
	}
	var result []response.PaymentSystem

	err := c.sendGetRequestWithQueryParams(ctx, endpoint, params, &result)

	return result, err
}

func (c *ApiKeyPayClient) ExchangeRateCalculate(ctx context.Context, options *entity.ExchangeRateCalculateOptions) (*response.ExchangeRateCalculate, error) {
	endpoint := "/1/exchange-rate-calculate"
	params := map[string]string{
		"nonce":  fmt.Sprintf("%d", options.Nonce),
		"akey":   c.aKey,
		"bkey":   c.bKey,
		"psid1":  fmt.Sprintf("%d", options.Psid1),
		"psid2":  fmt.Sprintf("%d", options.Psid2),
		"amount": fmt.Sprintf("%f", options.Amount),
		"direct": fmt.Sprintf("%d", options.Direct),
	}

	if options.CouponName != "" {
		params["coupon_name"] = options.CouponName
		params["skey"] = c.sKey
	}

	if options.AmountValute != "" {
		params["amount_valute"] = options.AmountValute
	}

	var result response.ExchangeRateCalculate

	err := c.sendGetRequestWithQueryParams(ctx, endpoint, params, &result)

	return &result, err
}

func (c *ApiKeyPayClient) OrderPropsGet(ctx context.Context, nonce uint32, psID1, psID2 int) ([]response.Requisite, error) {
	endpoint := "/1/order-props-get"
	params := map[string]string{
		"nonce": fmt.Sprintf("%d", nonce),
		"akey":  c.aKey,
		"bkey":  c.bKey,
		"psid1": fmt.Sprintf("%d", psID1),
		"psid2": fmt.Sprintf("%d", psID2),
	}
	var result []response.Requisite

	err := c.sendGetRequestWithQueryParams(ctx, endpoint, params, &result)

	return result, err
}

func (c *ApiKeyPayClient) CreateOrder(ctx context.Context, nonce uint32, order *request.NewOrderData, userOrderID string, redirect *request.Redirect) (*response.Order, error) {
	endpoint := fmt.Sprintf("1/order-create/?bkey=%s", c.bKey)

	orderRequestBody := request.NewOrder{
		Nonce:       nonce,
		Akey:        c.aKey,
		Bkey:        c.bKey,
		Order:       *order,
		Confirm:     false,
		PayInfo:     false,
		Skey:        "",
		UserOrderID: userOrderID,
		Redirect:    redirect,
	}

	jsonData, err := json.Marshal(orderRequestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", apiKeyPayHost, endpoint), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	res := response.Order{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *ApiKeyPayClient) OrderValidate(ctx context.Context, nonce uint32, orderID uint64) (*response.Order, error) {
	endpoint := "/1/order-validate"
	params := map[string]string{
		"nonce":    fmt.Sprintf("%d", nonce),
		"order_id": fmt.Sprintf("%d", orderID),
		"akey":     c.aKey,
		"bkey":     c.bKey,
	}
	var result response.OrdersValidate

	err := c.sendGetRequestWithQueryParams(ctx, endpoint, params, &result)

	if result.Status == "card" {
		return &result.Order, errors.New(result.Msg)
	}

	return &result.Order, err
}

func (c *ApiKeyPayClient) OrderPayInfo(ctx context.Context, nonce uint32, orderID uint64, paymentDetail bool) (*response.OrderPayInfo, error) {
	endpoint := "/1/order-pay-info"
	params := map[string]string{
		"nonce":    fmt.Sprintf("%d", nonce),
		"order_id": fmt.Sprintf("%d", orderID),
		"akey":     c.aKey,
		"bkey":     c.bKey,
		"pd":       fmt.Sprintf("%t", paymentDetail),
	}
	var result response.OrderPayInfo

	err := c.sendGetRequestWithQueryParams(ctx, endpoint, params, &result)

	return &result, err
}

func (c *ApiKeyPayClient) OrderConfirm(ctx context.Context, nonce uint32, orderID uint64) (*response.Order, error) {
	endpoint := "/1/order-confirm"
	params := map[string]string{
		"nonce":    fmt.Sprintf("%d", nonce),
		"order_id": fmt.Sprintf("%d", orderID),
		"akey":     c.aKey,
		"bkey":     c.bKey,
	}
	var result response.Order

	err := c.sendGetRequestWithQueryParams(ctx, endpoint, params, &result)

	return &result, err
}

func (c *ApiKeyPayClient) OrderCancel(ctx context.Context, nonce uint32, orderID uint64) error {
	endpoint := "/1/order-cancel"
	params := map[string]string{
		"nonce":    fmt.Sprintf("%d", nonce),
		"order_id": fmt.Sprintf("%d", orderID),
		"akey":     c.aKey,
		"bkey":     c.bKey,
	}

	err := c.sendGetRequestWithQueryParams(ctx, endpoint, params, nil)

	return err
}

func (c *ApiKeyPayClient) OrderGet(ctx context.Context, nonce uint32, orderID uint64) (*response.Order, error) {
	endpoint := "/1/order-get"
	params := map[string]string{
		"nonce":    fmt.Sprintf("%d", nonce),
		"order_id": fmt.Sprintf("%d", orderID),
		"akey":     c.aKey,
		"bkey":     c.bKey,
	}
	var result response.Order

	err := c.sendGetRequestWithQueryParams(ctx, endpoint, params, &result)

	return &result, err
}

func (c *ApiKeyPayClient) sendGetRequestWithQueryParams(ctx context.Context, endpoint string, params map[string]string, v interface{}) error {
	query := c.mapParamsToQueryString(params)
	link := fmt.Sprintf("%s/%s?%s", apiKeyPayHost, endpoint, query)
	req, err := http.NewRequest("GET", link, http.NoBody)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)

	if err := c.sendRequest(req, v); err != nil {
		return err
	}
	return nil
}

func (c *ApiKeyPayClient) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes response.FailResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	success := response.SuccessResponse{}
	err = json.NewDecoder(res.Body).Decode(&success)
	if err != nil {
		return err
	}

	if success.ErrorCode != 0 || success.ErrorName != "" {
		return errors.New(success.Message)
	}

	if v != nil {
		err = json.Unmarshal(success.Value, &v)
	}

	return err
}

func (c *ApiKeyPayClient) mapParamsToQueryString(params map[string]string) string {
	values := url.Values{}
	for k, v := range params {
		values.Add(k, v)
	}
	return values.Encode()
}
