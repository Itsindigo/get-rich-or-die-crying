# About

This repo is a project for automating crypto trades via [coinbase.com](https://coinbase.com). It uses a simple decision tree that parses the [CoinMarketCap](https://coinmarketcap.com/) fear and greed index to determine whether to buy, sell or hold.


## Flow

![Job Design Diagram](./docs/diagram.png)


## Technology Choices
- Go
- Colly for scraping
- Coinbase Advanced Trade API
- [modal.com](https://modal.com/) as platform to host cronjobs
- Twilio for texts?

