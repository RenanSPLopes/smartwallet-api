package repositories

import (
	"context"
	"fmt"
	mapper "github.com/PeteProgrammer/go-automapper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"smartwallet-api/domain/entities"
	"smartwallet-api/infrastructure/dtos"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MarketDataRepository interface {
	Save(marketData entities.MarketData)
	GetAll() []dtos.MarketData
}

type MongoDBMarketDataRepository struct {
	ConnectionString string
}

func NewMongoDBMarketDataRepository(conectionString string) *MongoDBMarketDataRepository {
	return &MongoDBMarketDataRepository{ConnectionString: conectionString}
}

func (m MongoDBMarketDataRepository) Save(marketData entities.MarketData) {
	var marketDataDto dtos.MarketData
	mapper.MapLoose(marketData, &marketDataDto)
	marketDataDto.ID = 	primitive.NewObjectID()
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

func (m MongoDBMarketDataRepository) GetAll() []dtos.MarketData {
	client, ctx := m.createClient()
	defer client.Disconnect(ctx)

	database := client.Database("SmartWallet")
	marketDataCollection := database.Collection("marketdata")

	cursor, err := marketDataCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	var marketData []dtos.MarketData
	if err = cursor.All(ctx, &marketData); err != nil {
		log.Fatal(err)
		panic(err)
	}

	return marketData
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
