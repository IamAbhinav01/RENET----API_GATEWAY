package services

import "renet/DB/repositories"

type UserService interface {
}

type UserServiceImpl struct {
	repo repositories.UserRepository
}

func NewUserService(_repo repositories.UserRepository) UserService {
	return &UserServiceImpl{
		repo: _repo,
	}
}

func (repo UserServiceImpl) SignUpUser(){

}