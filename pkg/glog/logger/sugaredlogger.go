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
	this.genLog(level).Commit(fmt.Sprint(args...))
}

func (this *SugaredLogger) Logf(level iface.Level, fmtstr string, args ...interface{}) {
	if !iface.LegalLogLevel(level) {
		panic(fmt.Sprintf("glog: illegal log level: %d", level))
	}
	this.genLog(level).Commit(fmt.Sprintf(fmtstr, args...))
}

func (this *SugaredLogger) Trace(args ...interface{}) {
	this.genLog(iface.Trace).Commit(fmt.Sprint(args...))
}

func (this *SugaredLogger) Tracef(fmtstr string, args ...interface{}) {
	this.genLog(iface.Trace).Commit(fmt.Sprintf(fmtstr, args...))
}

func (this *SugaredLogger) Debug(args ...interface{}) {
	this.genLog(iface.Debug).Commit(fmt.Sprint(args...))
}

func (this *SugaredLogger) Debugf(fmtstr string, args ...interface{}) {
	this.genLog(iface.Debug).Commit(fmt.Sprintf(fmtstr, args...))
}

func (this *SugaredLogger) Info(args ...interface{}) {
	this.genLog(iface.Info).Commit(fmt.Sprint(args...))
}

func (this *SugaredLogger) Infof(fmtstr string, args ...interface{}) {
	this.genLog(iface.Info).Commit(fmt.Sprintf(fmtstr, args...))
}

func (this *SugaredLogger) Warn(args ...interface{}) {
	this.genLog(iface.Warn).Commit(fmt.Sprint(args...))
}

func (this *SugaredLogger) Warnf(fmtstr string, args ...interface{}) {
	this.genLog(iface.Warn).Commit(fmt.Sprintf(fmtstr, args...))
}

func (this *SugaredLogger) Error(args ...interface{}) {
	this.genLog(iface.Error).Commit(fmt.Sprint(args...))
}

func (this *SugaredLogger) Errorf(fmtstr string, args ...interface{}) {
	this.genLog(iface.Error).Commit(fmt.Sprintf(fmtstr, args...))
}

func (this *SugaredLogger) Fatal(args ...interface{}) {
	this.genLog(iface.Fatal).Commit(fmt.Sprint(args...))
}

func (this *SugaredLogger) Fatalf(fmtstr string, args ...interface{}) {
	this.genLog(iface.Fatal).Commit(fmt.Sprintf(fmtstr, args...))
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
