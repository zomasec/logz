package logz

import (
	"fmt"
)


const (
	Red     = "\033[0;31m"
	Blue    = "\033[0;34m"
	Green   = "\033[0;32m"
	Yellow  = "\033[0;33m"
	Purple  = "\033[0;35m"
	NC      = "\033[0m"    // No Color
)

// Formatter represents the log message formatter including the color.
type Formatter struct {
	LogHolder   string
	Message     string
	Args        []interface{}
	Color       string // Added to hold the color for the holder
}

// DefaultLogger now includes a map for custom log level colors.
type DefaultLogger struct {
	Level          string
	Message        string
	Args           []interface{}
	Formatter      *Formatter
	ColorEnabled   bool // Controls color output
	DebugEnabled   bool // Controls debug log visibility
	InfoEnabled    bool // Controls info log visibility
	WarningEnabled bool // Controls warning log visibility
	ErrorEnabled   bool // Controls error log visibility
	FatalEnabled   bool // Controls fatal log visibility
}

// NewLogger creates a new Logger instance with the provided log holder, message, color, and optional arguments.
func NewLogger(holder, color, msg string, args ...interface{}) *Formatter {
	return &Formatter{
		LogHolder: holder,
		Color:     color,
		Message:   msg,
		Args:      args,
	}
}

// DefaultLogs returns a new DefaultLogger instance with default configuration.
func DefaultLogs() *DefaultLogger {
	return &DefaultLogger{
		Formatter:     &Formatter{},
		ColorEnabled:  true, // Colors are enabled by default
		DebugEnabled:  true, // All log levels are visible by default
		InfoEnabled:   true,
		WarningEnabled: true,
		ErrorEnabled:  true,
		FatalEnabled:  true,
	}
}

func (f Formatter) Format() string {
	color := f.Color
	if color == "" {
		color = NC // Default to no color if none is specified
	}
	return fmt.Sprintf("[%s%s%s] %s", color, f.LogHolder,NC, fmt.Sprintf(f.Message, f.Args...))
}

// Log prints the formatted log message.
func (f *Formatter) Log() {
	fmt.Println(f.Format())
}

func (l *DefaultLogger) colorize(holder, color string) {
	if l.ColorEnabled {
		l.Formatter.LogHolder = fmt.Sprintf("%s%s%s", color, holder, NC)
	} else {
		l.Formatter.LogHolder = holder
	}
}

func (l *DefaultLogger) INFO(msg string, args ...interface{}) {
	if !l.InfoEnabled {
		return
	}
	l.colorize("INFO", Green)
	l.Formatter.Message = msg
	l.Formatter.Args = args
	l.Formatter.Log()
}

func (l *DefaultLogger) DEBUG(msg string, args ...interface{}) {
	if !l.DebugEnabled {
		return
	}
	l.colorize("DEBUG", Blue)
	l.Formatter.Message = msg
	l.Formatter.Args = args
	l.Formatter.Log()
}

func (l *DefaultLogger) WARNING(msg string, args ...interface{}) {
	if !l.WarningEnabled {
		return
	}
	l.colorize("WARNING", Yellow)
	l.Formatter.Message = msg
	l.Formatter.Args = args
	l.Formatter.Log()
}

func (l *DefaultLogger) ERROR(msg string, args ...interface{}) {
	if !l.ErrorEnabled {
		return
	}
	l.colorize("ERROR", Red)
	l.Formatter.Message = msg
	l.Formatter.Args = args
	l.Formatter.Log()
}

func (l *DefaultLogger) FATAL(msg string, args ...interface{}) {
	if !l.FatalEnabled {
		return
	}
	l.colorize("FATAL", Purple)
	l.Formatter.Message = msg
	l.Formatter.Args = args
	l.Formatter.Log()
}
