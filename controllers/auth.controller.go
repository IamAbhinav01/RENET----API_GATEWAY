package controllers

import (
	"errors"
	"net/http"
	"renet/services"
	"renet/utils/converter"
	"renet/utils/formatters"
)

type AuthController struct {
	AuthService services.UserService
}

type signupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuthController(serv services.UserService) *AuthController {
	return &AuthController{
		AuthService: serv,
	}
}

func (cntrl AuthController) SignUp(w http.ResponseWriter, r *http.Request) {
	var request signupRequest
	if err := converter.DecodeFromJSON(r, &request); err != nil {
		_ = formatters.ErrorResponse(w, http.StatusBadRequest, err, nil)
		return
	}
	if request.Name == "" || request.Email == "" || request.Password == "" {
		_ = formatters.ErrorResponse(w, http.StatusBadRequest, errors.New("name, email, and password are required"), nil)
		return
	}

	if err := cntrl.AuthService.SignUpUser(request.Name, request.Email, request.Password); err != nil {
		_ = formatters.ErrorResponse(w, http.StatusInternalServerError, err, nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = formatters.SuccessResponse(w, http.StatusCreated, map[string]string{"message": "user created"})
}
func Login()  {}
func LogOut() {}
