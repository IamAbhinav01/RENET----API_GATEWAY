package services

import (
	"log"
	"renet/DB/repositories"
	dtos "renet/DTOs"
)

type UserService interface {
	SignUpUser(payload dtos.SignupRequestDTO) error
}

type UserServiceImpl struct {
	repo repositories.UserRepository
}

func NewUserService(_repo repositories.UserRepository) UserService {
	return &UserServiceImpl{
		repo: _repo,
	}
}

func (user UserServiceImpl) SignUpUser(payload dtos.SignupRequestDTO) error {

	err := user.repo.CreateUser(payload.Name,payload.Email,payload.Password)
	if err != nil {
		log.Println("Error occure while sending data from service to repo", err)
		return err
	} else {

		log.Println("Successfully sent data from service layer to repo layer")
	}

	return nil
}
