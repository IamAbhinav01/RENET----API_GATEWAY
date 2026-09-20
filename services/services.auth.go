package services

import (
	"log"
	"renet/DB/repositories"
)

type UserService interface {
	SignUpUser(_name string, email string, password string) error
}

type UserServiceImpl struct {
	repo repositories.UserRepository
}

func NewUserService(_repo repositories.UserRepository) UserService {
	return &UserServiceImpl{
		repo: _repo,
	}
}

func (user UserServiceImpl) SignUpUser(_name string, email string, password string) error {

	err := user.repo.CreateUser(_name, email, password)
	if err != nil {
		log.Println("Error occure while sending data from service to repo", err)
		return err
	} else {

		log.Println("Successfully sent data from service layer to repo layer")
	}

	return nil
}
