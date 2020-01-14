package logger

import (
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/gratonos/glog/internal/encoding/binary"
	ilogger "github.com/gratonos/glog/internal/logger"
	"github.com/gratonos/glog/pkg/glog/iface"
)

type Log struct {
	logger *ilogger.Logger
	buf    []byte
}

const logBufLen = 1024

var logPool = sync.Pool{
	New: func() interface{} {
		buf := make([]byte, 0, logBufLen)
		buf = binary.AppendBinaryMeta(buf)
		return &Log{
			buf: buf,
		}
	},
}

func genLog(logger *ilogger.Logger, level iface.Level, pkg string, frameSkip int) *Log {
	if logger.Level() > level {
		return nil
	}

	log := logPool.Get().(*Log)
	log.reset(logger)
	log.appendPreInfo(level, pkg, frameSkip+1)
	return log
}

func (this *Log) Bool(key string, value bool) *Log {
	if this != nil {
		this.buf = binary.AppendBoolContext(this.buf, key, value)
	}
	return this
}

func (this *Log) Byte(key string, value byte) *Log {
	if this != nil {
		this.buf = binary.AppendByteContext(this.buf, key, value)
	}
	return this
}

func (this *Log) Rune(key string, value rune) *Log {
	if this != nil {
		this.buf = binary.AppendRuneContext(this.buf, key, value)
	}
	return this
}

func (this *Log) Int(key string, value int) *Log {
	return this.Int64(key, int64(value))
}

func (this *Log) Int8(key string, value int8) *Log {
	if this != nil {
		this.buf = binary.AppendInt8Context(this.buf, key, value)
	}
	return this
}

func (this *Log) Int16(key string, value int16) *Log {
	if this != nil {
		this.buf = binary.AppendInt16Context(this.buf, key, value)
	}
	return this
}

func (this *Log) Int32(key string, value int32) *Log {
	if this != nil {
		this.buf = binary.AppendInt32Context(this.buf, key, value)
	}
	return this
}

func (this *Log) Int64(key string, value int64) *Log {
	if this != nil {
		this.buf = binary.AppendInt64Context(this.buf, key, value)
	}
	return this
}

func (this *Log) Uint(key string, value uint) *Log {
	return this.Uint64(key, uint64(value))
}

func (this *Log) Uint8(key string, value uint8) *Log {
	if this != nil {
		this.buf = binary.AppendUint8Context(this.buf, key, value)
	}
	return this
}

func (this *Log) Uint16(key string, value uint16) *Log {
	if this != nil {
		this.buf = binary.AppendUint16Context(this.buf, key, value)
	}
	return this
}

func (this *Log) Uint32(key string, value uint32) *Log {
	if this != nil {
		this.buf = binary.AppendUint32Context(this.buf, key, value)
	}
	return this
}

func (this *Log) Uint64(key string, value uint64) *Log {
	if this != nil {
		this.buf = binary.AppendUint64Context(this.buf, key, value)
	}
	return this
}

func (this *Log) Uintptr(key string, value uintptr) *Log {
	if this != nil {
		this.buf = binary.AppendUintptrContext(this.buf, key, value)
	}
	return this
}

func (this *Log) Float32(key string, value float32) *Log {
	if this != nil {
		this.buf = binary.AppendFloat32Context(this.buf, key, value)
	}
	return this
}

func (this *Log) Float64(key string, value float64) *Log {
	if this != nil {
		this.buf = binary.AppendFloat64Context(this.buf, key, value)
	}
	return this
}

func (this *Log) Complex64(key string, value complex64) *Log {
	if this != nil {
		this.buf = binary.AppendComplex64Context(this.buf, key, value)
	}
	return this
}

func (this *Log) Complex128(key string, value complex128) *Log {
	if this != nil {
		this.buf = binary.AppendComplex128Context(this.buf, key, value)
	}
	return this
}

func (this *Log) Str(key, value string) *Log {
	if this != nil {
		this.buf = binary.AppendStringContext(this.buf, key, value)
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
		this.buf = binary.AppendTimeContext(this.buf, key, value)
	}
	return this
}

func (this *Log) Duration(key string, value time.Duration) *Log {
	if this != nil {
		this.buf = binary.AppendDurationContext(this.buf, key, value)
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

func (this *Log) reset(logger *ilogger.Logger) {
	this.logger = logger
	this.buf = binary.ResetBuf(this.buf)
}

func (this *Log) appendPreInfo(level iface.Level, pkg string, frameSkip int) {
	this.buf = binary.AppendLevel(this.buf, level)
	this.buf = binary.AppendPkg(this.buf, pkg)
	if this.logger.FileLine() {
		file, line := fileAndLine(frameSkip + 1)
		this.buf = binary.AppendFile(this.buf, file)
		this.buf = binary.AppendLine(this.buf, line)
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

func fileAndLine(frameSkip int) (string, int) {
	_, file, line, ok := runtime.Caller(frameSkip + 1)
	if ok {
		return filepath.Base(file), line
	} else {
		return "???", 0
	}
}
