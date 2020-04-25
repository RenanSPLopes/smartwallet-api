package main

func main() {
	r := registerRoutes()
	r.Run(":19000")

	config := provideConfig()
	rabbitmqClient := ProvideRabbitMQClient()
	rabbitmqClient.Listen(config.RabbitMQ.QueueName)
}
