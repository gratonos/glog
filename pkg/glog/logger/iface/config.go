package iface

type Logger struct {
	Level         Level
	FileLine      bool
	ConsoleWriter ConsoleWriter
}

type ConsoleWriter struct {
	Enable       bool
	Coloring     bool
	ErrorHandler ErrorHandler
}
