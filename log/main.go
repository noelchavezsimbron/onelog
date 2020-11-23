package log

import (
	"github.com/noelchavezsimbron/powerlog"
	"os"
	"time"
)

var logger = powerlog.New(os.Stdout, powerlog.ALL).Hook(func(e powerlog.Entry) {
	e.Int64("time", time.Now().Unix())
})

// Info prints a message with log level Info.
func Info(msg string) {
	logger.Info(msg)
}

// InfoWithFields prints a message with log level INFO and fields.
func InfoWithFields(msg string, fields func(e powerlog.Entry)) {
	logger.InfoWithFields(msg, fields)
}

// Debug prints a message with log level DEBUG.
func Debug(msg string) {
	logger.Debug(msg)
}

// DebugWithFields prints a message with log level DEBUG and fields.
func DebugWithFields(msg string, fields func(e powerlog.Entry)) {
	logger.DebugWithFields(msg, fields)
}

// Warn prints a message with log level INFO.
func Warn(msg string) {
	logger.Warn(msg)
}

// WarnWithFields prints a message with log level WARN and fields.
func WarnWithFields(msg string, fields func(e powerlog.Entry)) {
	logger.WarnWithFields(msg, fields)
}

// Error prints a message with log level ERROR.
func Error(msg string) {
	logger.Error(msg)
}

// ErrorWithFields prints a message with log level ERROR and fields.
func ErrorWithFields(msg string, fields func(e powerlog.Entry)) {
	logger.ErrorWithFields(msg, fields)
}

// Fatal prints a message with log level FATAL.
func Fatal(msg string) {
	logger.Fatal(msg)
}

// FatalWithFields prints a message with log level FATAL and fields.
func FatalWithFields(msg string, fields func(e powerlog.Entry)) {
	logger.FatalWithFields(msg, fields)
}
