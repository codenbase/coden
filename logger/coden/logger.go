package coden

import (
	"github.com/codenbase/coden/log"
	"github.com/codenbase/coden/logger"
)

// codenLogger provides an implementation of the logger.Logger interface.
type codenLogger struct{}

// Ensure that codenLogger implements the logger.Logger interface.
var _ logger.Logger = (*codenLogger)(nil)

// NewLogger creates a new instance of codenLogger.
func NewLogger() *codenLogger {
	return &codenLogger{}
}

// Debug logs a debug message with any additional key-value pairs.
func (l *codenLogger) Debug(msg string, kvs ...any) {
	log.Debugw(msg, kvs...)
}

// Warn logs a warning message with any additional key-value pairs.
func (l *codenLogger) Warn(msg string, kvs ...any) {
	log.Warnw(msg, kvs...)
}

// Info logs an informational message with any additional key-value pairs.
func (l *codenLogger) Info(msg string, kvs ...any) {
	log.Infow(msg, kvs...)
}

// Error logs an error message with any additional key-value pairs.
func (l *codenLogger) Error(msg string, kvs ...any) {
	log.Errorw(nil, msg, kvs...)
}
