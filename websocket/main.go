package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"github.com/rs/cors"
)

type Request struct {
	ID int
	Date_request string
	Id_user int
	Status string
}

func main() {
	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			polling.Default,
			websocket.Default,
		},
	})

	server.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("Nuevo usuario conectado: ", s.ID())
		s.SetContext("")
		return nil
	})

	server.OnEvent("/", "get_requests", func(s socketio.Conn, id_user int){
		fmt.Println("Obteniendo pedidos")
		// var input Request
		response, err := http.Get("http://localhost:9080/requests/" + strconv.Itoa(id_user))
		if err != nil {
			fmt.Println("Error al obtener los pedidos:", err)
			s.Emit("error", "Error al obtener los pedidos")
			return
		}
		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error al leer la respuesta:", err)
			s.Emit("error", "Error al procesar la respuesta")
			return
		}
		fmt.Println("Respuesta: ", string(body))
		s.Emit("requests_list", string(body))
	})

	server.OnEvent("/", "add_request", func(s socketio.Conn, id_request int)   {
		fmt.Println("Registrando pedido")
		response, err := http.Get("http://localhost:9080/requests/one/" + strconv.Itoa(id_request))
		if err != nil {
			fmt.Println("Error al obtener pedido:", err)
			s.Emit("error", "Error al obtener pedido")
			return
		}
		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error al leer la respuesta:", err)
			s.Emit("error", "Error al procesar la respuesta")
			return
		}
		fmt.Println("Respuesta: ", string(body))
		s.Emit("one_request", string(body))
	})

	server.OnDisconnect("/", func(c socketio.Conn, s string) {
		fmt.Println("Usuario desconectado: ", c.ID())
	})

	http.Handle("/socket.io/", server)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
		Debug:            true, // Para ver detalles en consola
	})

	handler := c.Handler(server)

	fmt.Println("Corriendo servidor")
	err := http.ListenAndServe(":7080", handler)
	if err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}