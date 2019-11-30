package logger

type Level int32

const (
	Trace Level = iota
	Debug
	Info
	Warn
	Error
	Fatal
	Off
)

const (
	logBufLen   = 512
	timeHolder  = "00:00:00.000000"
	hourBegin   = 0
	hourEnd     = 2
	minuteBegin = 3
	minuteEnd   = 5
	secondBegin = 6
	secondEnd   = 8
	microBegin  = 9
	microEnd    = len(timeHolder)
)

const (
	stackOffset = 3
)

var levelNames = []string{
	Trace: "TRACE",
	Debug: "DEBUG",
	Info:  "INFO ",
	Warn:  "WARN ",
	Error: "ERROR",
	Fatal: "FATAL",
}
