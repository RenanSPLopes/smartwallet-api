//+build wireinject

package main

import (
	"os"
	"smartwallet-api/application/controllers"
	"smartwallet-api/application/services"
	"smartwallet-api/infrastructure/repositories"
	"github.com/google/wire"
)

func provideConfig() Config {
	return Config{
		RabbitMQ: RabbitMQConfig{
			ConnectionString: os.Getenv("RABBIT_CONNECTIONSTRING"),
			QueueName:        os.Getenv("MARKETDATA_QUEUE_NAME"),
		},
		MongoDB: MongoDBConfig {
			ConnectionString: "",
			Collection: "",
		},
	}
}

func provideMongoDBMarketDataRepository(c Config) repositories.MongoDBMarketDataRepository{
	return repositories.NewMongoDBMarketDataRepository(c.MongoDB.ConnectionString, c.MongoDB.Collection)
}

func provideRabbitMQClient(c Config, m services.MarketDataProcessorService) controllers.RabbitMQClient {
	return controllers.NewRabbitMQClient(c.RabbitMQ.ConnectionString, m)
}

func provideMarketDataProcessor(m repositories.MongoDBMarketDataRepository) services.MarketDataProcessorService {
	return services.NewMarketDataProcessorService(m)
}

func ProvideRabbitMQClient() controllers.RabbitMQClient {
	panic(wire.Build(provideRabbitMQClient, provideConfig, provideMarketDataProcessor, provideMongoDBMarketDataRepository))
}
