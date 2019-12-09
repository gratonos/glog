package logger

import (
	"fmt"

	ilogger "github.com/gratonos/glog/internal/logger"
	"github.com/gratonos/glog/pkg/glog/logger/iface"
)

type Logger struct {
	logger *ilogger.Logger
	pkg    string
	marker string
}

func NewLogger(logger *ilogger.Logger, pkg string) *Logger {
	return &Logger{
		logger: logger,
		pkg:    pkg,
	}
}

func (self Logger) WithMark() *Logger {
	self.marker = logMarker
	return &self
}

func (this *Logger) Log(level iface.Level, msg string) *Log {
	if !iface.LegalLogLevel(level) {
		panic(fmt.Sprintf("glog: illegal log level: %d", level))
	}
	return genLog(this.logger, level, this.pkg, this.marker, msg)
}

func (this *Logger) Trace(msg string) *Log {
	return genLog(this.logger, iface.Trace, this.pkg, this.marker, msg)
}

func (this *Logger) Debug(msg string) *Log {
	return genLog(this.logger, iface.Debug, this.pkg, this.marker, msg)
}

func (this *Logger) Info(msg string) *Log {
	return genLog(this.logger, iface.Info, this.pkg, this.marker, msg)
}

func (this *Logger) Warn(msg string) *Log {
	return genLog(this.logger, iface.Warn, this.pkg, this.marker, msg)
}

func (this *Logger) Error(msg string) *Log {
	return genLog(this.logger, iface.Error, this.pkg, this.marker, msg)
}

func (this *Logger) Fatal(msg string) *Log {
	return genLog(this.logger, iface.Fatal, this.pkg, this.marker, msg)
}

func (this *Logger) Config() iface.Config {
	return this.logger.Config()
}

func (this *Logger) SetConfig(config iface.Config) error {
	return this.logger.SetConfig(config)
}

func (this *Logger) UpdateConfig(updater func(config iface.Config) iface.Config) error {
	return this.logger.UpdateConfig(updater)
}
