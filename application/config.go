package main

type RabbitMQConfig struct {
	ConnectionString         string
	MarketDataQueueName      string
	StocksQuotationQueueName string
}

type MongoDBConfig struct {
	ConnectionString string
}

type Config struct {
	RabbitMQ RabbitMQConfig
	MongoDB  MongoDBConfig
}
