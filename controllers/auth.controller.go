package controllers

import (
	"net/http"
	dtos "renet/DTOs"
	"renet/middlewares"
	"renet/services"
	"renet/utils/formatters"
	"strings"
)

type AuthController struct {
	AuthService services.UserService
}


func NewAuthController(serv services.UserService) *AuthController {
	return &AuthController{
		AuthService: serv,
	}
}

func (cntrl AuthController) SignUp(w http.ResponseWriter, r *http.Request) {
	payload, ok := r.Context().Value(middlewares.PayloadContextKet).(dtos.SignupRequestDTO)
	if !ok {
		formatters.ErrorResponse(w, http.StatusBadRequest, nil, "Invalid signup request")
		return
	}

	err := cntrl.AuthService.SignUpUser(payload)

	if err!= nil{
		status := http.StatusInternalServerError
		if strings.Contains(strings.ToLower(err.Error()),"duplicate") || strings.Contains(err.Error(), "1062"){
			status = http.StatusConflict
		}
		formatters.ErrorResponse(w,status,err,"Error occured while signing the user")
		return
	}
	formatters.SuccessResponse(w,http.StatusCreated,"User sign-up successfully")
}
func Login()  {}
func LogOut() {}
