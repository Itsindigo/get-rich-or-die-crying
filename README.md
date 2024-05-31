# About

This repo is a project for automating crypto trades via [coinbase.com](https://coinbase.com). It uses a simple decision tree that parses the [CoinMarketCap](https://coinmarketcap.com/) fear and greed index to determine whether to buy, sell or hold.


## Flow

![Job Design Diagram](./docs/diagram.png)


## Technology Choices
- Go
- Colly for scraping
- Coinbase Advanced Trade API
- [northflank](https://northflank.com/) for CI/CD / Cronjob platform
- Slack API for notifications



## Deploy

- Build Docker Image: `docker image build -t get-rich-or-die-trying/send-it:latest -f ./cmd/send-it/Dockerfile .`
- `docker run -e IS_REMOTE_ENVIRONMENT='1' --env-file .env get-rich-or-die-trying/send-it:latest`