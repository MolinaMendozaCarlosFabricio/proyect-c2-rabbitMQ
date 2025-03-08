package application

import (
	"fmt"

	"request_api.com/r/src/users/domain"
)

type DeleteUserUC struct {
	db domain.IUser
}

func NewDeleteUserUC(db domain.IUser)*DeleteUserUC{
	return&DeleteUserUC{db: db}
}

func(uc *DeleteUserUC)Execute(id int)error{
	fmt.Println("Ejecutando caso de uso")
	return uc.db.DeleteUserMethod(id)
}