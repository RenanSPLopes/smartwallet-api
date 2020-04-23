package services

import (
	"smartwallet-api/application/models"
	mapper "github.com/PeteProgrammer/go-automapper"
	"smartwallet-api/domain/entities"
	"smartwallet-api/infrastructure/repositories"
	"log"
	"encoding/json"
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
	// marketData.SetIndicators()
	mappedMarketData := &marketData
	mappedMarketData.SetIndicators()

	jsonObject, _ := json.Marshal(mappedMarketData)
	log.Printf("##################################")
	log.Printf("Sending to Repository: " + string(jsonObject))
	log.Printf("##################################")
	m.MarketDataRepository.Save(*mappedMarketData)
}