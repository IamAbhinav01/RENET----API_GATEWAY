package router

import (
	"github.com/go-chi/chi/v5"
)

func Router()*chi.Mux{
	router := chi.NewRouter()

	router.Route("/ap1/v1/health",func(r chi.Router){
		RegisterHealthRoutes(r)
	})

}