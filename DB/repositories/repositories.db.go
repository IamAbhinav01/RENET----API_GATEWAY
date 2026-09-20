package repositories

type UserRepository interface{
	CreateUser(_name string,email string,password string)
}

func New