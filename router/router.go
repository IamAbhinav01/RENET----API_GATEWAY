package router

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"renet/controllers"
	"renet/middlewares"
	"strings"

	"github.com/go-chi/chi/v5"
)

func createProxy(target string) *httputil.ReverseProxy {
	url, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ModifyResponse = func(response *http.Response) error {
		for name := range response.Header {
			if strings.HasPrefix(strings.ToLower(name), "access-control-") {
				response.Header.Del(name)
			}
		}
		return nil
	}
	return proxy
}

func createCatalogProxy(target string) http.Handler {
	return http.StripPrefix("/catalog", createProxy(target))
}

func createRecommendationProxy(target string) http.Handler {
	return http.StripPrefix("/recommend", createProxy(target))
}

func createStreamingProxy(target string) http.Handler {
	return http.StripPrefix("/streaming", createProxy(target))
}

func Router(Authcntrl *controllers.AuthController) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middlewares.CORSMiddleware)

	// Auth routes
	router.Route("/api/v1/auth", func(r chi.Router) {
		RegisterAuthRouter(r, Authcntrl)
	})

	// Reverse Proxies
	catalogHistoryProxy := createCatalogProxy("http://127.0.0.1:3000")
	router.With(
		middlewares.RequireSession(Authcntrl.SessionManager),
		middlewares.ForwardSessionUserID,
	).Handle("/catalog/api/history", catalogHistoryProxy)
	router.Handle("/catalog/*", createCatalogProxy("http://127.0.0.1:3000"))
	streamingProxy := createStreamingProxy("http://127.0.0.1:5000")
	router.Handle("/streaming/api/v1/videos/availability", streamingProxy)
	recommendationProxy := createRecommendationProxy("http://127.0.0.1:8000")
	router.With(
		middlewares.RequireSession(Authcntrl.SessionManager),
		middlewares.ForwardSessionUserID,
	).Handle("/recommend/api/user/recommend", recommendationProxy)
	router.Handle("/recommend/*", recommendationProxy)
	router.Handle("/streaming/health", streamingProxy)
	router.Handle("/streaming/streams/*", streamingProxy)
	router.Handle("/streaming/output/*", streamingProxy)
	router.Handle("/streaming/api/v1/videos/random*", streamingProxy)
	router.With(middlewares.RequireSession(Authcntrl.SessionManager)).Handle("/streaming/*", streamingProxy)

	return router
}
