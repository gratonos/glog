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

func (self Level) LegalForLog() bool {
	return self >= Trace && self <= Fatal
}

func (self Level) LegalForLogger() bool {
	return self.LegalForLog() || self == Off
}
