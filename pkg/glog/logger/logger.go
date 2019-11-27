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

func (this *Logger) Commit(log *Log) {
	os.Stderr.Write(log.buf)
}
