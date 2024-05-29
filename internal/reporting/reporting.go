package reporting

type TradeReporter struct{}

func (tr *TradeReporter) ReportNoAction() {

}

func (tr *TradeReporter) ReportSale() {

}

func (tr *TradeReporter) ReportBuy() {

}

func (tr *TradeReporter) ReportError(err error) {

}

func NewTradeReporter() *TradeReporter {
	return &TradeReporter{}
}
