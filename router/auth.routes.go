package router

import (
	"renet/controllers"
	"renet/middlewares"

	"github.com/go-chi/chi/v5"
)

func RegisterAuthRouter(r chi.Router, authController *controllers.AuthController) {
	r.With(middlewares.SignUpRequestValidation).Post("/signup", authController.SignUp)

}
