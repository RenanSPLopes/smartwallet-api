package main

func main() {
	config := provideConfig()
	go initializeMarketDataListener(config)

	r := registerRoutes()
	r.Run(":19001")
}

func initializeMarketDataListener(c Config) {
	rabbitmqClient := ProvideRabbitMQClient()
	rabbitmqClient.ListenMarketDataQueue(c.RabbitMQ.MarketDataQueueName)
}

func initializeStocksQuotationListener(c Config) {
	rabbitMQClient := ProvideRabbitMQClient()
	rabbitMQClient.ListenStocksQuotation(c.RabbitMQ.StocksQuotationQueueName)
}
