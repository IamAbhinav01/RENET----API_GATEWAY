package services

import (
	"fmt"
	"log"
	"renet/DB/repositories"
	dtos "renet/DTOs"
	"renet/security/argon2id"
)

type UserService interface {
	SignUpUser(payload dtos.SignupRequestDTO) (int, error)
	SignInUser(payload dtos.SignInRequestDTO) error
}

type UserServiceImpl struct {
	repo repositories.UserRepository
}

func NewUserService(_repo repositories.UserRepository) UserService {
	return &UserServiceImpl{
		repo: _repo,
	}
}

func (user UserServiceImpl) SignUpUser(payload dtos.SignupRequestDTO) (int, error) {

	cfg := argon2id.NewArgonConfig()
	hashedPassword,err := cfg.HashPassword(payload.Password)
	log.Println("Hashed Password : ",hashedPassword)
	if err != nil{
		log.Println("Error occured while hashing the password")
		return 0, err
	}

	userID, err := user.repo.CreateUser(payload.Name,payload.Email,hashedPassword)
	if err != nil {
		log.Println("Error occure while sending data from service to repo", err)
		return 0, err
	} else {

		log.Println("Successfully sent data from service layer to repo layer")
	}

	return userID, nil
}

func(user UserServiceImpl) SignInUser(payload dtos.SignInRequestDTO) error{

	cfg := argon2id.NewArgonConfig()
	

	userDetails,err := user.repo.GetUserCredentialByEmail(payload.Email)
	if err != nil{
		log.Println("Error occured while accessing user details")
		return err
	}

	match,err := cfg.VerifyPassword(userDetails.Password,payload.Password)
	if err != nil{
		log.Println("Error occured while verifying the user password")
		return err
	}

	if(!match){
		log.Println("Invalid Password send the correct password")
		return fmt.Errorf("Invalid credential,check your password")
	}

	return nil

}