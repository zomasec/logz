package logger


type Level int8 

const (
	Info Level = iota
	Debug
	Warn
	Error
	Fatal
	Test
)

func (l Level) String() string {
	return [...]string{"INF", "DBG", "WRN", "ERR", "FTL", "TST"}[l]
}