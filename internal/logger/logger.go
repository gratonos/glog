package logger

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

type Logger struct {
	level Level
}

func New(level Level) *Logger {
	return &Logger{
		level: level,
	}
}

func (this *Logger) Log(level Level, callDepth int) *Log {
	if level < Trace || level > Fatal {
		panic(fmt.Sprintf("glog: illegal log level: %d", level))
	}
	return this.genLog(level, callDepth)
}

func (this *Logger) Trace() *Log {
	return this.genLog(Trace, 0)
}

func (this *Logger) Debug() *Log {
	return this.genLog(Debug, 0)
}

func (this *Logger) Info() *Log {
	return this.genLog(Info, 0)
}

func (this *Logger) Warn() *Log {
	return this.genLog(Warn, 0)
}

func (this *Logger) Error() *Log {
	return this.genLog(Error, 0)
}

func (this *Logger) Fatal() *Log {
	return this.genLog(Fatal, 0)
}

func (this *Logger) genLog(level Level, callDepth int) *Log {
	if this.getLevel() > level {
		return nil
	}

	log := getLog(this)
	log.reserveTimestamp()
	log.appendLevel(level)

	return log
}

func (this *Logger) getLevel() Level {
	return Level(atomic.LoadInt32((*int32)(&this.level)))
}

func (this *Logger) commit(log *Log) {
	tm := time.Now()
	log.fillTimestamp(tm)

	os.Stderr.Write(log.buf)

	putLog(log)
}
