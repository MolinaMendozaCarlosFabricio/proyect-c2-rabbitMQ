package adapters

import (
	"context"
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
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

func(r *RequestRepoRabbitMQ)ConfirmValidationRequestMethod(id_request int, id_status int)error{
	conn, err := amqp.Dial("amqp://charly:666demon@13.217.73.3:5672/")
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

	var request_info struct {
		Id_request int
		Status string
	}

	request_info.Id_request = id_request

	if id_status == 1 {
		request_info.Status = "Aceptado"
	} else if id_status == 2 {
		request_info.Status = "Cancelado"
	} else if id_status == 3 {
		request_info.Status = "Pendiente"
	}
	
	err = ch.ExchangeDeclare(
		"inventory_analiser2",   // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")


	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
    defer cancel()


	jsonBody, err := json.Marshal(request_info)
	failOnError(err, "Error al serializar JSON")

	err = ch.PublishWithContext(ctx,
		"inventory_analiser2",           // exchange
		"queue_analiser",     // routing key
		false,        // mandatory
		false,
		amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         jsonBody,
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", jsonBody)

	return err
}