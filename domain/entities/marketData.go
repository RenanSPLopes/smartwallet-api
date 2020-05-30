package entities

import "strings"

type MarketData struct {
	Name           string
	Sector         string
	SubSector      string
	Segmentation   string
	B3Segmentation string
	TagAlong       string
	FreeFloat      string
	Stocks         []Stock
	Results        []Result
}

type Stock struct {
	Code             string
	Type             string
	Quotes           float32
	MarketIndicators []MarketIndicators
}

type BalanceSheet struct {
	TotalAsset float64
	NetEquity  float64
	GrossDebt  float64
	Cash       float64
	NetDebt    float64
}

type OperatingResult struct {
	NetIncome                   float64
	BookBalance                 float64
	EBITDA                      float64
	DepreciationAndAmortization float64
	EBIT                        float64
	NetProfit                   float64
}

type Result struct {
	Date                string
	BalanceSheet        BalanceSheet
	OperatingResults    OperatingResult
	Market              Market
	FinancialIndicators FinancialIndicators
}

type Market struct {
	MarketValue     float64
	EnterpriseValue float64
	StocksCount     float64
}

type MarketIndicators struct {
	PriceEarningsRatio float32
	PriceAssetValue    float32
	PriceEBITDA        float32
	PriceEBIT          float32
}

type FinancialIndicators struct {
	MarginEBITDA  float32
	MarginEBIT    float32
	NetMargin     float32
	ROE           float32
	DebitToEBITDA float32
	DebitToEBIT   float32
}

func (m *MarketData) CalculateResultIndicators() {
	var results []Result

	r := m.Results[0]
	r.FinancialIndicators = FinancialIndicators{
		MarginEBITDA:  r.OperatingResults.calculateMarginEBITDA(),
		MarginEBIT:    r.OperatingResults.calculateMarginEBIT(),
		NetMargin:     r.OperatingResults.calculateNetMargin(),
		ROE:           r.OperatingResults.calculateROE(r.BalanceSheet.NetEquity),
		DebitToEBITDA: r.OperatingResults.calculateDebitToEBITDA(r.BalanceSheet.NetDebt),
		DebitToEBIT:   r.OperatingResults.calculateDebitToEBIT(r.BalanceSheet.NetDebt),
	}
	results = append(results, r)

	m.Results = results
}

func (m *MarketData) CalculateStocksIndicators() {
	var stocks []Stock
	r := m.Results[0]

	for _, s := range m.Stocks {
		marketIndicators := MarketIndicators{
			PriceEarningsRatio: s.calculePriceEarningsRatio(r.OperatingResults.NetProfit, r.Market.StocksCount),
			PriceAssetValue:    s.calculatePriceAssetValue(r.BalanceSheet.NetEquity, r.Market.StocksCount),
			PriceEBITDA:        s.calculatePriceEBITDA(r.OperatingResults.EBITDA, r.Market.StocksCount),
			PriceEBIT:          s.calculatePriceEBIT(r.OperatingResults.EBIT, r.Market.StocksCount),
		}

		s.Code = strings.ToUpper(s.Code)
		s.MarketIndicators = append(s.MarketIndicators, marketIndicators)
		stocks = append(stocks, s)
	}

	m.Stocks = stocks
}

func (s Stock) calculePriceEarningsRatio(netProfit float64, stocksCount float64) float32 {
	return s.Quotes / float32(netProfit/stocksCount)
}

func (s Stock) calculatePriceAssetValue(netEquity float64, stocksCount float64) float32 {
	return s.Quotes / float32(netEquity/stocksCount)
}

func (s Stock) calculatePriceEBITDA(ebitda float64, stocksCount float64) float32 {
	return s.Quotes / float32(ebitda/stocksCount)
}

func (s Stock) calculatePriceEBIT(ebit float64, stocksCount float64) float32 {
	return s.Quotes / float32(ebit/stocksCount)
}

func (r OperatingResult) calculateROE(netEquity float64) float32 {
	return float32(r.NetProfit / netEquity)
}

func (r OperatingResult) calculateDebitToEBITDA(netDebt float64) float32 {
	return float32(netDebt / r.EBITDA)
}

func (r OperatingResult) calculateDebitToEBIT(netDebt float64) float32 {
	return float32(netDebt / r.EBIT)
}

func (r OperatingResult) calculateMarginEBITDA() float32 {
	return float32(r.NetIncome / r.EBITDA)
}

func (r OperatingResult) calculateMarginEBIT() float32 {
	return float32(r.NetIncome / r.EBIT)
}

func (r OperatingResult) calculateNetMargin() float32 {
	return float32(r.NetProfit / r.NetIncome)
}
