package services

import (
	"log"
	"smartwallet-api/application/models"
	"smartwallet-api/domain/entities"
	"smartwallet-api/infrastructure/dtos"
	"smartwallet-api/infrastructure/repositories"
	"smartwallet-api/utils"

	mapper "github.com/PeteProgrammer/go-automapper"
)

type MarketDataProcessor interface {
	Process(marketData models.MarketData)
}

type MarketDataProcessorService struct {
	MarketDataRepository repositories.MarketDataRepository
}

func NewMarketDataProcessorService(m repositories.MarketDataRepository) MarketDataProcessorService {
	return MarketDataProcessorService{MarketDataRepository: m}
}

func (m MarketDataProcessorService) Process(marketDataModel models.MarketData) {
	var marketData entities.MarketData
	mapper.MapLoose(marketDataModel, &marketData)
	marketData.CalculateResultIndicators()

	marketDataFromDb := m.MarketDataRepository.GetByCode(marketData.Stocks[0].Code)

	if marketDataFromDb.Name != "" {
		result := marketData.Results[0]
		dates := extractResultsDates(marketDataFromDb.Results)

		if utils.Contains(dates, marketData.Results[0].Date) {
			log.Println("Result already exist on database.")
			return
		}

		m.MarketDataRepository.UpdateResults(marketDataFromDb.ID, result)
		return
	}

	marketData.CalculateStocksIndicators()
	m.MarketDataRepository.Save(marketData)
}

func extractResultsDates(results []dtos.Result) []string {
	var dates []string
	for _, x := range results {
		dates = append(dates, x.Date)
	}

	return dates
}
