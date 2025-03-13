package main

import (
	"encoding/json"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Request struct {
	ID int
	Date_request string
	Id_user int
	Status string
}

type Responses struct{
	Id_request int
	Status string
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func consumeRabbitMQ(server *socketio.Server){
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

	for d := range msgs {
		log.Printf("Mensaje recibido: %s", d.Body)
		var rsp Responses
		if err := json.Unmarshal(d.Body, &rsp); err != nil {
			log.Printf("Error al deserializar el mensaje: %v", err)
			d.Nack(false, false)
			continue
		}

		log.Printf("Pedido recibido: %+v", rsp)
		server.BroadcastToNamespace("/", "rabbitmq_message", rsp)
	}
}

func main() {
	// Crear servidor Socket.IO
	server := socketio.NewServer(nil)
	// Manejar eventos de conexión
	server.OnConnect("/", func(s socketio.Conn) error {
		log.Println("Nuevo cliente conectado:", s.ID())
		return nil
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("Error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("Cliente desconectado:", s.ID(), reason)
	})

	// Iniciar el consumidor de RabbitMQ en un goroutine
	go consumeRabbitMQ(server)

	// Iniciar servidor HTTP
	// Middleware para permitir CORS
	mux := http.NewServeMux()

	mux.Handle("/socket.io/", server)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // Manejo de CORS
    	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	    w.Header().Set("Access-Control-Allow-Credentials", "true")
    	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, X-Requested-With")
	    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

    // Responde al pre-flight request
	    if r.Method == "OPTIONS" {
    	    return
    	}

	    mux.ServeHTTP(w, r) // Llama al siguiente manejador
	})

	log.Println("Servidor escuchando en :7080")
	log.Fatal(http.ListenAndServe(":7080", handler))
}