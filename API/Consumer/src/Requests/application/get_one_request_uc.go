package application

import (
	"fmt"

	"products_api.com/p/src/Requests/domain"
)

type GetOneOfMyRequestsMethodUC struct {
	db domain.IRequest
}

func NewGetOneOfMyRequestsMethodUC(db domain.IRequest)*GetOneOfMyRequestsMethodUC{
	return&GetOneOfMyRequestsMethodUC{db: db}
}

func(uc *GetOneOfMyRequestsMethodUC)Execute(id_request int)([]domain.Request, error){
	fmt.Println("Ejecutando caso de uso")
	return uc.db.GetOneOfMyRequestsMethod(id_request)
}