package controllers

import (
	"log"
	"smartwallet-api/application/models"
	"smartwallet-api/application/services"

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

func (r RabbitMQClient) ListenMarketDataQueue(queueName string) {
	conn, ch, msgs := r.connectRabbitMQ(queueName)
	defer conn.Close()
	defer ch.Close()

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			// log.Printf("Received a message: %s", d.Body)
			var m models.MarketData
			err := json.Unmarshal(d.Body, &m)

			if err != nil {
				log.Fatal("Failed to decode message.")
			}

			r.MarketDataProcessor.Process(m)
		}
	}()

	log.Printf(" [*] Waiting for messages.")

	<-forever
}

func (r RabbitMQClient) ListenStocksQuotation(queueName string) {
	conn, ch, msgs := r.connectRabbitMQ(queueName)
	defer conn.Close()
	defer ch.Close()

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Stocks Quotation: %s", d.Body)
			var m models.StocksQuote
			err := json.Unmarshal(d.Body, &m)

			if err != nil {
				log.Fatal("Failed to decode stocks quotes message.")
				panic(err.Error())
			}
		}
	}()

	log.Printf(" [*] Waiting for messages.")

	<-forever
}

func (r RabbitMQClient) connectRabbitMQ(queueName string) (*amqp.Connection, *amqp.Channel, <-chan amqp.Delivery) {
	conn, err := amqp.Dial(r.ConnectionString)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

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

	return conn, ch, msgs
}
