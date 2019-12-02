package logger

import (
	"fmt"

	"github.com/gratonos/glog/internal/logger"
	"github.com/gratonos/glog/pkg/glog"
)

type Logger struct {
	log *logger.Logger
}

func New(level glog.Level) *Logger {
	return &Logger{
		log: logger.New(level),
	}
}

func (this *Logger) Log(level glog.Level) *logger.Log {
	if level < glog.Trace || level > glog.Fatal {
		panic(fmt.Sprintf("glog: illegal log level: %d", level))
	}
	return this.log.GenLog(level)
}

func (this *Logger) Trace() *logger.Log {
	return this.log.GenLog(glog.Trace)
}

func (this *Logger) Debug() *logger.Log {
	return this.log.GenLog(glog.Debug)
}

func (this *Logger) Info() *logger.Log {
	return this.log.GenLog(glog.Info)
}

func (this *Logger) Warn() *logger.Log {
	return this.log.GenLog(glog.Warn)
}

func (this *Logger) Error() *logger.Log {
	return this.log.GenLog(glog.Error)
}

func (this *Logger) Fatal() *logger.Log {
	return this.log.GenLog(glog.Fatal)
}
