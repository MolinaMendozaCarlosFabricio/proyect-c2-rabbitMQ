package application

import (
	"fmt"

	"request_api.com/r/src/users/domain"
)

type EditUserUC struct {
	db domain.IUser
}

func NewEditUserUC(db domain.IUser)*EditUserUC{
	return&EditUserUC{db: db}
}

func(uc *EditUserUC)Execute(
	id int, 
	name string, 
	last_name string, 
	email string, 
	cellphone string, 
	password string,
	)error{
	user := &domain.User{
		ID: id, 
		Name: name, 
		Last_name: last_name, 
		Email: email, 
		Cellphone: cellphone, 
		Password: password,
	}
	fmt.Println("Ejecutando caso de uso")
	return uc.db.EditUserMethod(*user)
}