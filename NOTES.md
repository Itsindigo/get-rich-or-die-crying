

A market pair is a pair of assets, the first currency is the base, and the second currency is the quote:

[Base-Quote]
[BTC-GBP]
[ETH-GBP]

To buy a product, provide a quote size or a base size:

`{"order_configuration":{"market_market_ioc":{"quote_size":"9"}}}`)

{
  "client_order_id": "some-id-provided-by-me",
  "product_id": "ETH-GBP",
  "side": "BUY" / "SELL",
  "order_configuration": {
    "market_market_ioc": {
      "quote_size": "1000", // GBP (Only BUY)
      "base_size": 1 // ETH/BTC (Buy or SELL)
    }
  },
  "self_trade_prevention_id": "some-id",
  "leverage": "1.0", // no leverage, omit?
  "retail_portfolio_id": "default portfolio id", // omit?
}


<!-- How to sell £10 of Eth??? -->
<!-- Think need to query trade price and reconstruct in other direction, Create Convert Quote -->

<!-- Sell 0.00001 Eth -->
{
  "client_order_id": "some-id-provided-by-me",
  "product_id": "ETH-GBP",
  "side": "SELL",
  "order_configuration": {
    "market_market_ioc": {
      "base_size": "0.00001"
    }
  },
}

<!-- Buy £10 of Eth -->
{
  "client_order_id": "some-id-provided-by-me",
  "product_id": "ETH-GBP",
  "side": "BUY",
  "order_configuration": {
    "market_market_ioc": {
      "quote_size": "10"
    }
  },
}



<!-- RESPONSES -->

<!-- SUCCESSFUL ORDER -->
{
    "success": true,
    "failure_reason": "UNKNOWN_FAILURE_REASON",
    "order_id": "81dab6fb-64aa-46b0-9151-90c92f7f7be0",
    "success_response": {
        "order_id": "81dab6fb-64aa-46b0-9151-90c92f7f7be0",
        "product_id": "ETH-GBP",
        "side": "SELL",
        "client_order_id": "some-id-provided-by-me"
    },
    "order_configuration": {
        "market_market_ioc": {
            "base_size": "0.00001"
        }
    }
}

<!-- UNSUCCESSFUL ORDER -->
{
    "success": false,
    "failure_reason": "UNKNOWN_FAILURE_REASON",
    "order_id": "",
    "error_response": {
        "error": "UNKNOWN_FAILURE_REASON",
        "message": "",
        "error_details": "",
        "preview_failure_reason": "PREVIEW_INVALID_BASE_SIZE_TOO_LARGE"
    },
    "order_configuration": {
        "market_market_ioc": {
            "base_size": "100000.00001"
        }
    }
}