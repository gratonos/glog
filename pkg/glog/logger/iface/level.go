package iface

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

func LegalLogLevel(level Level) bool {
	return level >= Trace && level <= Fatal
}
