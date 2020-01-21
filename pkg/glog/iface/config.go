package iface

type Logger struct {
	Level         Level
	FileLine      bool
	ConsoleWriter ConsoleWriter
	FileWriter    FileWriter
}

type ConsoleWriter struct {
	Enable       bool
	Coloring     bool
	ErrorHandler ErrorHandler
}

type FileWriter struct {
	Enable       bool
	MaxFileSize  int64
	Dir          string
	ErrorHandler ErrorHandler
}
