package trading

import "fmt"

func (cb *CoinbaseAPI) CreateOrderRaw() (interface{}, error) {
	var order interface{}
	url := "/orders"
	method := "POST"

	_, err := cb.Request(method, url, nil, &order)

	if err != nil {
		return nil, fmt.Errorf("CreateOrderRaw: %w", err)
	}

	return order, err
}

func (cb *CoinbaseAPI) CreateOrder() (interface{}, error) {
	order, err := cb.CreateOrderRaw()

	if err != nil {
		return nil, fmt.Errorf("CreateOrder: %w", err)
	}

	return order, nil
}

func (cb *CoinbaseAPI) MarketBuy() (interface{}, error) {
	return nil, nil
}

func (cb *CoinbaseAPI) MarketSell() (interface{}, error) {
	return nil, nil
}
