package main

func main() {
	config := provideConfig()
	go initializeRabbitMq(config)

	r := registerRoutes()
	r.Run(":19001")
}

func initializeRabbitMq(c Config) {
	rabbitmqClient := ProvideRabbitMQClient()
	rabbitmqClient.Listen(c.RabbitMQ.QueueName)
}
