package handler

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

type contextKey string

const requestIDKey contextKey = "request_id"

func withMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
	return requestIDMiddleware(loggingMiddleware(logger, next))
}

func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			// TODO: Replace this with a stronger ID generator and make sure the
			// value is passed to logs and responses consistently.
			requestID = strconv.FormatInt(time.Now().UnixNano(), 10)
		}

		ctx := context.WithValue(r.Context(), requestIDKey, requestID)
		w.Header().Set("X-Request-ID", requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func loggingMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)

		// TODO: Log status code, duration, request ID, and any useful domain
		// fields. Start simple, then decide what is actually worth emitting.
		logger.Info("request finished",
			"method", r.Method,
			"path", r.URL.Path,
			"duration_ms", time.Since(start).Milliseconds(),
		)
	})
}
