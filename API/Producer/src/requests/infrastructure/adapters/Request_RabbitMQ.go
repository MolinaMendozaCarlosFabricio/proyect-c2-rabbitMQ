package adapters

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"request_api.com/r/src/requests/domain"
)

type RequestRepoRabbitMQ struct {
}

func NewRequestRepoRabbitMQ()*RequestRepoRabbitMQ{
	return&RequestRepoRabbitMQ{}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func SendRequestMethod(acquire domain.Acquires){
	conn, err := amqp.Dial("amqp://charly:666demon@13.217.73.3:5672/")
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()
/*
	q, err := ch.QueueDeclare(
		"check_inventory2", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")
*/
	err = ch.ExchangeDeclare(
		"inventory_distributor",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")


	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
    defer cancel()

	jsonBody, err := json.Marshal(acquire)
	failOnError(err, "Error al serializar JSON")

	err = ch.PublishWithContext(ctx,
		"inventory_distributor",           // exchange
		"",       // routing key
		false,        // mandatory
		false,
		amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         jsonBody,
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", jsonBody)
}



func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
			s = "hello"
	} else {
			s = strings.Join(args[1:], " ")
	}
	return s
}

//func(r *RequestRepoRabbitMQ)AddProductToRequestMethod(id_request int, id_product int)