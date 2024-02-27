package logz

import (
	"fmt"
	
)

type StandardLogger interface {
	Log(level string, message string, args ...interface{})
	Info(message string, args ...interface{})
	Debug(message string, args ...interface{})
	Warning(message string, args ...interface{})
	Error(message string, args ...interface{})
	Fatal(message string, args ...interface{})
}


type Logger struct {
	Formatter	*Formatter
	Default 	*DefaultLogger
}

type DefaultLogger struct {
	Level string
	Message		string
	Args		[]interface{}
	Formatter	*Formatter
}

type Formatter struct {
	LogHolder	string
	Message		string
	Args		[]interface{}
	
}



func NewLogger(holder string, msg string, args ...interface{}) *Formatter {	
	format := &Formatter{
		LogHolder: holder,
		Message: msg,
		Args: args,
	} 
	return format
}

func DefaultLogs() *DefaultLogger {
	return &DefaultLogger{Formatter: &Formatter{}}
}


func (f Formatter) Format(level string, msg string, args ...interface{}) string {
	return fmt.Sprintf("[%v] %s", level, fmt.Sprintf(msg, args...) )
}

func (f *Formatter) Log() {
	formatted := f.Format(f.LogHolder, f.Message, f.Args...)
	fmt.Println(formatted)
}

func (l *DefaultLogger) INFO(msg string, args ...interface{}) {
	l.Formatter.LogHolder = "INFO" 
	l.Formatter.Message = msg 
	l.Formatter.Args = args 
	l.Formatter.Log()
}

func (l *DefaultLogger) DEBUG(msg string, args ...interface{}) {
	l.Formatter.LogHolder = "DEBUG" 
	l.Formatter.Message = msg 
	l.Formatter.Args = args 
	l.Formatter.Log()
}

func (l *DefaultLogger) WARNING(msg string, args ...interface{}) {
	l.Formatter.LogHolder = "WARNING" 
	l.Formatter.Message = msg 
	l.Formatter.Args = args 
	l.Formatter.Log()
}

func (l *DefaultLogger) ERROR(msg string, args ...interface{}) {
	l.Formatter.LogHolder = "ERROR" 
	l.Formatter.Message = msg 
	l.Formatter.Args = args 
	l.Formatter.Log()
}

func (l *DefaultLogger) FATAL(msg string, args ...interface{}) {
	l.Formatter.LogHolder = "FATAL" 
	l.Formatter.Message = msg 
	l.Formatter.Args = args 
	l.Formatter.Log()
}
