package logger

import (
	"strconv"
	"time"
)

type Log struct {
	logger *Logger
	buf    []byte
}

func getLog(logger *Logger) *Log {
	return &Log{
		logger: logger,
		buf:    make([]byte, 0, logBufLen),
	}
}

func (this *Log) Int(key string, value int) *Log {
	if this == nil {
		return nil
	}

	this.appendSeparator()

	this.buf = append(this.buf, key...)
	this.buf = append(this.buf, ':')
	this.buf = strconv.AppendInt(this.buf, int64(value), 10)

	return this
}

func (this *Log) Commit() {
	if this == nil {
		return
	}

	this.appendNewLine()
	this.logger.commit(this)
}

func (this *Log) reserveTimestamp() {
	this.buf = this.buf[:timeLen]
}

func (this *Log) appendLevel(level Level) {
	this.appendSeparator()
	this.buf = append(this.buf, levelNames[level]...)
}

func (this *Log) appendRuntimeInfo(file string, line int, fn string) {
	this.appendSeparator()

	this.buf = append(this.buf, file...)
	this.buf = append(this.buf, ':')
	this.buf = strconv.AppendInt(this.buf, int64(line), 10)

	this.appendSeparator()

	this.buf = append(this.buf, fn...)
}

func (this *Log) fillTimestamp(tm time.Time) {
	tm.AppendFormat(this.buf[:0], timeLayout)
}

func (this *Log) appendNewLine() {
	this.buf = append(this.buf, '\n')
}

func (this *Log) appendSeparator() {
	this.buf = append(this.buf, ' ')
}

func (this *Log) put() {
}
