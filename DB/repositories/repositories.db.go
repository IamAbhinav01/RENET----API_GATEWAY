package repositories

import (
	"renet/schemas"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(_name string, email string, password string) error
}

type UserRespositoryImpl struct {
	db *gorm.DB
}

func NewUserRespository(_db *gorm.DB)UserRepository{
	return &UserRespositoryImpl{
		db: _db,
	}
}

func(repo *UserRespositoryImpl)CreateUser(_name string, email string, password string) error{

	user := schemas.User{
		Name: _name,
		Email: email,
		Password: password,
	}

	result := repo.db.Create(&user)
	
	if result.Error != nil{
		return result.Error
	}

	return nil
}