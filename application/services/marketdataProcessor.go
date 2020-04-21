package services

import (
	"log"
	"smartwallet-api/application/models"
)

type MarketDataProcessor interface{
	Process(marketData models.MarketData)
}

type MarketDataProcessorService struct{

}

func NewMarketDataProcessorService() MarketDataProcessorService {
	return MarketDataProcessorService{}
}

func (marketDataProcessor MarketDataProcessorService) Process(marketData models.MarketData){
	log.Printf("Message received.")
}