package base

import (
	"context"

	"github.com/codenbase/coden/log"
)

// baseLogger is a logger that implements the Logger interface.
// It uses the log package to log error messages with additional context.
type baseLogger struct{}

// NewLogger creates and returns a new instance of baseLogger.
func NewLogger() *baseLogger {
	return &baseLogger{}
}

// Error logs an error message with the provided context using the log package.
func (l *baseLogger) Error(ctx context.Context, err error, msg string, kvs ...any) {
	log.W(ctx).Errorw(err, msg, kvs...)
}
