package main

type RabbitMQConfig struct {
	ConnectionString string
	QueueName        string
}

type MongoDBConfig struct {
	ConnectionString string
	Collection string
}

type Config struct {
	RabbitMQ RabbitMQConfig
	MongoDB MongoDBConfig
}
