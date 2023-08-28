package logger

import (
	"ahmadfarras/golang-http-base-template/app/configuration/constant"
	"context"
	"os"
	"sync"

	"go.uber.org/zap"
)

type ctxKey struct{}

var (
	once   sync.Once
	logger *zap.Logger
)

func InitLogger() *zap.Logger {

	once.Do(func() {
		if os.Getenv(constant.APP_ENV) == "PRODUCTION" {
			logger = zap.Must(zap.NewProduction())
		} else {
			logger = zap.Must(zap.NewDevelopment())
		}
	})

	return logger
}

func FromCtx(ctx context.Context) *zap.Logger {
	if l, ok := ctx.Value(ctxKey{}).(*zap.Logger); ok {
		return l
	} else if l := logger; l != nil {
		return l
	}

	return zap.NewNop()
}

func WithCtx(ctx context.Context, l *zap.Logger) context.Context {
	if lp, ok := ctx.Value(ctxKey{}).(*zap.Logger); ok {
		if lp == l {
			return ctx
		}
	}

	return context.WithValue(ctx, ctxKey{}, l)
}
