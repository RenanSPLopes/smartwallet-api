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

type MarketIndicators struct{
	PriceEarningsRatio float64
	PriceAssetValue float64
	PriceEBITDA float64
	PriceEBIT float64
}

type FinancialIndicators struct{
	MarginEBITDA float32
	MarginEBIT float32
	NetMargin float32
	ROE float32
	DebitToPriceEarningsRation float32
	DebitToEBITDA float32
}