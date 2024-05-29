package trading

import (
	"fmt"
	"time"
)

type SimpleAccount struct {
	Id        string
	Name      CoinbaseWalletName
	Currency  string
	IsActive  bool
	IsDefault bool
	IsReady   bool
	Type      string
	Balance   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (sa SimpleAccount) GoodToTrade() bool {
	return sa.IsActive && sa.IsReady && sa.DeletedAt == nil
}

func (sa SimpleAccount) String() string {
	deletedAtStr := "nil"
	if sa.DeletedAt != nil {
		deletedAtStr = sa.DeletedAt.Format(time.RFC3339)
	}

	return fmt.Sprintf(
		"SimpleAccount {\n"+
			"  Id: %q\n"+
			"  Name: %q\n"+
			"  Currency: %q\n"+
			"  IsActive: %t\n"+
			"  IsDefault: %t\n"+
			"  IsReady: %t\n"+
			"  Type: %s\n"+
			"  Balance: %s\n"+
			"  CreatedAt: %s\n"+
			"  UpdatedAt: %s\n"+
			"  DeletedAt: %s\n"+
			"}\n",
		sa.Id,
		sa.Name,
		sa.Currency,
		sa.IsActive,
		sa.IsDefault,
		sa.IsReady,
		sa.Type,
		sa.Balance,
		sa.CreatedAt.Format(time.RFC3339),
		sa.UpdatedAt.Format(time.RFC3339),
		deletedAtStr,
	)
}
