package application

import (
	"fmt"

	"products_api.com/p/src/Requests/domain"
)

type UpdateRequestStatusUC struct {
	db domain.IRequest
}

func NewUpdateRequestStatusUC(db domain.IRequest)*UpdateRequestStatusUC{
	return&UpdateRequestStatusUC{db: db}
}

func(uc *UpdateRequestStatusUC)Execute(id_request int)(bool, error){
	fmt.Println("Ejecutando caso de uso")
	id_status := 3
	flag, err := uc.db.ValidateRequestMethod(id_request)
	if err != nil {
		fmt.Println("Error al validar inventario")
	}
	if flag {
		id_status = 2
	} else {
		id_status = 1
		uc.db.ReduceStockMethod(id_request)
	}
	return (flag), (uc.db.UpdateRequestsStatusMethod(id_status, id_request))
}