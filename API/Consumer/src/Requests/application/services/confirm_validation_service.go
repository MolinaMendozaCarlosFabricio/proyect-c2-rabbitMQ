package services

import (
	"fmt"

	"products_api.com/p/src/Requests/domain"
)

type ConfirmValidationRequestService struct {
	db domain.IService
}

func NewConfirmValidationRequestService(db domain.IService)*ConfirmValidationRequestService{
	return&ConfirmValidationRequestService{db: db}
}

func(service *ConfirmValidationRequestService)Execute(id_request int, id_status int)error{
	fmt.Println("Ejecutando servicio")
	return service.db.ConfirmValidationRequestMethod(id_request, id_status)
}