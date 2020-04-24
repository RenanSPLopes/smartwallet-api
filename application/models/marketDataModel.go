package models

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
	Code string 
	Type string 
	Quotes float32 
}


type Result struct {
	Date                string               
	BalanceSheet   		BalanceSheet         
	OperatingResults    OperatingResult     
	Market         		Market               
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

type Market struct {
	MarketValue     	float64 
	EnterpriseValue 	float64 
	StocksCount         float64 
}