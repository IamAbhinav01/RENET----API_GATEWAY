package router

import (
	"renet/controllers"

	"github.com/go-chi/chi/v5"
)

func RegisterAuthRouter(r chi.Router, authController *controllers.AuthController) {
	r.Post("/signup", authController.SignUp)

}
