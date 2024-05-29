package trading

var FearBuyThreshold int = 45
var GreedSellThreshold int = 85

type CoinbaseWalletName string

const (
	ETHWallet  CoinbaseWalletName = "ETH Wallet"
	BTCWallet  CoinbaseWalletName = "BTC Wallet"
	GBPWallet  CoinbaseWalletName = "GBP Wallet"
	MINAWallet CoinbaseWalletName = "MINA Wallet" // wallet with no balance
)
