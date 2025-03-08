package infrastructure

import (
	"log"

	"request_api.com/r/src/core"
	"request_api.com/r/src/users/domain"
)

type UserRepoMySQL struct {
	Connection core.ConectionMySQL
}

func NewUserRepoMySQL() *UserRepoMySQL{
	conn := core.MySQLConection()
    if conn.Err != "" {
        log.Fatal("Error al configurar la pool de conexiones", conn.Err)
    }
	return &UserRepoMySQL{Connection: *conn}
}

func (r *UserRepoMySQL) SaveUserMethod(user domain.User)error{
	query := "INSERT INTO users (name, last_name, email, cellphone, password) VALUES (?, ?, ?, ?, ?)"
	_, err := r.Connection.ExecPreparedQuerys(query, user.Name, user.Last_name, user.Email, user.Cellphone, user.Password)
	if err != nil {
        log.Fatalf("Error al registrar Usuarios:", err)
    }
    return err
}

func (r *UserRepoMySQL) GetOneUserMethod(id int)([]domain.User, error){
	query := "SELECT id, name, last_name, email, cellphone FROM users WHERE id = ?"
	rows, err := r.Connection.FetchRows(query, id)
	var users []domain.User
	if err != nil {
        log.Fatalf("Error al obtener Usuarios:", err)
    }
    defer rows.Close()
	for rows.Next(){
		var id int
		var name string
		var last_name string
		var email string
		var cellphone string

		if err := rows.Scan(&id, &name, &last_name, &email, &cellphone); err != nil{
			log.Println("Error al escanear la fila:", err)
		}
		user := domain.User{ID: id, Name: name, Last_name: last_name, Email: email, Cellphone: cellphone, Password: ""}
		users = append(users, user)
	}
	return users, err
}

func (r *UserRepoMySQL) EditUserMethod(user domain.User)error{
	query := "UPDATE users SET name = ?, last_name = ?, email = ?, cellphone = ?, password = ? WHERE id = ?"
	_, err := r.Connection.ExecPreparedQuerys(query, user.Name, user.Last_name, user.Email, user.Cellphone, user.Password, user.ID)
	if err != nil {
        log.Fatalf("Error al editar info. dl usuario:", err)
    }
    return err
}

func (r *UserRepoMySQL) DeleteUserMethod(id int)error{
	_, err := r.Connection.ExecPreparedQuerys("DELETE FROM users WHERE id = ?", id)
	if err != nil {
        log.Fatalf("Error al eliminar usuario:", err)
    }
	return err
}