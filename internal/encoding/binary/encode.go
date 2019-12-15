package binary

import (
	"encoding/binary"
	"math"
	"time"
	"unsafe"
)

func AppendBinaryMeta(dst []byte) []byte {
	dst = append(dst, binaryMagic...)
	dst = appendUint8(dst, binaryVersion)
	return dst
}

func ResetBuf(dst []byte) []byte {
	return dst[:sizeOfHeader]
}

func AppendTime(dst []byte, tm time.Time) []byte {
	dst = appendFieldKind(dst, Timestamp)
	dst = appendUint64(dst, uint64(tm.UnixNano()))
	return dst
}

func AppendLevel(dst []byte, level uint8) []byte {
	dst = appendFieldKind(dst, Level)
	dst = appendUint8(dst, level)
	return dst
}

func AppendPkg(dst []byte, pkg string) []byte {
	dst = appendFieldKind(dst, Pkg)
	dst = appendShortString(dst, pkg)
	return dst
}

func AppendFunc(dst []byte, fn string) []byte {
	dst = appendFieldKind(dst, Func)
	dst = appendShortString(dst, fn)
	return dst
}

func AppendFile(dst []byte, file string) []byte {
	dst = appendFieldKind(dst, File)
	dst = appendShortString(dst, file)
	return dst
}

func AppendLine(dst []byte, line int) []byte {
	dst = appendFieldKind(dst, Line)
	dst = appendUint32(dst, uint32(line))
	return dst
}

func AppendMark(dst []byte) []byte {
	return appendFieldKind(dst, Mark)
}

func AppendMsg(dst []byte, msg string) []byte {
	dst = appendFieldKind(dst, Msg)
	dst = appendString(dst, msg)
	return dst
}

func AppendBoolKV(dst []byte, key string, value bool) []byte {
	dst = appendKVMeta(dst, key, Bool)
	dst = appendBool(dst, value)
	return dst
}

func AppendByteKV(dst []byte, key string, value byte) []byte {
	dst = appendKVMeta(dst, key, Byte)
	dst = appendUint8(dst, value)
	return dst
}

func AppendRuneKV(dst []byte, key string, value rune) []byte {
	dst = appendKVMeta(dst, key, Rune)
	dst = appendUint32(dst, uint32(value))
	return dst
}

func AppendInt8KV(dst []byte, key string, value int8) []byte {
	dst = appendKVMeta(dst, key, Int8)
	dst = appendUint8(dst, uint8(value))
	return dst
}

func AppendInt16KV(dst []byte, key string, value int16) []byte {
	dst = appendKVMeta(dst, key, Int16)
	dst = appendUint16(dst, uint16(value))
	return dst
}

func AppendInt32KV(dst []byte, key string, value int32) []byte {
	dst = appendKVMeta(dst, key, Int32)
	dst = appendUint32(dst, uint32(value))
	return dst
}

func AppendInt64KV(dst []byte, key string, value int64) []byte {
	dst = appendKVMeta(dst, key, Int64)
	dst = appendUint64(dst, uint64(value))
	return dst
}

func AppendUint8KV(dst []byte, key string, value uint8) []byte {
	dst = appendKVMeta(dst, key, Uint8)
	dst = appendUint8(dst, value)
	return dst
}

func AppendUint16KV(dst []byte, key string, value uint16) []byte {
	dst = appendKVMeta(dst, key, Uint16)
	dst = appendUint16(dst, value)
	return dst
}

func AppendUint32KV(dst []byte, key string, value uint32) []byte {
	dst = appendKVMeta(dst, key, Uint32)
	dst = appendUint32(dst, value)
	return dst
}

func AppendUint64KV(dst []byte, key string, value uint64) []byte {
	dst = appendKVMeta(dst, key, Uint64)
	dst = appendUint64(dst, value)
	return dst
}

func AppendUintptrKV(dst []byte, key string, value uintptr) []byte {
	dst = appendKVMeta(dst, key, Uintptr)
	dst = appendUint64(dst, uint64(value))
	return dst
}

