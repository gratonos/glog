package logger

import (
	"fmt"

	ilogger "github.com/gratonos/glog/internal/logger"
	"github.com/gratonos/glog/pkg/glog/logger/iface"
)

type SugaredLogger struct {
	logger *ilogger.Logger
	pkg    string
	marker string
}

func NewSugaredLogger(logger *ilogger.Logger, pkg string) *SugaredLogger {
	return &SugaredLogger{
		logger: logger,
		pkg:    pkg,
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
	genLog(this.logger, level, this.pkg, this.marker, fmt.Sprint(args...)).Commit()
}

func (this *SugaredLogger) Logf(level iface.Level, fmtstr string, args ...interface{}) {
	if !iface.LegalLogLevel(level) {
		panic(fmt.Sprintf("glog: illegal log level: %d", level))
	}
	genLog(this.logger, level, this.pkg, this.marker, fmt.Sprintf(fmtstr, args...)).Commit()
}

func (this *SugaredLogger) Trace(args ...interface{}) {
	genLog(this.logger, iface.Trace, this.pkg, this.marker, fmt.Sprint(args...)).Commit()
}

func (this *SugaredLogger) Tracef(fmtstr string, args ...interface{}) {
	genLog(this.logger, iface.Trace, this.pkg, this.marker, fmt.Sprintf(fmtstr, args...)).Commit()
}

func (this *SugaredLogger) Debug(args ...interface{}) {
	genLog(this.logger, iface.Debug, this.pkg, this.marker, fmt.Sprint(args...)).Commit()
}

func (this *SugaredLogger) Debugf(fmtstr string, args ...interface{}) {
	genLog(this.logger, iface.Debug, this.pkg, this.marker, fmt.Sprintf(fmtstr, args...)).Commit()
}

func (this *SugaredLogger) Info(args ...interface{}) {
	genLog(this.logger, iface.Info, this.pkg, this.marker, fmt.Sprint(args...)).Commit()
}

func (this *SugaredLogger) Infof(fmtstr string, args ...interface{}) {
	genLog(this.logger, iface.Info, this.pkg, this.marker, fmt.Sprintf(fmtstr, args...)).Commit()
}

func (this *SugaredLogger) Warn(args ...interface{}) {
	genLog(this.logger, iface.Warn, this.pkg, this.marker, fmt.Sprint(args...)).Commit()
}

func (this *SugaredLogger) Warnf(fmtstr string, args ...interface{}) {
	genLog(this.logger, iface.Warn, this.pkg, this.marker, fmt.Sprintf(fmtstr, args...)).Commit()
}

func (this *SugaredLogger) Error(args ...interface{}) {
	genLog(this.logger, iface.Error, this.pkg, this.marker, fmt.Sprint(args...)).Commit()
}

func (this *SugaredLogger) Errorf(fmtstr string, args ...interface{}) {
	genLog(this.logger, iface.Error, this.pkg, this.marker, fmt.Sprintf(fmtstr, args...)).Commit()
}

func (this *SugaredLogger) Fatal(args ...interface{}) {
	genLog(this.logger, iface.Fatal, this.pkg, this.marker, fmt.Sprint(args...)).Commit()
}

func (this *SugaredLogger) Fatalf(fmtstr string, args ...interface{}) {
	genLog(this.logger, iface.Fatal, this.pkg, this.marker, fmt.Sprintf(fmtstr, args...)).Commit()
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
