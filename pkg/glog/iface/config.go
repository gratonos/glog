package iface

type Logger struct {
	Level         Level
	FileLine      bool
	ConsoleWriter ConsoleWriter
	FileWriter    FileWriter
}

type ConsoleWriter struct {
	Enable       bool
	TextConfig   TextConfig
	ErrorHandler ErrorHandler
}

type FileWriter struct {
	Enable       bool
	MaxFileSize  int64
	Dir          string
	ErrorHandler ErrorHandler
}

type TextConfig struct {
	Coloring bool
}
