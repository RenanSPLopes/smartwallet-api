package models

type MarketData struct {
	Name           string       `json:"name"`
	Sector         string       `json:"sector"`
	SubSector      string       `json:"subSector"`
	Segmentation   string       `json:"segmentation"`
	B3Segmentation string       `json:"b3Segmentation"`
	TagAlong       string       `json:"tagAlong"`
	FreeFloat      string       `json:"freeFloat"`
	Stocks         []Stock      `json:"stocks"`
	BalanceSheet   BalanceSheet `json:"balanceSheet"`
	Results        []Result     `json:"results"`
	Market         Market       `json:"market"`
}

type Stock struct {
	Code string `json:"code"`
	Type string `json:"type"`
	Quotes float32 `json:"quotes"`
}

type BalanceSheet struct {
	TotalAsset float64 `json:"totalAsset"`
	NetEquity  float64 `json:"netEquity"`
	GrossDebt  float64 `json:"grossDebt"`
	Cash       float64 `json:"cash"`
	NetDebt    float64 `json:"netDebt"`
}

type Result struct {
	Date                        string  `json:"date"`
	NetIncome                   float64 `json:"netIncome"`
	BookBalance                 float64 `json:"bookBalance"`
	EBITDA                      float64 `json:"ebitda"`
	DepreciationAndAmortization float64 `json:"depreciationAndAmortization"`
	EBIT                        float64 `json:"ebit"`
	NetProfit                   float64 `json:"netProfit"`
}

type Market struct {
	MarketValue     float64 `json:"marketValue"`
	EnterpriseValue float64 `json:"enterpriseValue"`
	Stocks          float64 `json:"stocks"`
}
