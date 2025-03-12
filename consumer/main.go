package main

import (
	"bytes"
	_"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	_"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Acquires struct{
	Id_request int
	Id_product int
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
		"inventory_distributor2",   // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")


	q, err := ch.QueueDeclare(
		"queue_of_distribution", // name
		true,         // durable
		false,        // delete when unused
		true,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name, // queue name
		"queue_distributor",     // routing key
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

	var forever chan struct{}

        go func() {
                for d := range msgs {
					log.Printf("Mensaje recibido: %s", d.Body)
					var acq Acquires
					if err := json.Unmarshal(d.Body, &acq); err != nil {
						log.Printf("Error al deserializar el mensaje: %v", err)
						d.Nack(false, false)
						continue
					}

					log.Printf("Pedido recibido: %+v", acq)

					data := []byte(`{
						"Id_request": ` + strconv.Itoa(acq.Id_request) + "," + `
						"Id_product": ` + strconv.Itoa(acq.Id_product) + "," + `
						"Quantity": ` + strconv.Itoa(acq.Quantity) + 
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