package logger

import (
	"strconv"
)

type Log struct {
	logger *Logger
	buf    []byte
}

func getLog(logger *Logger) *Log {
	return &Log{
		logger: logger,
	}
}

func (this *Log) Int(key string, value int) *Log {
	this.buf = append(this.buf, key...)
	this.buf = append(this.buf, ':')
	this.buf = strconv.AppendInt(this.buf, int64(value), 10)
	return this
}

func (this *Log) Commit() {
	this.buf = append(this.buf, '\n')
	this.logger.Commit(this)
}
