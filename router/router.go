package router

import (
	"renet/controllers"

	"github.com/go-chi/chi/v5"
)


func Router()*chi.Mux{
	router := chi.NewRouter()

	router.Route("/api/v1/health",func(r chi.Router){
		RegisterHealthRoutes(r,controllers.HealthController)
	})
	router.Route("/api/v1/login",func(r chi.Router) {
		RegisterAuthRouter(r,controllers.HealthController)
	})

	return router
}