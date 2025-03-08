package infrastructure

import (
	"log"
	//"time"

	"request_api.com/r/src/core"
	"request_api.com/r/src/requests/domain"
	"request_api.com/r/src/requests/infrastructure/adapters"
)

type RequestRepoMySQL struct {
	Connection core.ConectionMySQL
}

func NewRequestRepoMySQL()*RequestRepoMySQL{
	conn := core.MySQLConection()
    if conn.Err != "" {
        log.Fatal("Error al configurar la pool de conexiones", conn.Err)
    }
	return &RequestRepoMySQL{Connection: *conn}
}

func(r *RequestRepoMySQL)CreateRequestMethod(request domain.Request)(int, error){
	query := "INSERT INTO requests (date_request, id_user, id_status) VALUES (?, ?, ?)"
	result, err := r.Connection.ExecPreparedQuerys(query, request.Date_request, request.Id_user, request.Id_status)
	if err != nil {
        log.Fatalf("Error al registrar Pedido:", err)
    }
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("Error al obtener último id del pedido:", err)
	}
	request.ID = int(id)
	//time.Sleep(2*time.Minute)
	//adapters.SendRequestMethod(request)
    return request.ID,err
}

func(r *RequestRepoMySQL)AddProductToRequestMethod(id_request int, id_product int, quiantity int)error{
	query := "INSERT INTO acquires (id_request, id_product, quantity) VALUES (?, ?, ?)"
	_, err := r.Connection.ExecPreparedQuerys(query, id_request, id_product, quiantity)
	if err != nil {
        log.Fatalf("Error al enlazar prodcutos con pedidos:", err)
    }
	acquire := &domain.Acquires{Id_request: id_request, Id_product: id_product, Quantity: quiantity}
	adapters.SendRequestMethod(*acquire)
    return err
}