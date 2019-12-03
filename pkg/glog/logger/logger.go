package logger

import (
	"fmt"

	"github.com/gratonos/glog/internal/logger"
	"github.com/gratonos/glog/pkg/glog/logger/iface"
)

type Logger struct {
	log *logger.Logger
	pkg string
}

func New(log *logger.Logger, pkg string) *Logger {
	return &Logger{
		log: log,
		pkg: pkg,
	}
}

func (this *Logger) Log(level iface.Level) *logger.Log {
	if !iface.LegalLogLevel(level) {
		panic(fmt.Sprintf("glog: illegal log level: %d", level))
	}
	return this.log.GenLog(level, this.pkg)
}

func (this *Logger) Trace() *logger.Log {
	return this.log.GenLog(iface.Trace, this.pkg)
}

func (this *Logger) Debug() *logger.Log {
	return this.log.GenLog(iface.Debug, this.pkg)
}

func (this *Logger) Info() *logger.Log {
	return this.log.GenLog(iface.Info, this.pkg)
}

func (this *Logger) Warn() *logger.Log {
	return this.log.GenLog(iface.Warn, this.pkg)
}

func (this *Logger) Error() *logger.Log {
	return this.log.GenLog(iface.Error, this.pkg)
}

func (this *Logger) Fatal() *logger.Log {
	return this.log.GenLog(iface.Fatal, this.pkg)
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
