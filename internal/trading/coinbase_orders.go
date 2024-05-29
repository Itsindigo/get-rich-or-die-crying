package trading

import (
	"fmt"
	"strconv"
	"time"
)

type SuccessResponse struct {
	OrderID       string `json:"order_id"`
	ProductID     string `json:"product_id"`
	Side          string `json:"side"`
	ClientOrderID string `json:"client_order_id"`
}

type ErrorResponse struct {
	Error                string `json:"error"`
	Message              string `json:"message"`
	ErrorDetails         string `json:"error_details"`
	PreviewFailureReason string `json:"preview_failure_reason"`
}

type MarketMarketIOC struct {
	BaseSize  string `json:"base_size,omitempty"`
	QuoteSize string `json:"quote_size,omitempty"`
}

/*
This type is incomplete but sufficient for market sales, full docs @
https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_postorder/
*/
type OrderConfiguration struct {
	MarketMarketIOC MarketMarketIOC `json:"market_market_ioc"`
}

type CreateOrderResponse struct {
	Success            bool               `json:"success"`
	FailureReason      string             `json:"failure_reason"`
	OrderID            string             `json:"order_id"`
	SuccessResponse    *SuccessResponse   `json:"success_response,omitempty"`
	ErrorResponse      *ErrorResponse     `json:"error_response,omitempty"`
	OrderConfiguration OrderConfiguration `json:"order_configuration"`
}

func (cb *CoinbaseAPI) CreateOrderRaw(body OrderOptions) (CreateOrderResponse, error) {
	var order CreateOrderResponse
	url := "/orders"
	method := "POST"

	_, err := cb.Request(method, url, body, &order)

	if err != nil {
		return CreateOrderResponse{}, fmt.Errorf("CreateOrderRaw: %w", err)
	}

	return order, err
}

func (cb *CoinbaseAPI) CreateOrder(orderOptions OrderOptions) (CreateOrderResponse, error) {
	order, err := cb.CreateOrderRaw(orderOptions)

	if err != nil {
		return CreateOrderResponse{}, fmt.Errorf("CreateOrder: %w", err)
	}

	if order.ErrorResponse != nil {
		errors := []ErrorWithLabel{
			{Label: "type", Error: order.ErrorResponse.Error},
			{Label: "message", Error: order.ErrorResponse.Message},
			{Label: "details", Error: order.ErrorResponse.ErrorDetails},
			{Label: "preview_failure_reason", Error: order.ErrorResponse.PreviewFailureReason},
		}
		return CreateOrderResponse{}, CreateOrderError{errors}
	}

	return order, nil
}

func (cb *CoinbaseAPI) MarketBuy(productId MarketPair, purchaseAmount string) (CreateOrderResponse, error) {
	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

	fmt.Printf("PURCHASING %s", purchaseAmount)

	order, err := cb.CreateOrder(
		OrderOptions{
			ClientOrderId: timestamp,
			ProductId:     productId,
			Side:          TradeSide("BUY"),
			OrderConfiguration: OrderConfiguration{
				MarketMarketIOC: MarketMarketIOC{
					QuoteSize: purchaseAmount,
				},
			},
		},
	)

	if err != nil {
		return CreateOrderResponse{}, fmt.Errorf("MarketSell - %w", err)
	}

	return order, nil
}

type OrderOptions struct {
	ClientOrderId      string             `json:"client_order_id"`
	ProductId          MarketPair         `json:"product_id"`
	Side               TradeSide          `json:"side"`
	OrderConfiguration OrderConfiguration `json:"order_configuration"`
}

func (cb *CoinbaseAPI) MarketSell(productId MarketPair, saleAmount string) (CreateOrderResponse, error) {
	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

	order, err := cb.CreateOrder(
		OrderOptions{
			ClientOrderId: timestamp,
			ProductId:     productId,
			Side:          TradeSide("SELL"),
			OrderConfiguration: OrderConfiguration{
				MarketMarketIOC: MarketMarketIOC{
					BaseSize: saleAmount,
				},
			},
		},
	)

	if err != nil {
		return CreateOrderResponse{}, fmt.Errorf("MarketSell - %w", err)
	}

	return order, nil
}
