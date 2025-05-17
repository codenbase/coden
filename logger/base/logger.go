package base

import (
	"github.com/codenbase/coden/log"
	"github.com/codenbase/coden/logger"
)

// baseLogger provides an implementation of the logger.Logger interface.
type baseLogger struct{}

// Ensure that baseLogger implements the logger.Logger interface.
var _ logger.Logger = (*baseLogger)(nil)

// NewLogger creates a new instance of baseLogger.
func NewLogger() *baseLogger {
	return &baseLogger{}
}

// Debug logs a debug message with any additional key-value pairs.
func (l *baseLogger) Debug(msg string, kvs ...any) {
	log.Debugw(msg, kvs...)
}

// Warn logs a warning message with any additional key-value pairs.
func (l *baseLogger) Warn(msg string, kvs ...any) {
	log.Warnw(msg, kvs...)
}

// Info logs an informational message with any additional key-value pairs.
func (l *baseLogger) Info(msg string, kvs ...any) {
	log.Infow(msg, kvs...)
}

// Error logs an error message with any additional key-value pairs.
func (l *baseLogger) Error(msg string, kvs ...any) {
	log.Errorw(nil, msg, kvs...)
}
