package trading

var FearBuyThreshold int = 45
var GreedSellThreshold int = 82

type CoinbaseWalletName string

const (
	ETHWallet  CoinbaseWalletName = "ETH Wallet"
	BTCWallet  CoinbaseWalletName = "BTC Wallet"
	GBPWallet  CoinbaseWalletName = "GBP Wallet"
	MINAWallet CoinbaseWalletName = "MINA Wallet" // wallet with no balance
)

type MarketPair string

const (
	ETH_GBP MarketPair = "ETH-GBP"
)

type TradeSide string

const (
	Buy  TradeSide = "BUY"
	Sell TradeSide = "SELL"
)
