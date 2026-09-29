package middlewares

import (
	"context"
	"errors"
	"net/http"
	sessions "renet/redis"
	"renet/utils/formatters"

	goredis "github.com/redis/go-redis/v9"
)

type sessionContextKey struct{}

func RequireSession(sessionManager *sessions.SessionManager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("session_id")
			if err != nil || cookie.Value == "" {
				formatters.ErrorResponse(w, http.StatusUnauthorized, nil, "Authentication required")
				return
			}

			data, err := sessionManager.Store().Get(r.Context(), cookie.Value)
			if errors.Is(err, goredis.Nil) {
				formatters.ErrorResponse(w, http.StatusUnauthorized, nil, "Session is invalid or expired")
				return
			}
			if err != nil {
				formatters.ErrorResponse(w, http.StatusServiceUnavailable, err, "Unable to validate session")
				return
			}
			if data["user_id"] == "" {
				formatters.ErrorResponse(w, http.StatusUnauthorized, nil, "Session is invalid or expired")
				return
			}

			ctx := context.WithValue(r.Context(), sessionContextKey{}, data)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func SessionDataFromContext(ctx context.Context) (map[string]string, bool) {
	data, ok := ctx.Value(sessionContextKey{}).(map[string]string)
	return data, ok
}
