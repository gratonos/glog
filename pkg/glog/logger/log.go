package logger

import (
	"sync"
	"time"

	"github.com/gratonos/glog/internal/encoding/binary"
	ilogger "github.com/gratonos/glog/internal/logger"
)

type preInfo struct {
	Pkg   string
	Func  string
	File  string
	Line  int
	Level uint8
}

type Log struct {
	logger *ilogger.Logger
	buf    []byte
}

const logBufLen = 512

var logPool = sync.Pool{
	New: func() interface{} {
		buf := make([]byte, 0, logBufLen)
		buf = binary.AppendBinaryMeta(buf)
		return &Log{
			buf: buf,
		}
	},
}

func genLog(logger *ilogger.Logger, info *preInfo) *Log {
	log := getLog(logger)
	log.appendPreInfo(info)
	return log
}

func getLog(logger *ilogger.Logger) *Log {
	log := logPool.Get().(*Log)
	log.logger = logger
	log.buf = binary.ResetBuf(log.buf)
	return log
}

func (this *Log) Bool(key string, value bool) *Log {
	if this != nil {
		this.buf = binary.AppendBoolKV(this.buf, key, value)
	}
	return this
}

func (this *Log) Byte(key string, value byte) *Log {
	if this != nil {
		this.buf = binary.AppendByteKV(this.buf, key, value)
	}
	return this
}

func (this *Log) Rune(key string, value rune) *Log {
	if this != nil {
		this.buf = binary.AppendRuneKV(this.buf, key, value)
	}
	return this
}

func (this *Log) Int(key string, value int) *Log {
	return this.Int64(key, int64(value))
}

func (this *Log) Int8(key string, value int8) *Log {
	if this != nil {
		this.buf = binary.AppendInt8KV(this.buf, key, value)
	}
	return this
}

func (this *Log) Int16(key string, value int16) *Log {
	if this != nil {
		this.buf = binary.AppendInt16KV(this.buf, key, value)
	}
	return this
}

func (this *Log) Int32(key string, value int32) *Log {
	if this != nil {
		this.buf = binary.AppendInt32KV(this.buf, key, value)
	}
	return this
}

func (this *Log) Int64(key string, value int64) *Log {
	if this != nil {
		this.buf = binary.AppendInt64KV(this.buf, key, value)
	}
	return this
}

func (this *Log) Uint(key string, value uint) *Log {
	return this.Uint64(key, uint64(value))
}

func (this *Log) Uint8(key string, value uint8) *Log {
	if this != nil {
		this.buf = binary.AppendUint8KV(this.buf, key, value)
	}
	return this
}

func (this *Log) Uint16(key string, value uint16) *Log {
	if this != nil {
		this.buf = binary.AppendUint16KV(this.buf, key, value)
	}
	return this
}

func (this *Log) Uint32(key string, value uint32) *Log {
	if this != nil {
		this.buf = binary.AppendUint32KV(this.buf, key, value)
	}
	return this
}

func (this *Log) Uint64(key string, value uint64) *Log {
	if this != nil {
		this.buf = binary.AppendUint64KV(this.buf, key, value)
	}
	return this
}

func (this *Log) Uintptr(key string, value uintptr) *Log {
	if this != nil {
		this.buf = binary.AppendUintptrKV(this.buf, key, value)
	}
	return this
}

func (this *Log) Float32(key string, value float32) *Log {
	if this != nil {
		this.buf = binary.AppendFloat32KV(this.buf, key, value)
	}
	return this
}

func (this *Log) Float64(key string, value float64) *Log {
	if this != nil {
		this.buf = binary.AppendFloat64KV(this.buf, key, value)
	}
	return this
}

func (this *Log) Complex64(key string, value complex64) *Log {
	if this != nil {
		this.buf = binary.AppendComplex64KV(this.buf, key, value)
	}
	return this
}

func (this *Log) Complex128(key string, value complex128) *Log {
	if this != nil {
		this.buf = binary.AppendComplex128KV(this.buf, key, value)
	}
	return this
}

func (this *Log) Str(key, value string) *Log {
	if this != nil {
		this.buf = binary.AppendStringKV(this.buf, key, value)
	}
	return this
}

func (this *Log) Func(fn string) *Log {
	return this.Str("func", fn)
}

func (this *Log) Err(err error) *Log {
	return this.Str("error", err.Error())
}

func (this *Log) Time(key string, value time.Time) *Log {
	if this != nil {
		this.buf = binary.AppendTimeKV(this.buf, key, value)
	}
	return this
}

func (this *Log) Duration(key string, value time.Duration) *Log {
	if this != nil {
		this.buf = binary.AppendDurationKV(this.buf, key, value)
	}
	return this
}

func (this *Log) Mark() *Log {
	if this != nil {
		this.buf = binary.AppendMark(this.buf)
	}
	return this
}

func (this *Log) Commit(msg string) {
	if this != nil {
		this.buf = binary.AppendMsg(this.buf, msg)
		this.logger.Commit(this.emit, this.put)
	}
}

func (this *Log) appendPreInfo(info *preInfo) {
	this.buf = binary.AppendLevel(this.buf, info.Level)
	this.buf = binary.AppendPkg(this.buf, info.Pkg)
	if len(info.Func) != 0 {
		this.buf = binary.AppendFunc(this.buf, info.Func)
	}
	if len(info.File) != 0 {
		this.buf = binary.AppendFile(this.buf, info.File)
	}
	if info.Line != 0 {
		this.buf = binary.AppendLine(this.buf, info.Line)
	}
}

func (this *Log) emit(tm time.Time) []byte {
	this.buf = binary.AppendTime(this.buf, tm)
	this.buf = binary.AppendEnd(this.buf)
	return this.buf
}

func (this *Log) put() {
	logPool.Put(this)
}
