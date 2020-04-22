package services

import (
	"log"
	"smartwallet-api/application/models"
	mapper "github.com/PeteProgrammer/go-automapper"
	"smartwallet-api/domain/entities"
)

type MarketDataProcessor interface{
	Process(marketData models.MarketData)
}

type MarketDataProcessorService struct{

}

func NewMarketDataProcessorService() MarketDataProcessorService {
	return MarketDataProcessorService{}
}

func (marketDataProcessor MarketDataProcessorService) Process(marketDataModel models.MarketData){
	var marketData entities.MarketData
	mapper.Map(marketDataModel, &marketData)

}