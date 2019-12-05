package logger

import (
	"fmt"

	"github.com/gratonos/glog/internal/logger"
	"github.com/gratonos/glog/pkg/glog/logger/iface"
)

type SugaredLogger struct {
	log    *logger.Logger
	pkg    string
	marker string
}

func NewSugaredLogger(log *logger.Logger, pkg string) *SugaredLogger {
	return &SugaredLogger{
		log: log,
		pkg: pkg,
	}
}

func (self SugaredLogger) WithMark() *SugaredLogger {
	self.marker = logMarker
	return &self
}

func (this *SugaredLogger) Log(level iface.Level, args ...interface{}) {
	if !iface.LegalLogLevel(level) {
		panic(fmt.Sprintf("glog: illegal log level: %d", level))
	}
	this.log.GenLog(level, this.pkg, this.marker, fmt.Sprint(args...)).Commit()
}

func (this *SugaredLogger) Logf(level iface.Level, fmtstr string, args ...interface{}) {
	if !iface.LegalLogLevel(level) {
		panic(fmt.Sprintf("glog: illegal log level: %d", level))
	}
	this.log.GenLog(level, this.pkg, this.marker, fmt.Sprintf(fmtstr, args...)).Commit()
}

func (this *SugaredLogger) Trace(args ...interface{}) {
	this.log.GenLog(iface.Trace, this.pkg, this.marker, fmt.Sprint(args...)).Commit()
}

func (this *SugaredLogger) Tracef(fmtstr string, args ...interface{}) {
	this.log.GenLog(iface.Trace, this.pkg, this.marker, fmt.Sprintf(fmtstr, args...)).Commit()
}

func (this *SugaredLogger) Debug(args ...interface{}) {
	this.log.GenLog(iface.Debug, this.pkg, this.marker, fmt.Sprint(args...)).Commit()
}

func (this *SugaredLogger) Debugf(fmtstr string, args ...interface{}) {
	this.log.GenLog(iface.Debug, this.pkg, this.marker, fmt.Sprintf(fmtstr, args...)).Commit()
}

func (this *SugaredLogger) Info(args ...interface{}) {
	this.log.GenLog(iface.Info, this.pkg, this.marker, fmt.Sprint(args...)).Commit()
}

func (this *SugaredLogger) Infof(fmtstr string, args ...interface{}) {
	this.log.GenLog(iface.Info, this.pkg, this.marker, fmt.Sprintf(fmtstr, args...)).Commit()
}

func (this *SugaredLogger) Warn(args ...interface{}) {
	this.log.GenLog(iface.Warn, this.pkg, this.marker, fmt.Sprint(args...)).Commit()
}

func (this *SugaredLogger) Warnf(fmtstr string, args ...interface{}) {
	this.log.GenLog(iface.Warn, this.pkg, this.marker, fmt.Sprintf(fmtstr, args...)).Commit()
}

func (this *SugaredLogger) Error(args ...interface{}) {
	this.log.GenLog(iface.Error, this.pkg, this.marker, fmt.Sprint(args...)).Commit()
}

func (this *SugaredLogger) Errorf(fmtstr string, args ...interface{}) {
	this.log.GenLog(iface.Error, this.pkg, this.marker, fmt.Sprintf(fmtstr, args...)).Commit()
}

func (this *SugaredLogger) Fatal(args ...interface{}) {
	this.log.GenLog(iface.Fatal, this.pkg, this.marker, fmt.Sprint(args...)).Commit()
}

func (this *SugaredLogger) Fatalf(fmtstr string, args ...interface{}) {
	this.log.GenLog(iface.Fatal, this.pkg, this.marker, fmt.Sprintf(fmtstr, args...)).Commit()
}

func (this *SugaredLogger) Config() iface.Config {
	return this.log.Config()
}

func (this *SugaredLogger) SetConfig(config iface.Config) error {
	return this.log.SetConfig(config)
}

func (this *SugaredLogger) UpdateConfig(updater func(config iface.Config) iface.Config) error {
	return this.log.UpdateConfig(updater)
}
