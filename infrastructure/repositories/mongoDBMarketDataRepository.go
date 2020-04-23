package repositories

import (
	"smartwallet-api/domain/entities"
	"smartwallet-api/infrastructure/dtos"
	mapper "github.com/PeteProgrammer/go-automapper"

)

type MarketDataRepository interface {
	Save(marketData entities.MarketData)
}

type MongoDBMarketDataRepository struct {
	ConnectionString string
	Collection string
}

func NewMongoDBMarketDataRepository(conectionString string, collection string) MongoDBMarketDataRepository{
	return MongoDBMarketDataRepository{ConnectionString:conectionString, Collection:collection}
}

func (m MongoDBMarketDataRepository) Save(marketData entities.MarketData){
	var marketDataDto dtos.MarketData
	mapper.Map(marketData, &marketDataDto)
}