func AppendFloat32KV(dst []byte, key string, value float32) []byte {
	dst = appendKVMeta(dst, key, Float32)
	dst = appendFloat32(dst, value)
	return dst
}

func AppendFloat64KV(dst []byte, key string, value float64) []byte {
	dst = appendKVMeta(dst, key, Float64)
	dst = appendFloat64(dst, value)
	return dst
}

func AppendComplex64KV(dst []byte, key string, value complex64) []byte {
	dst = appendKVMeta(dst, key, Complex64)
	dst = appendFloat32(dst, real(value))
	dst = appendFloat32(dst, imag(value))
	return dst
}

func AppendComplex128KV(dst []byte, key string, value complex128) []byte {
	dst = appendKVMeta(dst, key, Complex128)
	dst = appendFloat64(dst, real(value))
	dst = appendFloat64(dst, imag(value))
	return dst
}

func AppendStringKV(dst []byte, key, value string) []byte {
	dst = appendKVMeta(dst, key, String)
	dst = appendString(dst, value)
	return dst
}

func AppendTimeKV(dst []byte, key string, value time.Time) []byte {
	dst = appendKVMeta(dst, key, Time)
	dst = appendUint64(dst, uint64(value.UnixNano()))
	return dst
}

func AppendDurationKV(dst []byte, key string, value time.Duration) []byte {
	dst = appendKVMeta(dst, key, Duration)
	dst = appendUint64(dst, uint64(value))
	return dst
}

func AppendEnd(dst []byte) []byte {
	return appendFieldKind(dst, End)
}

func appendKVMeta(dst []byte, key string, kind ValueKind) []byte {
	dst = appendFieldKind(dst, KeyValue)
	dst = appendKey(dst, key)
	dst = appendValueKind(dst, kind)
	return dst
}

func appendFieldKind(dst []byte, kind FieldKind) []byte {
	return appendUint8(dst, uint8(kind))
}

func appendKey(dst []byte, key string) []byte {
	return appendShortString(dst, key)
}

func appendValueKind(dst []byte, kind ValueKind) []byte {
	return appendUint8(dst, uint8(kind))
}

func appendBool(dst []byte, b bool) []byte {
	var n uint8
	if b {
		n = 1
	}
	return appendUint8(dst, n)
}

func appendUint8(dst []byte, value uint8) []byte {
	return append(dst, value)
}

func appendUint16(dst []byte, value uint16) []byte {
	buf := make([]byte, unsafe.Sizeof(value))
	binary.LittleEndian.PutUint16(buf, value)
	return append(dst, buf...)
}

func appendUint32(dst []byte, value uint32) []byte {
	buf := make([]byte, unsafe.Sizeof(value))
	binary.LittleEndian.PutUint32(buf, value)
	return append(dst, buf...)
}

func appendUint64(dst []byte, value uint64) []byte {
	buf := make([]byte, unsafe.Sizeof(value))
	binary.LittleEndian.PutUint64(buf, value)
	return append(dst, buf...)
}

func appendFloat32(dst []byte, value float32) []byte {
	buf := make([]byte, unsafe.Sizeof(value))
	binary.LittleEndian.PutUint32(buf, math.Float32bits(value))
	return append(dst, buf...)
}

func appendFloat64(dst []byte, value float64) []byte {
	buf := make([]byte, unsafe.Sizeof(value))
	binary.LittleEndian.PutUint64(buf, math.Float64bits(value))
	return append(dst, buf...)
}

func appendString(dst []byte, str string) []byte {
	size := uint16(len(str))
	buf := make([]byte, unsafe.Sizeof(size))
	binary.LittleEndian.PutUint16(buf, size)
	dst = append(dst, buf...)
	dst = append(dst, str[:size]...)
	return dst
}

func appendShortString(dst []byte, str string) []byte {
	size := uint8(len(str))
	dst = append(dst, size)
	dst = append(dst, str[:size]...)
	return dst
}
