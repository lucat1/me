package http

import (
	"context"
	"log/slog"
	"net/http"
)

const AUTH_CONTEXT_KEY = "auth"

type User struct {
	DN string `json:"string"`
}

func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "user", "fideloper")
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AuthMust(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(AUTH_CONTEXT_KEY).(*User)

		if !ok {
			// TODO: render page
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		slog.With("user", user).Debug("Granted authenticated request")
		h.ServeHTTP(w, r)
	})
}

// if must is set, GetUser can only be called in a http.Handler wrapped in a
// AuthMust middleware, but the nil checking on the returned variable can be avoided.
// Otherwise, the user should be always checked for existance (nil if unlogged)
func GetUser(w http.ResponseWriter, r *http.Request, must bool) (user *User) {
	user, ok := r.Context().Value(AUTH_CONTEXT_KEY).(*User)
	if !ok {
		user = nil
		if must {
			// TODO: render page
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	return user
}
