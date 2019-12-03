package iface

type Config struct {
	Level Level
}

func DefaultConfig() Config {
	return Config{
		Level: Trace,
	}
}
