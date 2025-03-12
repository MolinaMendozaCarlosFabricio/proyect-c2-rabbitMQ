package infrastructure

import (
	"log"

	"products_api.com/p/src/Requests/domain"
	acquires_domain "products_api.com/p/src/Products/domain"
	"products_api.com/p/src/core"
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

func(r *RequestRepoMySQL)ValidateRequestMethod(id int)(bool, error){
	query := "SELECT products.stock FROM products INNER JOIN acquires ON acquires.id_product = products.id  WHERE acquires.id_request = ?"
	rows, err := r.Connection.FetchRows(query, id)
	flag := false
	if err != nil {
        log.Fatalf("Error al obtener productos:", err)
    }
    defer rows.Close()
	for rows.Next(){
		var stock int

		if err := rows.Scan(&stock); err != nil{
			log.Println("Error al escanear la fila:", err)
		}

		if stock <= 0 {
			flag = true
		}
	}
	return flag, err
}

func(r *RequestRepoMySQL)UpdateRequestsStatusMethod(id_status int, id_request int)error{
	query2 := "UPDATE requests SET id_status = ? WHERE id = ?"
	_, err := r.Connection.ExecPreparedQuerys(query2, id_status, id_request)
	if err != nil {
        log.Fatalf("Error al editar producto: ", err)
    }
    return err
}

func(r *RequestRepoMySQL)GetAllMyRequestsMethod(id int)([]domain.Request, error){
	query := "SELECT requests.id, requests.date_request, status_request.name_status FROM requests INNER JOIN status_request ON requests.id_status = status_request.id WHERE requests.id_user = ?"
	rows, err := r.Connection.FetchRows(query, id)
	var requests []domain.Request
	if err != nil {
        log.Fatalf("Error al obtener pedidos:", err)
    }
    defer rows.Close()
	for rows.Next(){
		var id int
		var request_date string
		var status_request string
		
		if err := rows.Scan(&id, &request_date, &status_request); err != nil{
			log.Println("Error al escanear la fila:", err)
		}
		request := domain.Request{ID: id, Date_request: request_date, Status: status_request}
		requests = append(requests, request)
	}
	return requests, err
}

func(r *RequestRepoMySQL)ReduceStockMethod(id int)error{
	query := "SELECT products.id products.stock, acquires.quantity FROM products INNER JOIN acquires ON acquires.id_product = products.id WHERE acquires.id_request = ?"
	rows, err := r.Connection.FetchRows(query, id)
	if err != nil {
        log.Fatalf("Error al obtener productos:", err)
    }
    defer rows.Close()
	for rows.Next(){
		var id_product int
		var stock int
		var quantity int

		if err := rows.Scan(&id_product, &stock, &quantity); err != nil{
			log.Println("Error al escanear la fila:", err)
		}
		stock = stock - quantity
		query := "UPDATE products SET stock = ? WHERE id = ?"
		_, err := r.Connection.ExecPreparedQuerys(query, stock, id_product)
		if err != nil {
			log.Fatalf("Error al eliminar producto: ", err)
		}
	}
	return err
}

func(r *RequestRepoMySQL)ReduceStockOfAProductMethod(acquire acquires_domain.Acquires)(int, error){
	int_status := 0

	query := "SELECT id_status FROM requests WHERE id = ?"
	rows, err := r.Connection.FetchRows(query, acquire.Id_request)
	if err != nil {
        log.Fatalf("Error al obtener productos:", err)
    }
    defer rows.Close()
	for rows.Next(){
		var id_status int

		if err := rows.Scan(&id_status); err != nil{
			log.Println("Error al escanear la fila:", err)
		}
		if id_status != 2 {
			query := "SELECT stock FROM products WHERE id = ?"
			rows, err := r.Connection.FetchRows(query, acquire.Id_product)
			if err != nil {
        		log.Fatalf("Error al obtener productos:", err)
		    }
		    defer rows.Close()
			for rows.Next(){
				var stock int

				if err := rows.Scan(&stock); err != nil{
					log.Println("Error al escanear la fila:", err)
				}

				status := 3

				if 0 <= stock - acquire.Quantity {
					stock = stock - acquire.Quantity
					status = 1
					query := "UPDATE products SET stock = ? WHERE id = ?"
					_, err := r.Connection.ExecPreparedQuerys(query, stock, acquire.Id_product)
					if err != nil {
						log.Fatalf("Error al actualizar existencias: ", err)
					}
				} else {
					status = 2
				}

				id_status = status
				query := "UPDATE requests SET id_status = ? WHERE id = ?"
				_, err := r.Connection.ExecPreparedQuerys(query, status, acquire.Id_request)
				if err != nil {
					log.Fatalf("Error al eliminar producto: ", err)
				}
			}
		}
	}
	
	return int_status, err
}

func(r *RequestRepoMySQL)GetOneOfMyRequestsMethod(id_request int)([]domain.Request, error){
	query := "SELECT requests.id, requests.date_request, status_request.name_status FROM requests INNER JOIN status_request ON requests.id_status = status_request.id WHERE requests.id = ?"
	rows, err := r.Connection.FetchRows(query, id_request)
	var requests []domain.Request
	if err != nil {
        log.Fatalf("Error al obtener pedidos:", err)
    }
    defer rows.Close()
	for rows.Next(){
		var id int
		var request_date string
		var status_request string
		
		if err := rows.Scan(&id, &request_date, &status_request); err != nil{
			log.Println("Error al escanear la fila:", err)
		}
		request := domain.Request{ID: id, Date_request: request_date, Status: status_request}
		requests = append(requests, request)
	}
	return requests, err
}