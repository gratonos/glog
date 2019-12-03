package logger

import (
	"os"
	"time"

	"github.com/gratonos/glog/pkg/glog/logger/iface"
)

type Logger struct {
	level atomicLevel
}

func New() *Logger {
	return &Logger{
		level: atomicLevel{
			level: iface.Trace,
		},
	}
}

func (this *Logger) GenLog(level iface.Level) *Log {
	if this.level.Get() > level {
		return nil
	}

	log := getLog(this)
	log.reserveTimestamp()
	log.appendLevel(level)

	return log
}

func (this *Logger) commit(log *Log) {
	tm := time.Now()
	log.fillTimestamp(tm)

	os.Stderr.Write(log.buf)

	putLog(log)
}
