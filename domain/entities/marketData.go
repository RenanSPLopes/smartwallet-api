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

func (m MarketData) CalculePriceEarningsRatios() []float32{
	var priceEarningsRatios []float32
	for _, s := range m.Stocks{
		for _, r := range m.Results {
			priceEarningsRatio := s.Quotes/ float32(r.NetProfit/ m.Market.Stocks)
			priceEarningsRatios = append(priceEarningsRatios, priceEarningsRatio)
		}
	}

	return priceEarningsRatios
}

func (m MarketData) CalculatePriceEBITs()[]float32{
	var priceEBITs []float32
	for _, s := range m.Stocks{
		for _, r := range m.Results{
			priceEBIT := s.Quotes/ float32(r.EBIT / m.Market.Stocks)
			priceEBITs = append(priceEBITs, priceEBIT)
		}
	}

	return priceEBITs
}

func (m MarketData) CalculatePriceEBITDAs()[]float32{
	var priceEBITDAs []float32
	for _, s := range m.Stocks{
		for _, r := range m.Results{
			priceEBITDA := s.Quotes/ float32(r.EBITDA / m.Market.Stocks)
			priceEBITDAs = append(priceEBITDAs, priceEBITDA)
		}
	}

	return priceEBITDAs
}

func (m MarketData) CalculatePriceAssetValues() []float32{
	var priceAssetValues []float32
	for _, s := range m.Stocks{
			priceAssetValue := s.Quotes/ float32(m.BalanceSheet.NetEquity/ m.Market.Stocks)
			priceAssetValues = append(priceAssetValues, priceAssetValue)
	}

	return priceAssetValues
}

func (m MarketData) CalculateROEs()[]float32{
	var roes []float32
	for _, r := range m.Results{
		roe := float32(r.NetProfit/m.BalanceSheet.NetEquity)
		roes = append(roes, roe)
	}

	return roes
}

func (m MarketData) CalculateDebitToPriceEarningsRatios() []float32{
	var debitPriceEarningsRatios []float32
	priceEarningsRatios := m.CalculePriceEarningsRatios()
	for _, per := range priceEarningsRatios{
		debitPriceEarningsRatio := float32(m.BalanceSheet.NetDebt)/ per
		debitPriceEarningsRatios = append(debitPriceEarningsRatios, debitPriceEarningsRatio)
	}

	return debitPriceEarningsRatios
	
}

func (m MarketData) CalculateDebitToEBITDA()[]float32{
	var debitToEBITDAs []float32
	for _, r := range m.Results{
		debitToEBITDA := float32(m.BalanceSheet.NetDebt/ r.EBITDA)
		debitToEBITDAs = append(debitToEBITDAs, debitToEBITDA)
	}

	return debitToEBITDAs
}

func (m MarketData) CalculateMarginEBITDA()[]float32{
	var marginEBITDAs []float32
	for _, r := range m.Results{
		marginEBITDA := float32(r.NetIncome/r.EBITDA)
		marginEBITDAs = append(marginEBITDAs, marginEBITDA)
	}

	return marginEBITDAs
}

func (m MarketData) CalculateMarginEBIT()[]float32{
	var marginEBITs []float32
	for _, r := range m.Results{
		marginEBIT := float32(r.NetIncome/r.EBIT)
		marginEBITs = append(marginEBITs, marginEBIT)
	}

	return marginEBITs
}

func (m MarketData) CalculateNetMargin()[]float32{
	var netMargins []float32
	for _, r := range m.Results{
		netMargin := float32(r.NetProfit/r.NetIncome)
		netMargins = append(netMargins, netMargin)
	}

	return netMargins
}