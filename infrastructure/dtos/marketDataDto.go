package dtos

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MarketData struct {
	ID             primitive.ObjectID `bson:"_id"`
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

type BalanceSheet struct {
	TotalAsset                     float64
	CurrentAssets                  float64 // Ativo Circulante
	NonCurrentAssets               float64 // Ativo não Circulante
	TotalLiabilities               float64 // Passivo Total
	CurrentLiabilities             float64 // Passivo Circulante
	NonCurrentLiabilities          float64 // Passivo não Circulante
	NetEquity                      float64
	GrossDebt                      float64
	Cash                           float64
	NetDebt                        float64
	FinancialIntermediationRevenue float64
}

type Result struct {
	Date                string
	BalanceSheet        BalanceSheet
	OperatingResults    OperatingResult
	FinancialIndicators FinancialIndicators
	CashFlow            CashFlow
}

type OperatingResult struct {
	NetIncome        float64
	BookBalance      float64
	EBITDA           *float64
	NetProfit        float64
	EquityValue      float64
	Capex            float64
	PDD              *float64
	ServiceRevenue   *float64
	EarningsPerShare float64
}

type CashFlow struct {
	FreeCashFlow       float64 // Fluxo de Caixa Livre
	OperatingCashFlow  float64 //Fluxo de Caixa Operacional
	InvestmentCashFlow float64 // Fluxo de Caixa de Investimento
	FinancingCashFlow  float64 // Fluxo de Caixa de Financiamento
}

type MarketIndicators struct {
	PriceEarningsRatio float32
	PriceAssetValue    float32
	PriceEBITDA        float32
}

type FinancialIndicators struct {
	MarginEBITDA  float32
	NetMargin     float32
	ROE           float32
	DebitToEBITDA float32
	ROA           float32
}
