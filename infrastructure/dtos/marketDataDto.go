package dtos

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
	StocksCount     float64
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