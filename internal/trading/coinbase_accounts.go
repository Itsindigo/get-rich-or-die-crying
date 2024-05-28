package trading

import (
	"fmt"
	"time"

	"github.com/itsindigo/get-rich-or-die-crying/internal/utils"
)

type Balance struct {
	Currency string  `json:"currency"`
	Value    float64 `json:"value,string"`
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

type AccountsResponse struct {
	Accounts []Account `json:"accounts"`
	Cursor   string    `json:"cursor"`
	HasNext  bool      `json:"has_next"`
	Size     int       `json:"size"`
}

func (cb *CoinbaseAPI) AccountsRaw() (*AccountsResponse, error) {
	var accounts AccountsResponse
	url := "/accounts"
	method := "GET"

	_, err := cb.Request(method, url, nil, &accounts)

	if err != nil {
		return nil, fmt.Errorf("accounts request error: %v", err)
	}

	return &accounts, err
}

func (ar *AccountsResponse) ToSimpleAccounts(currencies []string) []SimpleAccount {
	var simpleAccounts []SimpleAccount

	for _, account := range ar.Accounts {
		if utils.StringInSlice(account.Name, currencies) {
			fmt.Println("BALANCE:")
			fmt.Println(account.AvailableBalance.Value)
			simpleAccounts = append(simpleAccounts, SimpleAccount{
				Id:        account.UUID,
				Name:      account.Name,
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

func (cb *CoinbaseAPI) GetWallets(wallets []string) ([]SimpleAccount, error) {
	ar, err := cb.AccountsRaw()

	if err != nil {
		return nil, fmt.Errorf("error fetching accounts: %v", err)
	}

	return ar.ToSimpleAccounts(wallets), nil
}
