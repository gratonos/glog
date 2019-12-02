package logger

import (
	"os"
	"sync/atomic"
	"time"

	"github.com/gratonos/glog/pkg/glog/logger/iface"
)

type Logger struct {
	level iface.Level
}

func New() *Logger {
	return &Logger{
		level: iface.Trace,
	}
}

func (this *Logger) GenLog(level iface.Level) *Log {
	if this.getLevel() > level {
		return nil
	}

	log := getLog(this)
	log.reserveTimestamp()
	log.appendLevel(level)

	return log
}

func (this *Logger) getLevel() iface.Level {
	return iface.Level(atomic.LoadInt32((*int32)(&this.level)))
}

func (this *Logger) commit(log *Log) {
	tm := time.Now()
	log.fillTimestamp(tm)

	os.Stderr.Write(log.buf)

	putLog(log)
}
