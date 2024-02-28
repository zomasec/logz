// Package logz provides a simple logging functionality with different log levels.
package logz

import (
	"fmt"
)

// StandardLogger defines the interface for a standard logger.
type StandardLogger interface {
	Log(level string, message string, args ...interface{})
	Info(message string, args ...interface{})
	Debug(message string, args ...interface{})
	Warning(message string, args ...interface{})
	Error(message string, args ...interface{})
	Fatal(message string, args ...interface{})
}

// Logger represents a logger object which holds a formatter and default logger instance.
type Logger struct {
	Formatter *Formatter
	Default   *DefaultLogger
}

// DefaultLogger represents the default logger configuration.
type DefaultLogger struct {
	Level     string
	Message   string
	Args      []interface{}
	Formatter *Formatter
}

// Formatter represents the log message formatter.
type Formatter struct {
	LogHolder string
	Message   string
	Args      []interface{}
}

// NewLogger creates a new Logger instance with the provided log holder, message, and optional arguments.
func NewLogger(holder string, msg string, args ...interface{}) *Formatter {
	format := &Formatter{
		LogHolder: holder,
		Message:   msg,
		Args:      args,
	}
	return format
}

// DefaultLogs returns a new DefaultLogger instance with default configuration.
func DefaultLogs() *DefaultLogger {
	return &DefaultLogger{Formatter: &Formatter{}}
}

// Format formats the log message with the provided level, message, and optional arguments.
func (f Formatter) Format(level string, msg string, args ...interface{}) string {
	return fmt.Sprintf("[%v] %s", level, fmt.Sprintf(msg, args...))
}

// Log prints the formatted log message.
func (f *Formatter) Log() {
	formatted := f.Format(f.LogHolder, f.Message, f.Args...)
	fmt.Println(formatted)
}

// INFO logs a message at INFO level.
func (l *DefaultLogger) INFO(msg string, args ...interface{}) {
	l.Formatter.LogHolder = "INFO"
	l.Formatter.Message = msg
	l.Formatter.Args = args
	l.Formatter.Log()
}

// DEBUG logs a message at DEBUG level.
func (l *DefaultLogger) DEBUG(msg string, args ...interface{}) {
	l.Formatter.LogHolder = "DEBUG"
	l.Formatter.Message = msg
	l.Formatter.Args = args
	l.Formatter.Log()
}

// WARNING logs a message at WARNING level.
func (l *DefaultLogger) WARNING(msg string, args ...interface{}) {
	l.Formatter.LogHolder = "WARNING"
	l.Formatter.Message = msg
	l.Formatter.Args = args
	l.Formatter.Log()
}

// ERROR logs a message at ERROR level.
func (l *DefaultLogger) ERROR(msg string, args ...interface{}) {
	l.Formatter.LogHolder = "ERROR"
	l.Formatter.Message = msg
	l.Formatter.Args = args
	l.Formatter.Log()
}

// FATAL logs a message at FATAL level.
func (l *DefaultLogger) FATAL(msg string, args ...interface{}) {
	l.Formatter.LogHolder = "FATAL"
	l.Formatter.Message = msg
	l.Formatter.Args = args
	l.Formatter.Log()
}
