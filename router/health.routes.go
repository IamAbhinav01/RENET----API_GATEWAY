package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterHealthRoutes(r chi.Router, healthController http.HandlerFunc) {
	r.Get("/", healthController)
}
