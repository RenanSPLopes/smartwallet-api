package services

import (
	"smartwallet-api/application/models"
	"smartwallet-api/infrastructure/repositories"
)

type StocksQuotationProcessor interface {
	Process(stocksQuote models.StocksQuote)
}

type StocksQuotationProcessorService struct {
	MarketDataRepository repositories.MarketDataRepository
}

func NewStocksQuotationProcessorService(m repositories.MarketDataRepository) StocksQuotationProcessorService {
	return StocksQuotationProcessorService{MarketDataRepository: m}
}

func (s StocksQuotationProcessorService) Process(stocksQuote models.StocksQuote) {
	s.MarketDataRepository.UpdateQuotes(stocksQuote.Code, stocksQuote.Quote)
}
