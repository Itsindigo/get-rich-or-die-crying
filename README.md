# Plan
  - Job one, run hourly:
    - Query the coinmarket cap fear and greed index each day and store the value in a database
  - Job two, run daily:
    - Evaluate whether to buy or sell:
      - If value < MIN_FEAR_SCORE (45?) in last 24 hours, spend all cash in coinbase account on crypto
      - if value > MAX_GREED_SCORE (80?) in last 24 hours sell all crypto in coinbase account

# Essential components
- Job to scrape the fear and greed index score
- Job to evaluate whether to buy or sell and act if conditions are met
- Send summary SMS notification each day with current fear and greed index score, whether or not any transactions were made, or if there was an error

# Technology Choices
- Go
- Postgres
- Scraping lib: unknown
- [modal.com](https://modal.com/) as platform to host cronjobs

# Data sources
- Coinbase API (transactions, balances)
- Coinmarketcap (fear and greed index / market data)
