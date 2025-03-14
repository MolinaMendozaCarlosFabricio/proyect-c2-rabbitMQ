package main

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Response_info struct{
	Id_request int
	Status string
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main(){
	conn, err := amqp.Dial("amqp://charly:666demon@13.217.73.3:5672/")
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

	ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

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


	q, err := ch.QueueDeclare(
		"queue_of_analisys", // name
		true,         // durable
		false,        // delete when unused
		true,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name, // queue name
		"queue_analiser",     // routing key
		"inventory_analiser2", // exchange
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
					var acq Response_info
					if err := json.Unmarshal(d.Body, &acq); err != nil {
						log.Printf("Error al deserializar el mensaje: %v", err)
						d.Nack(false, false)
						continue
					}

					log.Printf("Pedido recibido: %+v", acq)
                    d.Ack(false)
                }
        }()

		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		<-forever
}