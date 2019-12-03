package logger

import (
	"fmt"

	"github.com/gratonos/glog/internal/logger"
	"github.com/gratonos/glog/pkg/glog/logger/iface"
)

type Logger struct {
	log    *logger.Logger
	pkg    string
	marker string
}

func New(log *logger.Logger, pkg string) *Logger {
	return &Logger{
		log: log,
		pkg: pkg,
	}
}

func (self Logger) WithMark() *Logger {
	self.marker = logMarker
	return &self
}

func (this *Logger) Log(level iface.Level, msg string) *logger.Log {
	if !iface.LegalLogLevel(level) {
		panic(fmt.Sprintf("glog: illegal log level: %d", level))
	}
	return this.log.GenLog(level, this.pkg, this.marker, msg)
}

func (this *Logger) Trace(msg string) *logger.Log {
	return this.log.GenLog(iface.Trace, this.pkg, this.marker, msg)
}

func (this *Logger) Debug(msg string) *logger.Log {
	return this.log.GenLog(iface.Debug, this.pkg, this.marker, msg)
}

func (this *Logger) Info(msg string) *logger.Log {
	return this.log.GenLog(iface.Info, this.pkg, this.marker, msg)
}

func (this *Logger) Warn(msg string) *logger.Log {
	return this.log.GenLog(iface.Warn, this.pkg, this.marker, msg)
}

func (this *Logger) Error(msg string) *logger.Log {
	return this.log.GenLog(iface.Error, this.pkg, this.marker, msg)
}

func (this *Logger) Fatal(msg string) *logger.Log {
	return this.log.GenLog(iface.Fatal, this.pkg, this.marker, msg)
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
