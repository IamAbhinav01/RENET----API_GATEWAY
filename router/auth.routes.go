package router

import (
	"renet/controllers"
	"renet/middlewares"

	"github.com/go-chi/chi/v5"
)

func RegisterAuthRouter(r chi.Router, authController *controllers.AuthController) {
	r.With(middlewares.SignUpRequestValidation).Post("/signup", authController.SignUp)
	r.With(middlewares.SignInRequestValidation).Post("/login", authController.Login)
	r.With(middlewares.RequireSession(authController.SessionManager)).Get("/me", authController.CurrentUser)
	r.With(middlewares.RequireSession(authController.SessionManager)).Post("/logout", authController.LogOut)
}
