package middlewares

import (
	"net/http"
	"renet/redis"
	"renet/utils/formatters"
)

func AuthMiddleware(sm *redis.SessionManager) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("session_id")

			if err != nil {
				formatters.ErrorResponse(w, http.StatusUnauthorized,  err,"Unauthorized - No session cookie")
				return
			}

			sessionData, err := sm.Store().Get(r.Context(), cookie.Value)

			if err != nil || sessionData == nil {
				formatters.ErrorResponse(w, http.StatusUnauthorized,  err,"Unauthorized - Invalid or expired session")
				return
			}

			r.Header.Set("X-user-id", sessionData["user_id"])

			next.ServeHTTP(w, r)
		})
	}

}