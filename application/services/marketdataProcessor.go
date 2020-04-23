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
	marketData := &entities.MarketData{
		Name: marketDataModel.Name,
		Sector: marketDataModel.Sector,
		SubSector: marketDataModel.SubSector,
		Segmentation: marketDataModel.Segmentation,
		B3Segmentation: marketDataModel.B3Segmentation,
		TagAlong: marketDataModel.TagAlong,
		FreeFloat: marketDataModel.FreeFloat,
		Stocks: mapStockFromModel(marketDataModel.Stocks),
		BalanceSheet: entities.BalanceSheet{
			TotalAsset: marketDataModel.BalanceSheet.TotalAsset,
			NetEquity: marketDataModel.BalanceSheet.NetEquity,
			GrossDebt: marketDataModel.BalanceSheet.GrossDebt,
			Cash: marketDataModel.BalanceSheet.Cash,
			NetDebt: marketDataModel.BalanceSheet.NetDebt,
		},
		Results: mapResultFromModel(marketDataModel.Results),
		Market: mapMarketFromModel(marketDataModel.Market),
	}

	jsonObject, _ := json.Marshal(marketData)
	log.Printf("##################################")
	log.Printf("Sending to Repository: " + string(jsonObject))
	log.Printf("##################################")
	m.MarketDataRepository.Save(*marketData)
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