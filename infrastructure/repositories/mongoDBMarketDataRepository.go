package repositories

import (
	"context"
	"fmt"
	"log"
	"smartwallet-api/domain/entities"
	"smartwallet-api/infrastructure/dtos"
	"time"

	mapper "github.com/PeteProgrammer/go-automapper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MarketDataRepository interface {
	Save(marketData entities.MarketData)
	GetAll() []dtos.MarketData
	GetById(id string) dtos.MarketData
	GetByCode(code string) dtos.MarketData
	UpdateResults(id primitive.ObjectID, result entities.Result)
	UpdateQuotes(code string, quote float32)
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
	marketDataDto.ID = primitive.NewObjectID()

	ctx, client, marketDataCollection := m.GetMarketDataCollection()
	defer client.Disconnect(ctx)
	result, err := marketDataCollection.InsertOne(ctx, marketDataDto)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	fmt.Println(result.InsertedID)
}

func (m MongoDBMarketDataRepository) GetAll() []dtos.MarketData {
	ctx, client, marketDataCollection := m.GetMarketDataCollection()
	defer client.Disconnect(ctx)

	projection := bson.D{
		{"name", 1},
		{"sector", 1},
		{"subsector", 1},
		{"segmentation", 1},
		{"b3segmentation", 1},
		{"marketvalue", 1},
	}

	cursor, err := marketDataCollection.Find(ctx, bson.M{}, options.Find().SetProjection(projection))
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

func (m MongoDBMarketDataRepository) GetById(id string) dtos.MarketData {
	ctx, client, marketDataCollection := m.GetMarketDataCollection()
	defer client.Disconnect(ctx)

	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Fatal(err.Error())
		panic(err.Error())
	}
	var result dtos.MarketData

	err = marketDataCollection.FindOne(ctx, bson.M{"_id": _id}).Decode(&result)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return result
}

func (m MongoDBMarketDataRepository) GetByCode(code string) dtos.MarketData {
	ctx, client, marketDataCollection := m.GetMarketDataCollection()
	defer client.Disconnect(ctx)

	var result dtos.MarketData
	err := marketDataCollection.FindOne(ctx, bson.D{{"stocks.code", code}}).Decode(&result)

	if err != nil {
		log.Println("No document with code found")
		return dtos.MarketData{}
	}

	return result
}

func (m MongoDBMarketDataRepository) UpdateResults(_id primitive.ObjectID, result entities.Result) {
	ctx, client, marketDataCollection := m.GetMarketDataCollection()
	defer client.Disconnect(ctx)

	filter := bson.M{"_id": bson.M{"$eq": _id}}
	change := bson.M{"$push": bson.M{"results": result}}

	_, err := marketDataCollection.UpdateOne(ctx, filter, change)

	if err != nil {
		log.Println("Error updating document. " + err.Error())
		panic(err.Error())
	}
}

func (m MongoDBMarketDataRepository) UpdateQuotes(code string, quote float32) {
	ctx, client, marketDataCollection := m.GetMarketDataCollection()
	defer client.Disconnect(ctx)

	filter := bson.M{"stocks.code": bson.M{"$eq": code}}
	change := bson.M{"$set": bson.M{"stocks.quotes": quote}}

	_, err := marketDataCollection.UpdateOne(ctx, filter, change)

	if err != nil {
		log.Println("Error updating quotes. " + err.Error())
		panic(err.Error())
	}
}

func (m MongoDBMarketDataRepository) GetMarketDataCollection() (context.Context, *mongo.Client, *mongo.Collection) {
	client, ctx := m.createClient()

	database := client.Database("SmartWallet")
	marketDataCollection := database.Collection("marketdata")

	return ctx, client, marketDataCollection
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
