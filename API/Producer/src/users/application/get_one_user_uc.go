package application

import (
	"fmt"

	"request_api.com/r/src/users/domain"
)

type GetOneUserUC struct {
	db domain.IUser
}

func NewGetOneUserUC(db domain.IUser)*GetOneUserUC{
	return&GetOneUserUC{db: db}
}

func(uc *GetOneUserUC)Execute(id int)([]domain.User, error){
	fmt.Println("Ejecutando caso de uso")
	return uc.db.GetOneUserMethod(id)
}