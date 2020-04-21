package main

import (
	"os"
	"smartwallet-api/application/controllers"

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

func provideRabbitMQClient(c Config) *controllers.RabbitMQClient {
	return controllers.NewRabbitMQClient(c.RabbitMQ.ConnectionString)
}

func ProvideRabbitMQClient() controllers.RabbitMQClient {
	panic(wire.Build(provideRabbitMQClient, provideConfig))
}
