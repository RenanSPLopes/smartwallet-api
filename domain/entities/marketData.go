package entities

import (
    "log"
    "encoding/json"
)

type MarketData struct {
	Name           string
	Sector         string
	SubSector      string
	Segmentation   string
	B3Segmentation string
	TagAlong       string
	FreeFloat      string
	Stocks         []Stock
	BalanceSheet   BalanceSheet
	Results        []Result
	Market         Market
}

type Stock struct {
	Code string 
	Type string 
	Quotes float32 
	MarketIndicators []MarketIndicators
}

type BalanceSheet struct {
	TotalAsset float64
	NetEquity  float64
	GrossDebt  float64
	Cash       float64
	NetDebt    float64
}

type Result struct {
	Date                        string
	NetIncome                   float64
	BookBalance                 float64
	EBITDA                      float64
	DepreciationAndAmortization float64
	EBIT                        float64
	NetProfit                   float64
	FinancialIndicators FinancialIndicators
}

type Market struct {
	MarketValue     float64
	EnterpriseValue float64
	StocksCount          float64
}

type MarketIndicators struct{
	PriceEarningsRatio float32
	PriceAssetValue float32
	PriceEBITDA float32
	PriceEBIT float32
}

type FinancialIndicators struct{
	MarginEBITDA float32
	MarginEBIT float32
	NetMargin float32
	ROE float32
	DebitToEBITDA float32
	DebitToEBIT float32
}

func (m *MarketData) SetIndicators(){
	var results []Results
	var stocks []Stock
	for _, r := range m.Results{
		for _, s := range m.Stocks{
			marketIndicators := MarketIndicators{
				PriceEarningsRatio: s.calculePriceEarningsRatio(r.NetProfit, m.Market.StocksCount),
				PriceAssetValue: s.calculatePriceAssetValue(m.BalanceSheet.NetEquity, m.Market.StocksCount),
				PriceEBITDA: s.calculatePriceEBITDA(r.EBITDA, m.Market.StocksCount),
				PriceEBIT: s.calculatePriceEBIT(r.EBIT, m.Market.StocksCount),
			}
			s.MarketIndicators = append(s.MarketIndicators, marketIndicators)
			stocks = append(stocks, s)		
		}
		r.FinancialIndicators = FinancialIndicators{
			MarginEBITDA: r.calculateMarginEBITDA(),
			MarginEBIT: r.calculateMarginEBIT(),
			NetMargin: r.calculateNetMargin(),
			ROE: r.calculateROE(m.BalanceSheet.NetEquity),
			DebitToEBITDA: r.calculateDebitToEBITDA(m.BalanceSheet.NetDebt),
			DebitToEBIT: r.calculateDebitToEBIT(m.BalanceSheet.NetDebt),
		}
		results = append(results, r)
	}

	m.Results = results
	m.Stocks = stocks

	jsonTest, _ := json.Marshal(m)
	log.Printf("After : " + string(jsonTest))
}

func (s Stock) calculePriceEarningsRatio(netProfit float64 ,  stocksCount float64) float32{
	return s.Quotes/ float32(netProfit/ stocksCount)
}

func (s Stock) calculatePriceAssetValue(netEquity float64, stocksCount float64) float32{
	return s.Quotes/ float32(netEquity/ stocksCount)
}

func (s Stock) calculatePriceEBITDA(ebitda float64, stocksCount float64) float32{
	return s.Quotes/ float32(ebitda / stocksCount)
}

func (s Stock) calculatePriceEBIT(ebit float64, stocksCount float64) float32{
	return s.Quotes/ float32(ebit / stocksCount)
}

func (r Result) calculateROE(netEquity float64)float32{
	return float32(r.NetProfit/netEquity)
}

func (r Result) calculateDebitToEBITDA(netDebt float64)float32{
	return float32(netDebt/ r.EBITDA)
}

func (r Result) calculateDebitToEBIT(netDebt float64)float32{
	return float32(netDebt/ r.EBIT)
}

func (r Result) calculateMarginEBITDA()float32{
	return float32(r.NetIncome/r.EBITDA)
}

func (r Result) calculateMarginEBIT()float32{
	return float32(r.NetIncome/r.EBIT)
}

func (r Result) calculateNetMargin() float32{
	return float32(r.NetProfit/r.NetIncome)
}


