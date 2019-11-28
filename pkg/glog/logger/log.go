package logger

import (
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

type Log struct {
	logger *Logger
	buf    []byte
}

func getLog(logger *Logger) *Log {
	log := &Log{
		logger: logger,
		buf:    make([]byte, timeLen, logBufLen),
	}
	log.appendRuntimeInfo()
	return log
}

func (this *Log) Int(key string, value int) *Log {
	this.appendSeparator()

	this.buf = append(this.buf, key...)
	this.buf = append(this.buf, ':')
	this.buf = strconv.AppendInt(this.buf, int64(value), 10)

	return this
}

func (this *Log) Commit() {
	this.appendNewLine()
	this.logger.commit(this)
}

func (this *Log) appendRuntimeInfo() {
	this.appendSeparator()

	file, line, fn := runtimeInfo(stackOffset)
	this.buf = append(this.buf, file...)
	this.buf = append(this.buf, ':')
	this.buf = strconv.AppendInt(this.buf, int64(line), 10)

	this.appendSeparator()

	this.buf = append(this.buf, fn...)
}

func (this *Log) fillTimestamp() {
	time.Now().AppendFormat(this.buf[:0], timeLayout)
}

func (this *Log) appendNewLine() {
	this.buf = append(this.buf, '\n')
}

func (this *Log) appendSeparator() {
	this.buf = append(this.buf, ' ')
}

func runtimeInfo(callDepth int) (file string, line int, fn string) {
	var pc uintptr
	var ok bool
	pc, file, line, ok = runtime.Caller(callDepth)
	if ok {
		fn = runtime.FuncForPC(pc).Name()
	} else {
		file = "???"
		line = 0
		fn = "???"
	}
	return filepath.Base(file), line, fn
}
