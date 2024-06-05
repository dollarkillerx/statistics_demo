package test

import (
	"fmt"
	"testing"

	"github.com/dollarkillerx/urllib"
)

func TestApiInit(t *testing.T) {
	_, bytes, err := urllib.Post("http://127.0.0.1:8475/api/v1/init").
		SetJsonObject(map[string]interface{}{
			"account": "my_test",
			"balance": 1000,
			"lever":   2000,
		}).Byte()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestSymbolInfoTick(t *testing.T) {
	for i := 0; i < 1000; i++ {
		_, bytes, err := urllib.Post("http://127.0.0.1:8475/api/v1/symbol_info_tick").
			SetJsonObject(map[string]interface{}{
				"symbol": "EURUSD",
			}).Byte()
		if err != nil {
			panic(err)
		}

		fmt.Println(string(bytes))
	}
}

func TestOrderSend(t *testing.T) {
	_, bytes, err := urllib.Post("http://127.0.0.1:8475/api/v1/order_send").
		SetJsonObject(map[string]interface{}{
			"symbol":  "EURUSD",
			"volume":  0.05,
			"type":    0,
			"price":   1.06726,
			"account": "my_test",
		}).Byte()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestOrderSendClose(t *testing.T) {
	_, bytes, err := urllib.Post("http://127.0.0.1:8475/api/v1/order_send").
		SetJsonObject(map[string]interface{}{
			//"position": 11,
			"symbol":  "EURUSD",
			"volume":  0.05,
			"type":    1,
			"price":   1.06565,
			"account": "my_test",
			"":        "",
		}).Byte()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestOrderTotal(t *testing.T) {
	_, bytes, err := urllib.Post("http://127.0.0.1:8475/api/v1/positions_total").Byte()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestPositionsGet(t *testing.T) {
	_, bytes, err := urllib.Post("http://127.0.0.1:8475/api/v1/positions_get").
		SetJsonObject(map[string]interface{}{
			"symbol":  "EURUSD",
			"account": "my_test",
		}).Byte()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestAccountInfo(t *testing.T) {
	_, bytes, err := urllib.Post("http://127.0.0.1:8475/api/v1/account_info").
		SetJsonObject(map[string]interface{}{
			"account": "my_test",
		}).Byte()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
