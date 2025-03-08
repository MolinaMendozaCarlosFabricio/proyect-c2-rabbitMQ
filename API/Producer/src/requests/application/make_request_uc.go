package application

import (
	"fmt"
	"time"

	"request_api.com/r/src/requests/domain"
)

type MakeRequestUC struct {
	db domain.IRequest
}

func NewMakeRequestUC(db domain.IRequest)*MakeRequestUC{
	return&MakeRequestUC{db: db}
}

func(uc *MakeRequestUC)Execute(id_user int)(int, error){
	request := &domain.Request{ID: 0, Date_request: time.Now(), Id_user: id_user, Id_status: 3}
	fmt.Println("Ejecutando caso de uso")
	return uc.db.CreateRequestMethod(*request)
}