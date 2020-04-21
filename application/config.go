package main

type RabbitMQConfig struct {
	ConnectionString string
	QueueName        string
}

type Config struct {
	RabbitMQ RabbitMQConfig
}
