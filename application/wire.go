//+build wireinject

package main

import (
	"os"
	"smartwallet-api/application/controllers"
	"smartwallet-api/application/services"

	"github.com/google/wire"
)

func provideConfig() Config {
	return Config{
		RabbitMQ: RabbitMQConfig{
			ConnectionString: os.Getenv("RABBIT_CONNECTIONSTRING"),
			QueueName:        os.Getenv("MARKETDATA_QUEUE_NAME"),
		},
	}
}

func provideRabbitMQClient(c Config, m services.MarketDataProcessorService) controllers.RabbitMQClient {
	return controllers.NewRabbitMQClient(c.RabbitMQ.ConnectionString, m)
}

func provideMarketDataProcessor() services.MarketDataProcessorService {
	return services.NewMarketDataProcessorService()
}

func ProvideRabbitMQClient() controllers.RabbitMQClient {
	panic(wire.Build(provideRabbitMQClient, provideConfig, provideMarketDataProcessor))
}
