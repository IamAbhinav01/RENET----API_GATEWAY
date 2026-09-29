package router

import (
	"net/http/httputil"
	"net/url"
	"renet/controllers"
	"renet/middlewares"

	"github.com/go-chi/chi/v5"
)

func createProxy(target string) *httputil.ReverseProxy {
	url, _ := url.Parse(target)
	return httputil.NewSingleHostReverseProxy(url)
}

func Router(Authcntrl *controllers.AuthController) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middlewares.CORSMiddleware)

	// Auth routes
	router.Route("/api/v1/auth", func(r chi.Router) {
		RegisterAuthRouter(r, Authcntrl)
	})

	// Reverse Proxies
	router.Handle("/catalog/*", createProxy("http://localhost:3000"))
	router.Handle("/recommend/*", createProxy("http://localhost:8000"))
	router.Handle("/streaming/*", createProxy("http://localhost:5000"))

	return router
}
