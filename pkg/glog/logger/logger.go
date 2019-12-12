package logger

import (
	"fmt"

	ilogger "github.com/gratonos/glog/internal/logger"
	"github.com/gratonos/glog/pkg/glog/logger/iface"
)

type Logger struct {
	logger *ilogger.Logger
	pkg    string
	mark   bool
}

func NewLogger(logger *ilogger.Logger, pkg string) *Logger {
	return &Logger{
		logger: logger,
		pkg:    pkg,
	}
}

func (self Logger) WithMark() *Logger {
	self.mark = true
	return &self
}

func (this *Logger) Log(level iface.Level) *Log {
	if !iface.LegalLogLevel(level) {
		panic(fmt.Sprintf("glog: illegal log level: %d", level))
	}
	return this.genLog(level)
}

func (this *Logger) Trace() *Log {
	return this.genLog(iface.Trace)
}

func (this *Logger) Debug() *Log {
	return this.genLog(iface.Debug)
}

func (this *Logger) Info() *Log {
	return this.genLog(iface.Info)
}

func (this *Logger) Warn() *Log {
	return this.genLog(iface.Warn)
}

func (this *Logger) Error() *Log {
	return this.genLog(iface.Error)
}

func (this *Logger) Fatal() *Log {
	return this.genLog(iface.Fatal)
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

func (this *Logger) genLog(level iface.Level) *Log {
	if this.logger.Level() > level {
		return nil
	}

	info := &preInfo{
		Pkg:   this.pkg,
		Level: uint8(level),
		Mark:  this.mark,
	}
	return genLog(this.logger, info)
}
