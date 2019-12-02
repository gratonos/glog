package logger

import (
	"os"
	"sync/atomic"
	"time"

	"github.com/gratonos/glog/pkg/glog"
)

type Logger struct {
	level glog.Level
}

func New(level glog.Level) *Logger {
	return &Logger{
		level: level,
	}
}

func (this *Logger) GenLog(level glog.Level) *Log {
	if this.getLevel() > level {
		return nil
	}

	log := getLog(this)
	log.reserveTimestamp()
	log.appendLevel(level)

	return log
}

func (this *Logger) getLevel() glog.Level {
	return glog.Level(atomic.LoadInt32((*int32)(&this.level)))
}

func (this *Logger) commit(log *Log) {
	tm := time.Now()
	log.fillTimestamp(tm)

	os.Stderr.Write(log.buf)

	putLog(log)
}
