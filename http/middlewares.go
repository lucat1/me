package http

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

const (
	REQUEST_ID_CONTEXT_KEY = "requestId"
	LOGGER_CONTEXT_KEY     = "logger"
)

func IdMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New()
		ctx := context.WithValue(r.Context(), REQUEST_ID_CONTEXT_KEY, id)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetRequestId(r *http.Request) uuid.UUID {
	res, ok := r.Context().Value(REQUEST_ID_CONTEXT_KEY).(uuid.UUID)
	if !ok {
		panic("Could not get request id from context")
	}
	return res
}

func LogMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := slog.
			New(slog.Default().Handler()).
			With(
				"id", GetRequestId(r),
				"method", r.Method,
				"path", r.URL.Path,
			)

		logger.Debug("Started handling request")
		ctx := context.WithValue(r.Context(), LOGGER_CONTEXT_KEY, logger)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetLogger(r *http.Request) *slog.Logger {
	res, ok := r.Context().Value(LOGGER_CONTEXT_KEY).(*slog.Logger)
	if !ok {
		panic("Could not get request logger from context")
	}
	return res
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
