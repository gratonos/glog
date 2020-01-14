package logger

import (
	"fmt"

	ilogger "github.com/gratonos/glog/internal/logger"
	"github.com/gratonos/glog/pkg/glog/logger/iface"
)

type Logger struct {
	logger *ilogger.Logger
	pkg    string
}

func NewLogger(logger *ilogger.Logger, pkg string) *Logger {
	return &Logger{
		logger: logger,
		pkg:    pkg,
	}
}

func (this *Logger) Log(level iface.Level, frameSkip int) *Log {
	if !iface.LegalLogLevel(level) {
		panic(fmt.Sprintf("glog: illegal log level: %d", level))
	}
	return genLog(this.logger, level, this.pkg, frameSkip+1)
}

func (this *Logger) Trace() *Log {
	return genLog(this.logger, iface.Trace, this.pkg, 0+1)
}

func (this *Logger) Debug() *Log {
	return genLog(this.logger, iface.Debug, this.pkg, 0+1)
}

func (this *Logger) Info() *Log {
	return genLog(this.logger, iface.Info, this.pkg, 0+1)
}

func (this *Logger) Warn() *Log {
	return genLog(this.logger, iface.Warn, this.pkg, 0+1)
}

func (this *Logger) Error() *Log {
	return genLog(this.logger, iface.Error, this.pkg, 0+1)
}

func (this *Logger) Fatal() *Log {
	return genLog(this.logger, iface.Fatal, this.pkg, 0+1)
}

func (this *Logger) Config() iface.Logger {
	return this.logger.Config()
}

func (this *Logger) SetConfig(config iface.Logger) error {
	return this.logger.SetConfig(config)
}

func (this *Logger) UpdateConfig(updater func(config iface.Logger) iface.Logger) error {
	return this.logger.UpdateConfig(updater)
}
