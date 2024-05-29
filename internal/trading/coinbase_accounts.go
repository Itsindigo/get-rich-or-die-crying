package trading

import (
	"fmt"
	"slices"
	"time"
)

type Balance struct {
	Currency string `json:"currency"`
	Value    string `json:"value"`
}

type Account struct {
	Active            bool       `json:"active"`
	AvailableBalance  Balance    `json:"available_balance"`
	CreatedAt         time.Time  `json:"created_at"`
	Currency          string     `json:"currency"`
	Default           bool       `json:"default"`
	DeletedAt         *time.Time `json:"deleted_at"`
	Hold              Balance    `json:"hold"`
	Name              string     `json:"name"`
	Ready             bool       `json:"ready"`
	RetailPortfolioID string     `json:"retail_portfolio_id"`
	Type              string     `json:"type"`
	UpdatedAt         time.Time  `json:"updated_at"`
	UUID              string     `json:"uuid"`
}

type GetAccountsResponse struct {
	Accounts []Account `json:"accounts"`
	Cursor   string    `json:"cursor"`
	HasNext  bool      `json:"has_next"`
	Size     int       `json:"size"`
}

func (cb *CoinbaseAPI) AccountsRaw() (GetAccountsResponse, error) {
	var accounts GetAccountsResponse
	url := "/accounts"
	method := "GET"

	_, err := cb.Request(method, url, nil, &accounts)

	if err != nil {
		return GetAccountsResponse{}, fmt.Errorf("AccountsRaw: %w", err)
	}

	return accounts, err
}

func (ar *GetAccountsResponse) ToSimpleAccounts(currencies []CoinbaseWalletName) []SimpleAccount {
	var simpleAccounts []SimpleAccount

	for _, account := range ar.Accounts {
		if slices.Contains(currencies, CoinbaseWalletName(account.Name)) {
			simpleAccounts = append(simpleAccounts, SimpleAccount{
				Id:        account.UUID,
				Name:      CoinbaseWalletName(account.Name),
				Currency:  account.Currency,
				Balance:   account.AvailableBalance.Value,
				Type:      account.Type,
				IsActive:  account.Active,
				IsDefault: account.Default,
				IsReady:   account.Ready,
				CreatedAt: account.CreatedAt,
				UpdatedAt: account.UpdatedAt,
				DeletedAt: account.DeletedAt,
			})
		}
	}

	return simpleAccounts
}

func (cb *CoinbaseAPI) GetWallets(wallets []CoinbaseWalletName) ([]SimpleAccount, error) {
	ar, err := cb.AccountsRaw()

	if err != nil {
		return nil, fmt.Errorf("GetWallets: %w", err)
	}

	return ar.ToSimpleAccounts(wallets), nil
}
