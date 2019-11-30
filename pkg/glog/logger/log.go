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
	return this.appendInt64Field(key, int64(value))
}

func (this *Log) Int8(key string, value int8) *Log {
	return this.appendInt64Field(key, int64(value))
}

func (this *Log) Int16(key string, value int16) *Log {
	return this.appendInt64Field(key, int64(value))
}

func (this *Log) Int32(key string, value int32) *Log {
	return this.appendInt64Field(key, int64(value))
}

func (this *Log) Int64(key string, value int64) *Log {
	return this.appendInt64Field(key, value)
}

func (this *Log) Uint(key string, value uint) *Log {
	return this.appendUint64Field(key, uint64(value))
}

func (this *Log) Uint8(key string, value uint8) *Log {
	return this.appendUint64Field(key, uint64(value))
}

func (this *Log) Uint16(key string, value uint16) *Log {
	return this.appendUint64Field(key, uint64(value))
}

func (this *Log) Uint32(key string, value uint32) *Log {
	return this.appendUint64Field(key, uint64(value))
}

func (this *Log) Uint64(key string, value uint64) *Log {
	return this.appendUint64Field(key, value)
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

func (this *Log) appendInt64Field(key string, value int64) *Log {
	return this.appendField(key, func() {
		this.buf = strconv.AppendInt(this.buf, value, 10)
	})
}

func (this *Log) appendUint64Field(key string, value uint64) *Log {
	return this.appendField(key, func() {
		this.buf = strconv.AppendUint(this.buf, value, 10)
	})
}

func (this *Log) appendField(key string, appender func()) *Log {
	if this == nil {
		return nil
	}

	this.appendSeparator()

	this.buf = append(this.buf, '(')
	this.buf = append(this.buf, key...)
	this.buf = append(this.buf, ": "...)
	appender()
	this.buf = append(this.buf, ')')

	return this
}

func (this *Log) appendNewLine() {
	this.buf = append(this.buf, '\n')
}

func (this *Log) appendSeparator() {
	this.buf = append(this.buf, ' ')
}

func (this *Log) put() {
}
