// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"os"
	"smartwallet-api/application/controllers"
	"smartwallet-api/application/services"
)

// Injectors from wire.go:

func ProvideRabbitMQClient() controllers.RabbitMQClient {
	config := provideConfig()
	marketDataProcessorService := provideMarketDataProcessor()
	rabbitMQClient := provideRabbitMQClient(config, marketDataProcessorService)
	return rabbitMQClient
}

// wire.go:

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
