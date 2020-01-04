package iface

type Config struct {
	Level         Level
	ConsoleWriter bool
	ConsoleConfig ConsoleConfig
}

type ConsoleConfig struct {
	Coloring bool
}

func DefaultConfig() Config {
	return Config{
		Level:         Trace,
		ConsoleWriter: true,
		ConsoleConfig: ConsoleConfig{
			Coloring: true,
		},
	}
}
