package middlewares

import (
	"net/http"
	dtos "renet/DTOs"
	"renet/utils/helpers"
)

func SignUpRequestValidation(next http.Handler) http.Handler {
	return helpers.RequestValidation[dtos.SignupRequestDTO](next)
}

func SignInRequestValidation(next http.Handler) http.Handler {
	return helpers.RequestValidation[dtos.SignInRequestDTO](next)
}