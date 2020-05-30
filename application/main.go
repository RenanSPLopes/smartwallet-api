package main

func main() {
	config := provideConfig()
	go initializeMarketDataListener(config)

	r := registerRoutes()
	r.Run(":19001")
}

func initializeMarketDataListener(c Config) {
	// marketDataProcessor := ProvideMarketDataProcessor()
	rabbitmqClient := ProvideRabbitMQClient()
	rabbitmqClient.ListenMarketDataQueue(c.RabbitMQ.QueueName)
}

func initializeStocksQuotationListener() {

}
