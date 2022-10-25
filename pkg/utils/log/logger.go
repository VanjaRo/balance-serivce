package log

import (
	"context"
	"os"

	reqContext "github.com/VanjaRo/balance-serivce/pkg/utils/context"
	"github.com/sirupsen/logrus"
)

func New() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetOutput(os.Stdout)
	return logger
}

// Debug wraps logrus Debugf such that logger for the given context
func Debug(ctx context.Context, msg string, args ...interface{}) {
	getLogger(ctx).Debugf(msg, args...)
}

// Info wraps logrus Infof such that logger for the given context
func Info(ctx context.Context, msg string, args ...interface{}) {
	getLogger(ctx).Infof(msg, args...)
}

// Warn wraps logrus Warnf such that logger for the given context
func Warn(ctx context.Context, msg string, args ...interface{}) {
	getLogger(ctx).Warnf(msg, args...)
}

// Error wraps logrus Warnf such that logger for the given context
func Error(ctx context.Context, msg string, args ...interface{}) {
	getLogger(ctx).Errorf(msg, args...)
}

// Fatal wraps logrus Fatalf such that logger for the given context
func Fatal(ctx context.Context, msg string, args ...interface{}) {
	getLogger(ctx).Fatalf(msg, args...)
}

// Panic wraps logrus Panicf such that logger for the given context
func Panic(ctx context.Context, msg string, args ...interface{}) {
	getLogger(ctx).Panicf(msg, args...)
}

func getLogger(ctx context.Context) logrus.FieldLogger {
	if ctx != nil {
		return reqContext.GetRequestLogger(ctx)
	}
	return New()
}
