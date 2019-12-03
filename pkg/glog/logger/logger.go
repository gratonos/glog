package logger

import (
	"fmt"

	"github.com/gratonos/glog/internal/logger"
	"github.com/gratonos/glog/pkg/glog/logger/iface"
)

type Logger struct {
	log *logger.Logger
}

func New() *Logger {
	return &Logger{
		log: logger.New(),
	}
}

func (this *Logger) Log(level iface.Level) *logger.Log {
	if !iface.LegalLogLevel(level) {
		panic(fmt.Sprintf("glog: illegal log level: %d", level))
	}
	return this.log.GenLog(level)
}

func (this *Logger) Trace() *logger.Log {
	return this.log.GenLog(iface.Trace)
}

func (this *Logger) Debug() *logger.Log {
	return this.log.GenLog(iface.Debug)
}

func (this *Logger) Info() *logger.Log {
	return this.log.GenLog(iface.Info)
}

func (this *Logger) Warn() *logger.Log {
	return this.log.GenLog(iface.Warn)
}

func (this *Logger) Error() *logger.Log {
	return this.log.GenLog(iface.Error)
}

func (this *Logger) Fatal() *logger.Log {
	return this.log.GenLog(iface.Fatal)
}

func (this *Logger) Config() iface.Config {
	return this.log.Config()
}

func (this *Logger) SetConfig(config iface.Config) error {
	return this.log.SetConfig(config)
}

func (this *Logger) UpdateConfig(updater func(config iface.Config) iface.Config) error {
	return this.log.UpdateConfig(updater)
}
