package http

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

type ContextKey string

const (
	REQUEST_ID_CONTEXT_KEY ContextKey = "requestId"
	LOGGER_CONTEXT_KEY     ContextKey = "logger"
	ERROR_CONTEXT_KEY      ContextKey = "error"
)

type ErrorPageData struct {
	Status  int
	Partial bool
	Error   error
	Message string
}

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

func Error(w http.ResponseWriter, r *http.Request, status int, partial bool, err error, message string) {
	res, ok := r.Context().Value(ERROR_CONTEXT_KEY).(*ErrorPageData)
	if !ok {
		panic("Could not get request error value from context")
	}

	w.WriteHeader(status)
	res.Status = status
	res.Partial = partial
	res.Error = err
	res.Message = message
}

func ErrorMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := ErrorPageData{
			Status: http.StatusOK,
		}
		ctx := context.WithValue(r.Context(), ERROR_CONTEXT_KEY, &err)
		h.ServeHTTP(w, r.WithContext(ctx))
		if err.Status != http.StatusOK {
			logger := GetLogger(r)
			logger.With("err", err.Error).Error(err.Message)

			var renderErr error
			if err.Partial {
				renderErr = RenderBlockPage[ErrorPageData](w, r, "error", err)
			} else {
				renderErr = RenderPage[ErrorPageData](w, r, "error", "Error", err)
			}
			if renderErr != nil {
				logger.With("err", renderErr).Error("Could not render page")
			}
		}
	})
}
