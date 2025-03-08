package application

import (
	"fmt"

	"request_api.com/r/src/users/domain"
)

type RegisterUserUC struct {
	db domain.IUser
}

func NewRegisterUserUC(db domain.IUser)*RegisterUserUC{
	return&RegisterUserUC{db: db}
}

func(uc *RegisterUserUC)Execute(
	name string, 
	last_name string, 
	email string, 
	cellphone string, 
	password string,
	)error{
	user := &domain.User{
		ID: 0, 
		Name: name, 
		Last_name: last_name, 
		Email: email, 
		Cellphone: cellphone, 
		Password: password,
	}
	fmt.Println("Ejecutando Caso de Uso")
	return uc.db.SaveUserMethod(*user)
}