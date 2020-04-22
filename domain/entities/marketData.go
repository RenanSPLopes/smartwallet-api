package entities

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
	Code string `json:"id"`
	Type string 
	Quotes float32 
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
}

type Market struct {
	MarketValue     float64
	EnterpriseValue float64
	Stocks          float64
}

func (m MarketData) CalculePriceEarningsRatio(){
	return m.Stock.Quotes / (m.Result.NetProfit / m.Market.Stocks)
}

func (m MarketData) CalculatePriceEBIT(){
	return m.Stock.Quotes/ (m.Result.EBIT / m.Market.Stocks)
}

func (m MarketData) CalculatePriceEBITDA(){
	return m.Stock.Quotes/ (m.Result.EBITDA / m.Market.Stocks)
}

func (m MarketData) CalculatePriceAssetValue(){
	return m.StockQuotes/ (m.BalanceSheet.NetEquity/ m.Market.Stocks)
}

func (m MarketData) CalculateROE(){
	return m.Result.NetProfit/m.BalanceSheet.NetEquity
}

func (m MarketData) CalculateDebitToPriceEarningsRatio(){
	return m.BalanceSheet.NetDebt/ m.CalculateDebitToPriceEarningsRatio()
}

func (m MarketData) CalculateDebitToEBITDA(){
	return m.BalanceSheet.NetDebt/ m.Result.EBITDA
}

func (m MarketData) CalculateMarginEBITDA(){
	return m.Result.NetIncome/m.Result.EBITDA
}

func (m MarketData) CalculateMarginEBIT(){
	return m.Result.NetIncome/m.Result.EBIT
}

func (m MarketData) CalculateNetMargin(){
	return m.Result.NetProfit/m.Result.NetIncome
}