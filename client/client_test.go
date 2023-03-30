package client

import (
	"context"
	"github.com/daniildulin/api-key-pay-go/entity"
	"github.com/daniildulin/api-key-pay-go/request"
	"github.com/daniildulin/api-key-pay-go/response"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestDirections(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Error(".env file not found")
	}
	client := New(os.Getenv("APP_AKEY"), os.Getenv("APP_BKEY"), os.Getenv("APP_SKEY"))

	data, err := client.Directions(context.Background(), 1)
	if err != nil {
		t.Error(err)
	}

	if data == nil {
		t.Error("data is nil")
	}
}

func TestDirectionsAvailable(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Error(".env file not found")
	}
	client := New(os.Getenv("APP_AKEY"), os.Getenv("APP_BKEY"), os.Getenv("APP_SKEY"))

	data, err := client.DirectionsAvailable(context.Background(), 1, 319, 1)
	if err != nil {
		t.Error(err)
	}

	if data == nil {
		t.Error("data is nil")
	}
}

func TestPaymentSystems(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Error(".env file not found")
	}
	client := New(os.Getenv("APP_AKEY"), os.Getenv("APP_BKEY"), os.Getenv("APP_SKEY"))

	data, err := client.PaymentSystems(context.Background(), 1)
	if err != nil {
		t.Error(err)
	}

	if data == nil {
		t.Error("data is nil")
	}
}

func TestExchangeRateCalculate(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Error(".env file not found")
	}
	client := New(os.Getenv("APP_AKEY"), os.Getenv("APP_BKEY"), os.Getenv("APP_SKEY"))

	data, err := client.ExchangeRateCalculate(context.Background(), &entity.ExchangeRateCalculateOptions{
		Nonce:  1,
		Psid1:  12,
		Psid2:  9,
		Amount: 1000,
		Direct: 0,
	})
	if err != nil {
		t.Error(err)
	}

	if data == nil {
		t.Error("data is nil")
	}
}

func TestOrderCreate(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Error(".env file not found")
	}
	client := New(os.Getenv("APP_AKEY"), os.Getenv("APP_BKEY"), os.Getenv("APP_SKEY"))

	nonce := uint32(1)
	value := 1.0
	psID := 42
	userAgent := os.Getenv("APP_USER_AGENT")
	clientIP := os.Getenv("APP_USER_IP")

	var direction response.Direction

	directionsList, err := client.DirectionsAvailable(context.Background(), 1, psID, 1)
	if err != nil {
		t.Error(err)
	}

	if len(directionsList) == 0 {
		t.Error("directions list is empty")
	}

	for i := range directionsList {
		if directionsList[i].Psid1 == 7 || directionsList[i].Psid1 == 8 || directionsList[i].Psid1 == 18 {
			direction = directionsList[i]
			break
		}
	}
	nonce++

	exchangeRate, err := client.ExchangeRateCalculate(context.Background(), &entity.ExchangeRateCalculateOptions{
		Nonce:  nonce,
		Psid1:  direction.Psid1,
		Psid2:  direction.Psid2,
		Amount: value,
		Direct: 1,
	})
	if err != nil {
		t.Error(err)
	}
	nonce++

	orderPropsRequired, err := client.OrderPropsGet(context.Background(), nonce, exchangeRate.Direction.Psid1, exchangeRate.Direction.Psid2)
	if err != nil {
		t.Error(err)
	}
	nonce++

	orderProps := make([]request.RequisiteShort, len(orderPropsRequired))
	for i := range orderPropsRequired {
		switch orderPropsRequired[i].Name {
		case "from_acc":
			orderProps[i] = request.RequisiteShort{
				Name:  orderPropsRequired[i].Name,
				Value: os.Getenv("APP_USER_CARD"),
			}
		case "from_fio":
			orderProps[i] = request.RequisiteShort{
				Name:  orderPropsRequired[i].Name,
				Value: os.Getenv("APP_USER_NAME"),
			}
		case "email":
			orderProps[i] = request.RequisiteShort{
				Name:  orderPropsRequired[i].Name,
				Value: os.Getenv("APP_USER_EMAIL"),
			}
		case "to_acc":
			orderProps[i] = request.RequisiteShort{
				Name:  orderPropsRequired[i].Name,
				Value: os.Getenv("APP_USER_BLOCKCHAIN_ADDRESS"),
			}
		}
	}

	in, err := exchangeRate.Direction.In.Float64()
	if err != nil {
		t.Error(err)
	}

	out, err := exchangeRate.Direction.Out.Float64()
	if err != nil {
		t.Error(err)
	}

	newOrderData := &request.NewOrderData{
		IP:        clientIP,
		UserAgent: userAgent,
		Agreement: "yes",
		Props:     orderProps,
		Psid1:     exchangeRate.Direction.Psid1,
		Psid2:     exchangeRate.Direction.Psid2,
		In:        in,
		Out:       out,
		Direct:    1,
	}

	data, err := client.CreateOrder(context.Background(), nonce, newOrderData, "test-id", nil)
	if err != nil {
		t.Error(err)
	}

	if data == nil {
		t.Error("data is nil")
	}
}
