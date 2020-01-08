package logger

import (
	"fmt"

	ilogger "github.com/gratonos/glog/internal/logger"
	"github.com/gratonos/glog/pkg/glog/logger/iface"
)

type SugaredLogger struct {
	logger *ilogger.Logger
	pkg    string
}

func NewSugaredLogger(logger *ilogger.Logger, pkg string) *SugaredLogger {
	return &SugaredLogger{
		logger: logger,
		pkg:    pkg,
	}
}

func (this *SugaredLogger) Log(level iface.Level, frameSkip int, args ...interface{}) {
	if !iface.LegalLogLevel(level) {
		panic(fmt.Sprintf("glog: illegal log level: %d", level))
	}
	this.log(level, frameSkip+1, args...)
}

func (this *SugaredLogger) Logf(level iface.Level, frameSkip int,
	fmtstr string, args ...interface{}) {

	if !iface.LegalLogLevel(level) {
		panic(fmt.Sprintf("glog: illegal log level: %d", level))
	}
	this.logf(level, frameSkip+1, fmtstr, args...)
}

func (this *SugaredLogger) Trace(args ...interface{}) {
	this.log(iface.Trace, 0+1, args...)
}

func (this *SugaredLogger) Tracef(fmtstr string, args ...interface{}) {
	this.logf(iface.Trace, 0+1, fmtstr, args...)
}

func (this *SugaredLogger) Debug(args ...interface{}) {
	this.log(iface.Debug, 0+1, args...)
}

func (this *SugaredLogger) Debugf(fmtstr string, args ...interface{}) {
	this.logf(iface.Debug, 0+1, fmtstr, args...)
}

func (this *SugaredLogger) Info(args ...interface{}) {
	this.log(iface.Info, 0+1, args...)
}

func (this *SugaredLogger) Infof(fmtstr string, args ...interface{}) {
	this.logf(iface.Info, 0+1, fmtstr, args...)
}

func (this *SugaredLogger) Warn(args ...interface{}) {
	this.log(iface.Warn, 0+1, args...)
}

func (this *SugaredLogger) Warnf(fmtstr string, args ...interface{}) {
	this.logf(iface.Warn, 0+1, fmtstr, args...)
}

func (this *SugaredLogger) Error(args ...interface{}) {
	this.log(iface.Error, 0+1, args...)
}

func (this *SugaredLogger) Errorf(fmtstr string, args ...interface{}) {
	this.logf(iface.Error, 0+1, fmtstr, args...)
}

func (this *SugaredLogger) Fatal(args ...interface{}) {
	this.log(iface.Fatal, 0+1, args...)
}

func (this *SugaredLogger) Fatalf(fmtstr string, args ...interface{}) {
	this.logf(iface.Fatal, 0+1, fmtstr, args...)
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

func (this *SugaredLogger) log(level iface.Level, frameSkip int, args ...interface{}) {
	log := genLog(this.logger, level, this.pkg, frameSkip+1)
	if log != nil {
		log.Commit(fmt.Sprint(args...))
	}
}

func (this *SugaredLogger) logf(level iface.Level, frameSkip int,
	fmtstr string, args ...interface{}) {

	log := genLog(this.logger, level, this.pkg, frameSkip+1)
	if log != nil {
		log.Commit(fmt.Sprintf(fmtstr, args...))
	}
}
