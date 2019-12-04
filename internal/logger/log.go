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

func (this *Log) Bool(key string, value bool) *Log {
	if value {
		return this.appendStrField(key, "true")
	} else {
		return this.appendStrField(key, "false")
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

func (this *Log) Float32(key string, value float32) *Log {
	return this.appendFloatField(key, float64(value), 32)
}

func (this *Log) Float64(key string, value float64) *Log {
	return this.appendFloatField(key, value, 64)
}

func (this *Log) Complex64(key string, value complex64) *Log {
	return this.appendComplexField(key, complex128(value), 32)
}

func (this *Log) Complex128(key string, value complex128) *Log {
	return this.appendComplexField(key, value, 64)
}

func (this *Log) Byte(key string, value byte) *Log {
	return this.appendByteField(key, value)
}

func (this *Log) Rune(key string, value rune) *Log {
	return this.appendStrField(key, string(value))
}

func (this *Log) Uintptr(key string, value uintptr) *Log {
	return this.appendUintptrField(key, value)
}

func (this *Log) Str(key, value string) *Log {
	return this.appendStrField(key, value)
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

func (this *Log) appendInfo(info string) {
	this.appendSeparator()
	this.buf = append(this.buf, info...)
}

func (this *Log) appendMsg(msg string) {
	this.appendSeparator()
	this.appendMsgLeftBound()
	this.buf = append(this.buf, msg...)
	this.appendMsgRightBound()
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
		this.buf = strconv.AppendInt(this.buf, value, integerBase)
		this.appendEnd()
	}
	return this
}

func (this *Log) appendUint64Field(key string, value uint64) *Log {
	if this != nil {
		this.appendKey(key)
		this.buf = strconv.AppendUint(this.buf, value, integerBase)
		this.appendEnd()
	}
	return this
}

func (this *Log) appendFloatField(key string, value float64, bitSize int) *Log {
	if this != nil {
		this.appendKey(key)
		this.appendFloat(value, bitSize)
		this.appendEnd()
	}
	return this
}

func (this *Log) appendComplexField(key string, value complex128, bitSize int) *Log {
	if this != nil {
		this.appendKey(key)
		this.appendFloat(real(value), bitSize)
		this.buf = append(this.buf, '+')
		this.appendFloat(imag(value), bitSize)
		this.buf = append(this.buf, 'i')
		this.appendEnd()
	}
	return this
}

func (this *Log) appendByteField(key string, value byte) *Log {
	if this != nil {
		this.appendKey(key)
		this.appendByte(value)
		this.appendEnd()
	}
	return this
}

func (this *Log) appendUintptrField(key string, value uintptr) *Log {
	if this != nil {
		this.appendKey(key)
		this.buf = append(this.buf, "0x"...)
		this.appendUintptr(value)
		this.appendEnd()
	}
	return this
}

func (this *Log) appendStrField(key, value string) *Log {
	if this != nil {
		this.appendKey(key)
		this.buf = append(this.buf, value...)
		this.appendEnd()
	}
	return this
}

func (this *Log) appendFloat(value float64, bitSize int) {
	this.buf = strconv.AppendFloat(this.buf, value, floatFormat, floatPrecision, bitSize)
}

func (this *Log) appendByte(value byte) {
	this.buf = append(this.buf, byteReps[value]...)
}

func (this *Log) appendUintptr(value uintptr) {
	if value == 0 {
		this.buf = append(this.buf, "00"...)
		return
	}

	const bufLen = 16
	var buf [bufLen]byte
	n := 0
	for value != 0 {
		rep := byteReps[value&0xff]
		buf[bufLen-n-1] = rep[1]
		buf[bufLen-n-2] = rep[0]
		value >>= 8
		n += 2
	}
	this.buf = append(this.buf, buf[bufLen-n:]...)
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

func (this *Log) appendMsgLeftBound() {
	this.buf = append(this.buf, '<')
}

func (this *Log) appendMsgRightBound() {
	this.buf = append(this.buf, '>')
}

// assert(n >= 0 && len(buf) >= digits(n) && len(buf) % 2 == 0)
func fillInt(buf []byte, n int) {
	i := len(buf)
	for n > 0 {
		i -= 2
		j := n % 100
		buf[i+1] = smallIntReps[j][1]
		buf[i+0] = smallIntReps[j][0]
		n /= 100
	}
}
