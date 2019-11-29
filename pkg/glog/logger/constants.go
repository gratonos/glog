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
	logBufLen  = 256
	timeLayout = "2006-01-02 15:04:05.000000"
	timeLen    = len(timeLayout)
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
