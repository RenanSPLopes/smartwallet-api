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
	marketData.SetIndicators()

	jsonObject, _ := json.Marshal(marketData)
	log.Printf("##################################")
	log.Printf("Sending to Repository: " + string(jsonObject))
	log.Printf("##################################")
	m.MarketDataRepository.Save(marketData)
}

func mapMarketFromModel(marketModel models.Market) entities.Market{
	var market entities.Market
	mapper.Map(marketModel, &market)
	return market
}

func mapResultFromModel(resultsModel []models.Result) []entities.Result{
	var results []entities.Result
	for _, r := range resultsModel {
		var result entities.Result 
		 mapper.MapLoose(r, &result)
		 results = append(results, result)
	}

	return results
}

func mapStockFromModel(stocksModel []models.Stock) []entities.Stock{
	var stocks []entities.Stock
	for _, stockModel  := range stocksModel {
		stock := entities.Stock{
			Code: stockModel.Code,
			Type: stockModel.Type,
			Quotes: stockModel.Quotes,
		}

		stocks = append(stocks, stock)
	}
	
	return stocks
}