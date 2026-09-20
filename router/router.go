package router

import (
	"renet/controllers"

	"github.com/go-chi/chi/v5"
)

func Router(Authcntrl *controllers.AuthController) *chi.Mux {
	router := chi.NewRouter()

	router.Route("/api/v1/auth", func(r chi.Router) {
		RegisterAuthRouter(r, Authcntrl)
	})

	return router
}
