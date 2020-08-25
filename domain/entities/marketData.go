package entities

import "strings"

type MarketData struct {
	Name           string
	Sector         string
	SubSector      string
	Segmentation   string
	B3Segmentation string
	MarketValue    float64
	Stocks         []Stock
	Results        []Result
}

type Stock struct {
	Code   string
	Type   string
	Quotes float32
}

type BalanceSheet struct {
	TotalAsset                     float64
	NetEquity                      float64
	GrossDebt                      float64
	Cash                           float64
	NetDebt                        float64
	FinancialIntermediationRevenue float64
}

type OperatingResult struct {
	NetIncome      float64
	BookBalance    float64
	EBITDA         *float64
	NetProfit      float64
	EquityValue    float64
	Capex          float64
	PDD            *float64
	ServiceRevenue *float64
}

type Result struct {
	Date                string
	BalanceSheet        BalanceSheet
	OperatingResults    OperatingResult
	FinancialIndicators FinancialIndicators
}

type FinancialIndicators struct {
	MarginEBITDA  float32
	NetMargin     float32
	ROE           float32
	DebitToEBITDA float32
}

func (m *MarketData) CalculateResultIndicators() {
	var results []Result

	for _, r := range m.Results {
		r.FinancialIndicators = FinancialIndicators{
			MarginEBITDA:  r.OperatingResults.calculateMarginEBITDA(),
			NetMargin:     r.OperatingResults.calculateNetMargin(),
			ROE:           r.OperatingResults.calculateROE(r.BalanceSheet.NetEquity),
			DebitToEBITDA: r.OperatingResults.calculateDebitToEBITDA(r.BalanceSheet.NetDebt),
		}
		results = append(results, r)
	}
	m.Results = results
}

func (m *MarketData) CalculateStocksIndicators() {
	var stocks []Stock

	for _, s := range m.Stocks {

		s.Code = strings.ToUpper(s.Code)
		stocks = append(stocks, s)
	}

	m.Stocks = stocks
}

func (r OperatingResult) calculateROE(netEquity float64) float32 {

	if netEquity == 0 {
		return 0
	}

	return float32(r.NetProfit / netEquity)
}

func (r OperatingResult) calculateDebitToEBITDA(netDebt float64) float32 {

	if r.EBITDA == nil {
		return 0
	}

	return float32(netDebt / *r.EBITDA)
}

func (r OperatingResult) calculateMarginEBITDA() float32 {

	if r.NetIncome == 0 {
		return 0
	}

	return float32(*r.EBITDA / r.NetIncome)
}

func (r OperatingResult) calculateNetMargin() float32 {

	if r.NetIncome == 0 {
		return 0
	}

	return float32(r.NetProfit / r.NetIncome)
}
