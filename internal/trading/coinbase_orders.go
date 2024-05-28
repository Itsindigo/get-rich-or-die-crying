package trading

import "fmt"

func (cb *CoinbaseAPI) CreateOrderRaw() (interface{}, error) {
	var order interface{}
	url := "/orders"
	method := "POST"

	_, err := cb.Request(method, url, nil, &order)

	if err != nil {
		return nil, fmt.Errorf("orders request error: %v", err)
	}

	return order, err
}

func (cb *CoinbaseAPI) CreateOrder() (interface{}, error) {
	order, err := cb.CreateOrderRaw()

	if err != nil {
		return nil, fmt.Errorf("error creating order: %v", err)
	}

	return order, nil
}

func (cb *CoinbaseAPI) MarketBuy() {}

func (cb *CoinbaseAPI) MarketSell() {}
