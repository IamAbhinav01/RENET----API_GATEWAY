package router

import (
	"net/http"
	"net/http/httptest"
	"renet/middlewares"
	"testing"
)

func TestCatalogProxyStripsPrefixAndPreservesQuery(t *testing.T) {
	var gotPath string
	var gotQuery string
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotQuery = r.URL.RawQuery
	}))
	defer backend.Close()

	request := httptest.NewRequest(http.MethodGet, "/catalog/api/movies?page=1&limit=20", nil)
	response := httptest.NewRecorder()
	createCatalogProxy(backend.URL).ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("proxy returned status %d, want %d", response.Code, http.StatusOK)
	}
	if gotPath != "/api/movies" {
		t.Errorf("backend received path %q, want %q", gotPath, "/api/movies")
	}
	if gotQuery != "page=1&limit=20" {
		t.Errorf("backend received query %q, want %q", gotQuery, "page=1&limit=20")
	}
}

func TestProxyUsesSingleCredentialedCORSOrigin(t *testing.T) {
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.WriteHeader(http.StatusOK)
	}))
	defer backend.Close()

	request := httptest.NewRequest(http.MethodGet, "/recommend/api/recommend", nil)
	request.Header.Set("Origin", "http://localhost:5500")
	response := httptest.NewRecorder()
	middlewares.CORSMiddleware(createProxy(backend.URL)).ServeHTTP(response, request)

	origins := response.Header().Values("Access-Control-Allow-Origin")
	if len(origins) != 1 || origins[0] != "http://localhost:5500" {
		t.Fatalf("Access-Control-Allow-Origin = %#v, want one echoed origin", origins)
	}
	if got := response.Header().Values("Access-Control-Allow-Credentials"); len(got) != 1 || got[0] != "true" {
		t.Fatalf("Access-Control-Allow-Credentials = %#v, want one true value", got)
	}
}

func TestRecommendationProxyStripsPrefixAndPreservesQuery(t *testing.T) {
	var gotPath string
	var gotQuery string
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotQuery = r.URL.RawQuery
	}))
	defer backend.Close()

	request := httptest.NewRequest(http.MethodGet, "/recommend/api/recommend?movie_name=Toy+Story&n=8", nil)
	response := httptest.NewRecorder()
	createRecommendationProxy(backend.URL).ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("proxy returned status %d, want %d", response.Code, http.StatusOK)
	}
	if gotPath != "/api/recommend" {
		t.Errorf("backend received path %q, want %q", gotPath, "/api/recommend")
	}
	if gotQuery != "movie_name=Toy+Story&n=8" {
		t.Errorf("backend received query %q, want %q", gotQuery, "movie_name=Toy+Story&n=8")
	}
}
