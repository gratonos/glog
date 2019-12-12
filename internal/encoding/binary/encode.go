package binary

import (
	"encoding/binary"
	"fmt"
	"math"
	"time"
	"unsafe"
)

func AppendBegin(dst []byte) []byte {
	dst = appendUint8(dst, binaryVersion)
	dst = appendUint32(dst, 0) // reserved for size of load
	return dst
}

func AppendTime(dst []byte, tm time.Time) []byte {
	dst = appendFieldKind(dst, timeField)
	dst = appendUint64(dst, uint64(tm.UnixNano()))
	return dst
}

func AppendLevel(dst []byte, level uint8) []byte {
	dst = appendFieldKind(dst, levelField)
	dst = appendUint8(dst, level)
	return dst
}

func AppendPkg(dst []byte, pkg string) []byte {
	dst = appendFieldKind(dst, pkgField)
	dst = appendShortString(dst, pkg)
	return dst
}

func AppendFunc(dst []byte, fn string) []byte {
	dst = appendFieldKind(dst, funcField)
	dst = appendShortString(dst, fn)
	return dst
}

func AppendFile(dst []byte, file string) []byte {
	dst = appendFieldKind(dst, fileField)
	dst = appendShortString(dst, file)
	return dst
}

func AppendLine(dst []byte, line int) []byte {
	dst = appendFieldKind(dst, lineField)
	dst = appendUint32(dst, uint32(line))
	return dst
}

func AppendMark(dst []byte, mark bool) []byte {
	dst = appendFieldKind(dst, markField)
	dst = appendBool(dst, mark)
	return dst
}

func AppendMsg(dst []byte, msg string) []byte {
	dst = appendFieldKind(dst, msgField)
	dst = appendString(dst, msg)
	return dst
}

func AppendBoolKV(dst []byte, key string, value bool) []byte {
	dst = appendKVMeta(dst, key, boolValue)
	dst = appendBool(dst, value)
	return dst
}

func AppendByteKV(dst []byte, key string, value byte) []byte {
	dst = appendKVMeta(dst, key, byteValue)
	dst = appendUint8(dst, value)
	return dst
}

func AppendRuneKV(dst []byte, key string, value rune) []byte {
	dst = appendKVMeta(dst, key, runeValue)
	dst = appendUint32(dst, uint32(value))
	return dst
}

func AppendInt8KV(dst []byte, key string, value int8) []byte {
	dst = appendKVMeta(dst, key, int8Value)
	dst = appendUint8(dst, uint8(value))
	return dst
}

func AppendInt16KV(dst []byte, key string, value int16) []byte {
	dst = appendKVMeta(dst, key, int16Value)
	dst = appendUint16(dst, uint16(value))
	return dst
}

func AppendInt32KV(dst []byte, key string, value int32) []byte {
	dst = appendKVMeta(dst, key, int32Value)
	dst = appendUint32(dst, uint32(value))
	return dst
}

func AppendInt64KV(dst []byte, key string, value int64) []byte {
	dst = appendKVMeta(dst, key, int64Value)
	dst = appendUint64(dst, uint64(value))
	return dst
}

func AppendUint8KV(dst []byte, key string, value uint8) []byte {
	dst = appendKVMeta(dst, key, uint8Value)
	dst = appendUint8(dst, value)
	return dst
}

func AppendUint16KV(dst []byte, key string, value uint16) []byte {
	dst = appendKVMeta(dst, key, uint16Value)
	dst = appendUint16(dst, value)
	return dst
}

func AppendUint32KV(dst []byte, key string, value uint32) []byte {
	dst = appendKVMeta(dst, key, uint32Value)
	dst = appendUint32(dst, value)
	return dst
}

func AppendUint64KV(dst []byte, key string, value uint64) []byte {
	dst = appendKVMeta(dst, key, uint64Value)
	dst = appendUint64(dst, value)
	return dst
}

func AppendUintptrKV(dst []byte, key string, value uintptr) []byte {
	dst = appendKVMeta(dst, key, uintptrValue)
	dst = appendUint64(dst, uint64(value))
	return dst
}

func AppendFloat32KV(dst []byte, key string, value float32) []byte {
	dst = appendKVMeta(dst, key, float32Value)
	dst = appendFloat32(dst, value)
	return dst
}

func AppendFloat64KV(dst []byte, key string, value float64) []byte {
	dst = appendKVMeta(dst, key, float64Value)
	dst = appendFloat64(dst, value)
	return dst
}

func AppendComplex64KV(dst []byte, key string, value complex64) []byte {
	dst = appendKVMeta(dst, key, complex64Value)
	dst = appendFloat32(dst, real(value))
	dst = appendFloat32(dst, imag(value))
	return dst
}

func AppendComplex128KV(dst []byte, key string, value complex128) []byte {
	dst = appendKVMeta(dst, key, complex128Value)
	dst = appendFloat64(dst, real(value))
	dst = appendFloat64(dst, imag(value))
	return dst
}

func AppendStringKV(dst []byte, key, value string) []byte {
	dst = appendKVMeta(dst, key, stringValue)
	dst = appendString(dst, value)
	return dst
}

func AppendTimeKV(dst []byte, key string, value time.Time) []byte {
	dst = appendKVMeta(dst, key, timeValue)
	dst = appendUint64(dst, uint64(value.UnixNano()))
	return dst
}

func AppendDurationKV(dst []byte, key string, value time.Duration) []byte {
	dst = appendKVMeta(dst, key, durationValue)
	dst = appendUint64(dst, uint64(value))
	return dst
}

func AppendEnd(dst []byte) []byte {
	dst = appendFieldKind(dst, endIdentifier)

	loadLen := len(dst) - sizeOfHeader
	if loadLen > math.MaxUint32 {
		panic(fmt.Sprintf("glog: log buffer is too large, beyond %d bytes", math.MaxUint32))
	}
	appendUint32(dst[sizeOfVersion:sizeOfVersion], uint32(loadLen))

	return dst
}

func appendKVMeta(dst []byte, key string, kind valueKind) []byte {
	dst = appendFieldKind(dst, keyValueField)
	dst = appendKey(dst, key)
	dst = appendValueKind(dst, kind)
	return dst
}

func appendFieldKind(dst []byte, kind fieldKind) []byte {
	return append(dst, byte(kind))
}

func appendKey(dst []byte, key string) []byte {
	return appendShortString(dst, key)
}

func appendValueKind(dst []byte, kind valueKind) []byte {
	return append(dst, byte(kind))
}

func appendBool(dst []byte, b bool) []byte {
	if b {
		return append(dst, 1)
	} else {
		return append(dst, 0)
	}
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
