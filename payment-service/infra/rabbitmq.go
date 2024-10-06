package infra

import (
	"encoding/json"
	"log"
	"os"
	"payment-service/models"

	"github.com/streadway/amqp"
)

var channel *amqp.Channel
var queue *amqp.Queue

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func InitConfig() {

	host := os.Getenv("RABBIT_MQ_HOST")

	if host == "" {
		host = "localhost"
	}

	conn, err := amqp.Dial("amqp://guest:guest@" + host + ":5672/")

	log.Printf(host)

	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		"payment-queue", // name
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	failOnError(err, "Failed to declare a queue")

	queue = &q
	channel = ch
}

func SendMessage(payment models.Payment) {
	paymentJson, err := json.Marshal(payment)

	failOnError(err, "Failed to marshal payment")

	errCh := channel.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        paymentJson,
		})

	if errCh != nil {
		return
	}

	failOnError(err, "Failed to publish a message")
	log.Printf("sending message: %s", paymentJson)
}
