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
	MarketValue    float64
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
