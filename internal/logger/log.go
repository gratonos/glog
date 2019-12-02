package logger

import (
	"strconv"
	"sync"
	"time"

	"github.com/gratonos/glog/pkg/glog/logger/iface"
)

type Log struct {
	logger *Logger
	buf    []byte
}

var logPool = sync.Pool{
	New: func() interface{} {
		return &Log{
			buf: make([]byte, 0, logBufLen),
		}
	},
}

func getLog(logger *Logger) *Log {
	log := logPool.Get().(*Log)
	log.logger = logger
	log.buf = log.buf[:0]
	return log
}

func putLog(log *Log) {
	logPool.Put(log)
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
	this.buf = append(this.buf, timeHolder...)
}

func (this *Log) appendLevel(level iface.Level) {
	this.appendSeparator()
	this.buf = append(this.buf, levelNames[level]...)
}

func (this *Log) fillTimestamp(tm time.Time) {
	fillInt(this.buf[hourBegin:hourEnd], tm.Hour())
	fillInt(this.buf[minuteBegin:microEnd], tm.Minute())
	fillInt(this.buf[secondBegin:secondEnd], tm.Second())
	fillInt(this.buf[microBegin:microEnd], tm.Nanosecond()/1000)
}

func (this *Log) appendInt64Field(key string, value int64) *Log {
	if this != nil {
		this.appendKey(key)
		this.buf = strconv.AppendInt(this.buf, value, 10)
		this.appendEnd()
	}
	return this
}

func (this *Log) appendUint64Field(key string, value uint64) *Log {
	if this != nil {
		this.appendKey(key)
		this.buf = strconv.AppendUint(this.buf, value, 10)
		this.appendEnd()
	}
	return this
}

func (this *Log) appendKey(key string) {
	this.appendSeparator()
	this.buf = append(this.buf, '(')
	this.buf = append(this.buf, key...)
	this.buf = append(this.buf, ": "...)
}

func (this *Log) appendEnd() {
	this.buf = append(this.buf, ')')
}

func (this *Log) appendNewLine() {
	this.buf = append(this.buf, '\n')
}

func (this *Log) appendSeparator() {
	this.buf = append(this.buf, ' ')
}

// assert(n >= 0 && len(buf) >= digits(n) && len(buf) % 2 == 0)
func fillInt(buf []byte, n int) {
	const smallsString = "00010203040506070809" +
		"10111213141516171819" +
		"20212223242526272829" +
		"30313233343536373839" +
		"40414243444546474849" +
		"50515253545556575859" +
		"60616263646566676869" +
		"70717273747576777879" +
		"80818283848586878889" +
		"90919293949596979899"
	i := len(buf)
	for n > 0 {
		i -= 2
		j := (n % 100) << 1
		buf[i+1] = smallsString[j+1]
		buf[i+0] = smallsString[j+0]
		n /= 100
	}
}
