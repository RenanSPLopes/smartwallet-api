package main

func main() {
	r := registerRoutes()

	initRabbitMQClient()

	r.Run(":19000")
}
