package models

type BankMarketData struct {
	Name           string
	Sector         string
	SubSector      string
	Segmentation   string
	B3Segmentation string
	MarketValue    float64
	Stocks         []Stock
	Results        []BankResult
}

type BankResult struct {
	Date             string
	BalanceSheet     BankBalanceSheet
	OperatingResults BankOperatingResult
}

type BankBalanceSheet struct {
	NetEquity                      float64
	Cash                           float64
	FinancialIntermediationRevenue float64
}

type BankOperatingResult struct {
	BookBalance    float64
	NetProfit      float64
	EquityValue    float64
	PDD            float64
	ServiceRevenue float64 //Receita Serviços
}
