package logger


type Color int8 

const (
	Red Color = iota
	Blue
	Green
	Yellow
	Purple
	Cyan
	Gray
	Orange
	Reset
)

var Colors = map[Color]string{
	Red:    "\033[0;31m",
	Blue:   "\033[0;34m",
	Green:  "\033[0;32m",
	Yellow: "\033[0;33m",
	Purple: "\033[0;35m",
	Cyan:   "\033[0;36m",
	Gray:   "\033[0;37m",
	Orange: "\033[0;91m",
	Reset:   "\033[0m",
}
