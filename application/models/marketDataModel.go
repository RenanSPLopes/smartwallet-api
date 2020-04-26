package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MarketData struct {
	ID             primitive.ObjectID
	Name           string
	Sector         string
	SubSector      string
	Segmentation   string
	B3Segmentation string
	TagAlong       string    `json:",omitempty"`
	FreeFloat      string    `json:",omitempty"`
	Stocks         []*Stock  `json:",omitempty"`
	Results        []*Result `json:",omitempty"`
}

type Stock struct {
	Code   string
	Type   string
	Quotes float32
}

type Result struct {
	Date             string
	BalanceSheet     BalanceSheet
	OperatingResults OperatingResult
	Market           Market
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
	MarketValue     float64
	EnterpriseValue float64
	StocksCount     float64
}
