package main

import (
	"fmt"
	"net/http"
	"strconv"

	socketio "github.com/googollee/go-socket.io"
)

type Request struct {
	ID int
	Date_request string
	Id_user int
	Status string
}

func main() {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("Nuevo usuario conectado: ", s.ID())
		s.SetContext("")
		return nil
	})

	server.OnEvent("/", "get_requests", func(s socketio.Conn, id_user int){
		fmt.Println("Obteniendo pedidos")
		// var input Request
		response, _ := http.Get("http://localhost:9080/requests/" + strconv.Itoa(id_user))
		s.Emit("requests_list", response)
	})

	server.OnEvent("/", "add_request", func (s socketio.Conn, id_request int)   {
		
	})
}