package api

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Middleware struct {
	logger *zap.Logger
}

func NewMiddleware(logger *zap.Logger) *Middleware {
	return &Middleware{logger: logger}
}

func (m *Middleware) panicRecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				m.logger.Error("panic recover", zap.Any("err", err))
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (m *Middleware) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()

		next.ServeHTTP(w, r)

		m.logger.Info("request log",
			zap.String("leadTime", time.Since(timeStart).String()),
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path))
	})
}
