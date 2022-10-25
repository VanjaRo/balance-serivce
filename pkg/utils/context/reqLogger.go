package context

import (
	"context"

	"github.com/sirupsen/logrus"
)

type contextKey int

const (
	_ contextKey = iota
	reqLoggerKey
	reqIdKey
)

func SetRequestLogger(ctx context.Context, logger *logrus.Logger) context.Context {
	return context.WithValue(ctx, reqLoggerKey, logger)
}

func GetRequestLogger(ctx context.Context) logrus.FieldLogger {
	logger := ctx.Value(reqLoggerKey)
	if logger != nil {
		return logger.(logrus.FieldLogger)
	}
	return logrus.New()
}

func SetReqId(ctx context.Context, reqId string) context.Context {
	return context.WithValue(ctx, reqIdKey, reqId)
}

func GetReqId(ctx context.Context) string {
	reqId := ctx.Value(reqIdKey)
	if reqId != nil {
		return reqId.(string)
	}
	return ""
}
