package logger

import (
	"fmt"

	ilogger "github.com/gratonos/glog/internal/logger"
	"github.com/gratonos/glog/pkg/glog/logger/iface"
)

type SugaredLogger struct {
	logger *ilogger.Logger
	pkg    string
	mark   bool
}

func NewSugaredLogger(logger *ilogger.Logger, pkg string) *SugaredLogger {
	return &SugaredLogger{
		logger: logger,
		pkg:    pkg,
	}
}

func (self SugaredLogger) WithMark() *SugaredLogger {
	self.mark = true
	return &self
}

func (this *SugaredLogger) Log(level iface.Level, args ...interface{}) {
	if !iface.LegalLogLevel(level) {
		panic(fmt.Sprintf("glog: illegal log level: %d", level))
	}
	this.log(level, args...)
}

func (this *SugaredLogger) Logf(level iface.Level, fmtstr string, args ...interface{}) {
	if !iface.LegalLogLevel(level) {
		panic(fmt.Sprintf("glog: illegal log level: %d", level))
	}
	this.logf(level, fmtstr, args...)
}

func (this *SugaredLogger) Trace(args ...interface{}) {
	this.log(iface.Trace, args...)
}

func (this *SugaredLogger) Tracef(fmtstr string, args ...interface{}) {
	this.logf(iface.Trace, fmtstr, args...)
}

func (this *SugaredLogger) Debug(args ...interface{}) {
	this.log(iface.Debug, args...)
}

func (this *SugaredLogger) Debugf(fmtstr string, args ...interface{}) {
	this.logf(iface.Debug, fmtstr, args...)
}

func (this *SugaredLogger) Info(args ...interface{}) {
	this.log(iface.Info, args...)
}

func (this *SugaredLogger) Infof(fmtstr string, args ...interface{}) {
	this.logf(iface.Info, fmtstr, args...)
}

func (this *SugaredLogger) Warn(args ...interface{}) {
	this.log(iface.Warn, args...)
}

func (this *SugaredLogger) Warnf(fmtstr string, args ...interface{}) {
	this.logf(iface.Warn, fmtstr, args...)
}

func (this *SugaredLogger) Error(args ...interface{}) {
	this.log(iface.Error, args...)
}

func (this *SugaredLogger) Errorf(fmtstr string, args ...interface{}) {
	this.logf(iface.Error, fmtstr, args...)
}

func (this *SugaredLogger) Fatal(args ...interface{}) {
	this.log(iface.Fatal, args...)
}

func (this *SugaredLogger) Fatalf(fmtstr string, args ...interface{}) {
	this.logf(iface.Fatal, fmtstr, args...)
}

func (this *SugaredLogger) Config() iface.Config {
	return this.logger.Config()
}

func (this *SugaredLogger) SetConfig(config iface.Config) error {
	return this.logger.SetConfig(config)
}

func (this *SugaredLogger) UpdateConfig(updater func(config iface.Config) iface.Config) error {
	return this.logger.UpdateConfig(updater)
}

func (this *SugaredLogger) log(level iface.Level, args ...interface{}) {
	log := this.genLog(level)
	if log != nil {
		log.Commit(fmt.Sprint(args...))
	}
}

func (this *SugaredLogger) logf(level iface.Level, fmtstr string, args ...interface{}) {
	log := this.genLog(level)
	if log != nil {
		log.Commit(fmt.Sprintf(fmtstr, args...))
	}
}

func (this *SugaredLogger) genLog(level iface.Level) *Log {
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
