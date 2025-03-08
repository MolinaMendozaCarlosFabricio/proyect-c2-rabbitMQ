package application

import (
	"fmt"

	"products_api.com/p/src/Requests/domain"
)

type GetAllMyRequestsMethodUC struct {
	db domain.IRequest
}

func NewGetAllMyRequestsMethodUC(db domain.IRequest)*GetAllMyRequestsMethodUC{
	return&GetAllMyRequestsMethodUC{db: db}
}

func(uc *GetAllMyRequestsMethodUC)Execute(id int)([]domain.Request, error){
	fmt.Println("Ejecutando caso de uso")
	return uc.db.GetAllMyRequestsMethod(id)
}