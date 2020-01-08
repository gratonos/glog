package iface

type Config struct {
	Level         Level
	SrcPos        bool
	ConsoleWriter bool
	ConsoleConfig ConsoleConfig
}

type ConsoleConfig struct {
	ErrorHandler string
	Coloring     bool
}

func DefaultConfig() Config {
	return Config{
		Level:         Trace,
		SrcPos:        true,
		ConsoleWriter: true,
		ConsoleConfig: ConsoleConfig{
			Coloring: true,
		},
	}
}
