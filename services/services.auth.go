package services

import (
	"log"
	"renet/DB/repositories"
	dtos "renet/DTOs"
	"renet/security/argon2id"
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

	cfg := argon2id.NewArgonConfig()
	hashedPassword,err := cfg.HashPassword(payload.Password)
	log.Println("Hashed Password : ",hashedPassword)
	if err != nil{
		log.Println("Error occured while hashing the password")
		return err
	}

	err = user.repo.CreateUser(payload.Name,payload.Email,hashedPassword)
	if err != nil {
		log.Println("Error occure while sending data from service to repo", err)
		return err
	} else {

		log.Println("Successfully sent data from service layer to repo layer")
	}

	return nil
}
