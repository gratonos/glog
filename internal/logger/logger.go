package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/gratonos/glog/pkg/glog/logger/iface"
)

type Logger struct {
	level  atomicLevel
	config iface.Config
}

func New() *Logger {
	config := iface.DefaultConfig()
	return &Logger{
		level: atomicLevel{
			level: config.Level,
		},
		config: config,
	}
}

func (this *Logger) Level() iface.Level {
	return this.level.Get()
}

func (this *Logger) Config() iface.Config {
	return this.config
}

func (this *Logger) SetConfig(config iface.Config) error {
	level := config.Level
	if !iface.LegalLoggerLevel(level) {
		return fmt.Errorf("glog: illegal logger level: %d", level)
	}

	this.level.Set(level)
	this.config = config

	return nil
}

func (this *Logger) UpdateConfig(updater func(config iface.Config) iface.Config) error {
	if updater == nil {
		panic("glog: updater is nil")
	}
	return this.SetConfig(updater(this.config))
}

func (this *Logger) Commit(emit func(time.Time) []byte, done func()) {
	tm := time.Now()
	log := emit(tm)

	os.Stderr.Write(log)

	done()
}
