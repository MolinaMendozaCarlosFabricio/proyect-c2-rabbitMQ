package application

import (
	"fmt"

	"request_api.com/r/src/requests/domain"
)

type AddProductUC struct {
	db domain.IRequest
}

func NewAddProductUC(db domain.IRequest)*AddProductUC{
	return&AddProductUC{db: db}
}

func(uc *AddProductUC)Execute(id_request int, id_product int, quantity int)error{
	fmt.Println("Ejecutando caso de uso")
	return uc.db.AddProductToRequestMethod(id_request, id_product, quantity)
}