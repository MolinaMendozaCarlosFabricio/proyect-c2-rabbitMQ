package services

import (
	"fmt"

	"request_api.com/r/src/requests/domain"
)

type SendRequestToVerifyMethodService struct {
	db domain.IService
}

func NewSendRequestToVerifyMethodService(db domain.IService)*SendRequestToVerifyMethodService{
	return&SendRequestToVerifyMethodService{db: db}
}

func(s *SendRequestToVerifyMethodService)Execute(id_request int, id_product int, quiantity int)error{
	fmt.Println("Ejecutando servicio")
	return s.db.SendRequestToVerifyMethod(id_request, id_product, quiantity)
}