package main

func main() {
	r := registerRoutes()

	config := provideConfig()
	rabbitmqClient := ProvideRabbitMQClient()
	rabbitmqClient.Listen(config.RabbitMQ.QueueName)

	r.Run(":19001")
}
