package http

import (
	"log/slog"
	"net/http"
)

func LogMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.With("method", r.Method, "path", r.URL.Path).Debug("Received request")
		h.ServeHTTP(w, r)
	})
}

func RateLimitMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if rateLimitBreached(r) {
		//   TODO: render page
		//   w.WriteHeader(http.StatusTooManyRequests)
		//   return
		// }

		h.ServeHTTP(w, r)
	})
}
