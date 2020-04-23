package repositories

import (
	"smartwallet-api/domain/entities"
	"smartwallet-api/infrastructure/dtos"
	mapper "github.com/PeteProgrammer/go-automapper"
	"log"
	"encoding/json"
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

	jsonMessage, err := json.Marshal(marketDataDto)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Saved message " + string(jsonMessage))
}