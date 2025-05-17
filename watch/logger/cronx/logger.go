package cronx

import (
	"github.com/codenbase/coden/log"
)

// cronxLogger implement the cron.Logger interface.
type cronxLogger struct{}

// NewLogger returns a cron logger.
func NewLogger() *cronxLogger {
	return &cronxLogger{}
}

// Debug logs routine messages about cron's operation.
func (l *cronxLogger) Debug(msg string, kvs ...any) {
	log.Debugw(msg, kvs...)
}

// Info logs routine messages about cron's operation.
func (l *cronxLogger) Info(msg string, kvs ...any) {
	log.Infow(msg, kvs...)
}

// Error logs an error condition.
func (l *cronxLogger) Error(err error, msg string, kvs ...any) {
	log.Errorw(err, msg, kvs...)
}
