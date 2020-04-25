package main

func main() {
	r := registerRoutes()
	r.Run(":19001")

	config := provideConfig()
	rabbitmqClient := ProvideRabbitMQClient()
	rabbitmqClient.Listen(config.RabbitMQ.QueueName)
}
