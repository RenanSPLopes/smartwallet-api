package repositories

import (
	"context"
	"fmt"
	"log"
	"smartwallet-api/domain/entities"
	"smartwallet-api/infrastructure/dtos"
	"time"

	mapper "github.com/PeteProgrammer/go-automapper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MarketDataRepository interface {
	Save(marketData entities.MarketData)
}

type MongoDBMarketDataRepository struct {
	ConnectionString string
}

func NewMongoDBMarketDataRepository(conectionString string) MongoDBMarketDataRepository {
	return MongoDBMarketDataRepository{ConnectionString: conectionString}
}

func (m MongoDBMarketDataRepository) Save(marketData entities.MarketData) {
	var marketDataDto dtos.MarketData
	mapper.MapLoose(marketData, &marketDataDto)

	client, ctx := m.createClient()
	defer client.Disconnect(ctx)

	database := client.Database("SmartWallet")
	marketdataCollection := database.Collection("marketdata")
	result, err := marketdataCollection.InsertOne(ctx, marketDataDto)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	fmt.Println(result.InsertedID)
}

func (m MongoDBMarketDataRepository) createClient() (*mongo.Client, context.Context) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.ConnectionString))

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return client, ctx
}
