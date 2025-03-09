package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Acquires struct{
	Id_request int
	Id_prodcut int
	Quantity int
}

func main(){
	conn, err := amqp.Dial("amqp://charly:666demon@13.217.73.3:5672/")
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

	ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

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


	q, err := ch.QueueDeclare(
		"", // name
		false,         // durable
		false,        // delete when unused
		true,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		"inventory_distributor", // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	err = ch.ExchangeDeclare(
		"inventory_analiser",   // name
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


	var forever chan struct{}

        go func() {
                for d := range msgs {
					var acq Acquires
					if err := json.Unmarshal(d.Body, &acq); err != nil {
						log.Printf("Error al deserializar el mensaje: %v", err)
						d.Nack(false, false)
						continue
					}

					log.Printf("Pedido recibido: %+v", acq)

					data := []byte(`{
						"id_request": ` + strconv.Itoa(acq.Id_request) + "," + `
						"id_product": ` + strconv.Itoa(acq.Id_prodcut) + "," + `
						"quantity": ` + strconv.Itoa(acq.Quantity) + 
					`}`,)

					request, err := http.NewRequest(
						"PUT",
						"http://localhost:9080/requests/product",
						bytes.NewBuffer(data),
					)

					if err != nil {
						log.Printf("Error al preparar la solicitud HTTP: %v", err)
						d.Nack(false, false)
						continue
					}

					request.Header.Set("Content-Type", "application/json")

					client := http.Client{}

					response, err := client.Do(request)
		
					if err != nil {
						log.Printf("Error al hacer la solicitud HTTP: %v", err)
						d.Nack(false, false)
						continue
					}

					log.Printf("Respuesta recibida: %+v", response)
                    d.Ack(false)

					jsonBody, err := json.Marshal("Análisis completo")
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
        }()

		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}