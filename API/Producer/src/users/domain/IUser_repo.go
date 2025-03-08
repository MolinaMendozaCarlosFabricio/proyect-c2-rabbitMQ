package domain

type IUser interface{
	SaveUserMethod(user User) error
	GetOneUserMethod(id int) ([]User, error)
	EditUserMethod(user User) error
	DeleteUserMethod(id int) error
}

