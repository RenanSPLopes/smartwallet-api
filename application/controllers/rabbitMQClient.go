package controllers

import (
	"log"
	"smartwallet/smartwallet-api/application/models"
	"smartwallet/smartwallet-api/application/services"

	"github.com/streadway/amqp"

	"encoding/json"
)

type RabbitMQClient struct {
	ConnectionString    string
	MarketDataProcessor services.MarketDataProcessor
}

func NewRabbitMQClient(connectrionString string, marketDataProcessor services.MarketDataProcessor) RabbitMQClient {
	return RabbitMQClient{ConnectionString: connectrionString, MarketDataProcessor: marketDataProcessor}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatal("%s: %s", msg, err)
	}
}

func (r RabbitMQClient) Listen(queueName string) {
	conn, err := amqp.Dial(r.ConnectionString)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			var m models.MarketData
			err := json.Unmarshal(d.Body, &m)

			if err != nil {
				log.Fatal("Failed to decode message.")
			}

			MarketDataProcessor.Process(m)
		}
	}()

	log.Printf(" [*] Waiting for messages.")

	<-forever
}
