package logger

import (
	"os"
)

type Logger struct {
}

func New() *Logger {
	return &Logger{}
}

func (this *Logger) Trace() *Log {
	return getLog(this)
}

func (this *Logger) commit(log *Log) {
	log.fillTimestamp()
	os.Stderr.Write(log.buf)
}
