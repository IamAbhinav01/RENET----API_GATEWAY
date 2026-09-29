package repositories

import (
	"log"
	"renet/schemas"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(_name string, email string, password string) error
	GetUserCredentialByEmail(email string) (*schemas.SignUp, error)
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

	user := schemas.SignUp{
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

func(repo *UserRespositoryImpl) GetUserCredentialByEmail(email string) (*schemas.SignUp, error) {

	user := schemas.SignUp{
		Email: email,
	}

	result := repo.db.Select("id","email","password").Where("email=?",email).First(&user)
	if result.Error != nil{
		return nil,result.Error
	}

	log.Println("Successfully fetched the credential using email")
	return &user,nil

}