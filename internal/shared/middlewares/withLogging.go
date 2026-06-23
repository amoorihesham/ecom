package middlewares

import (
	"log/slog"
	"net/http"
	"time"
)

func WithLogging(logger *slog.Logger) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			reqId, ok := r.Context().Value(ReqIdKey).(string)
			if !ok {
				reqId = "unknown"
			}

			logger.Info(
				"incoming request",
				"req-id",
				reqId,
				"method", r.Method,
				"path", r.URL.Path,
			)

			next(w, r)

			logger.Info("request complete", "method", r.Method, "path", r.URL.Path, "duration", time.Since(start))
		})
	}
}
