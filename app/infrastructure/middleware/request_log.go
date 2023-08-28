package middleware

import (
	"ahmadfarras/golang-http-base-template/app/configuration/logger"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	requestIdKey = "request_id"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := logger.InitLogger()

		requestId := uuid.New()

		ctx := context.WithValue(
			r.Context(),
			requestIdKey,
			requestId,
		)

		r = r.WithContext(ctx)

		l = l.With(zap.String(string(requestIdKey), requestId.String()))

		w.Header().Add("X-Request-ID", requestId.String())

		lrw := newLoggingResponseWriter(w)

		r = r.WithContext(logger.WithCtx(ctx, l))

		defer func(start time.Time) {
			l.Info(
				fmt.Sprintf(
					"%s request to %s completed",
					r.Method,
					r.RequestURI,
				),
				zap.String("method", r.Method),
				zap.String("url", r.RequestURI),
				zap.String("user_agent", r.UserAgent()),
				zap.Int("status_code", lrw.statusCode),
				zap.Duration("elapsed_ms", time.Since(start)),
			)
		}(time.Now())

		next.ServeHTTP(lrw, r)
	})
}
