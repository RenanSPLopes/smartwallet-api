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

type Result struct {
	Date                string
	BalanceSheet        BalanceSheet
	OperatingResults    OperatingResult
	FinancialIndicators FinancialIndicators
	CashFlow            CashFlow
}

//Balanço Patrimonial
type BalanceSheet struct {
	TotalAsset            float64 // Ativo Total
	CurrentAssets         float64 // Ativo Circulante
	NonCurrentAssets      float64 // Ativo não Circulante
	TotalLiabilities      float64 // Passivo Total
	CurrentLiabilities    float64 // Passivo Circulante
	NonCurrentLiabilities float64 // Passivo não Circulante
	NetEquity             float64 // Patrimônio Líquido
	GrossDebt             float64
	Cash                  float64
	NetDebt               float64
}

type OperatingResult struct {
	NetIncome        float64 //Receita Líquida
	BookBalance      float64 // Receita Operacional Bruta
	EBITDA           *float64
	NetProfit        float64 // Lucro Líquido
	EquityValue      float64 // Valor Patrimonial
	Capex            float64
	PDD              *float64
	ServiceRevenue   *float64
	EarningsPerShare float64 //Lucro por acao
}

type CashFlow struct {
	FreeCashFlow       float64 // Fluxo de Caixa Livre
	OperatingCashFlow  float64 // Fluxo de Caixa Operacional
	InvestmentCashFlow float64 // Fluxo de Caixa de Investimento
	FinancingCashFlow  float64 // Fluxo de Caixa de Financiamento
}

type FinancialIndicators struct {
	MarginEBITDA  float32
	NetMargin     float32
	ROE           float32
	DebitToEBITDA float32
	ROA           float32
}

func (m *MarketData) CalculateResultIndicators() {
	var results []Result

	for _, r := range m.Results {
		r.FinancialIndicators = FinancialIndicators{
			MarginEBITDA:  r.OperatingResults.calculateMarginEBITDA(),
			NetMargin:     r.OperatingResults.calculateNetMargin(),
			ROE:           r.OperatingResults.calculateROE(r.BalanceSheet.NetEquity),
			DebitToEBITDA: r.OperatingResults.calculateDebitToEBITDA(r.BalanceSheet.NetDebt),
			ROA:           r.OperatingResults.calculateROA(r.BalanceSheet.TotalAsset),
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

func (r OperatingResult) calculateROA(totalAsset float64) float32 {

	if totalAsset == 0 {
		return 0
	}

	return float32(r.NetProfit / totalAsset)
}

func (r OperatingResult) calculateDebitToEBITDA(netDebt float64) float32 {

	if *r.EBITDA == 0 {
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
