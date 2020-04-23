package services

import (
	"smartwallet-api/application/models"
	mapper "github.com/PeteProgrammer/go-automapper"
	"smartwallet-api/domain/entities"
	"smartwallet-api/infrastructure/repositories"
)

type MarketDataProcessor interface{
	Process(marketData models.MarketData)
}

type MarketDataProcessorService struct{
	MarketDataRepository repositories.MarketDataRepository
}

func NewMarketDataProcessorService(m repositories.MarketDataRepository) MarketDataProcessorService {
	return MarketDataProcessorService{MarketDataRepository: m}
}

func (m MarketDataProcessorService) Process(marketDataModel models.MarketData){
	var marketData entities.MarketData
	mapper.MapLoose(marketDataModel, &marketData)
	marketData.SetIndicators()
	m.MarketDataRepository.Save(marketData)
}