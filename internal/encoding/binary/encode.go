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
	dst = appendUint32(dst, 0) // reserve for load len
	return dst
}

func AppendTimestamp(dst []byte, tm time.Time) []byte {
	dst = appendFieldKind(dst, timestampField)
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
	dst = appendString(dst, pkg)
	return dst
}

func AppendFileName(dst []byte, name string) []byte {
	dst = appendFieldKind(dst, fileNameField)
	dst = appendString(dst, name)
	return dst
}

func AppendFileLine(dst []byte, line int) []byte {
	dst = appendFieldKind(dst, fileLineField)
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

func AppendInt8KV(dst []byte, key string, value int8, format string) []byte {
	dst = appendNumberKVMeta(dst, key, intValue, unsafe.Sizeof(value), format)
	dst = appendUint8(dst, uint8(value))
	return dst
}

func AppendInt16KV(dst []byte, key string, value int16, format string) []byte {
	dst = appendNumberKVMeta(dst, key, intValue, unsafe.Sizeof(value), format)
	dst = appendUint16(dst, uint16(value))
	return dst
}

func AppendInt32KV(dst []byte, key string, value int32, format string) []byte {
	dst = appendNumberKVMeta(dst, key, intValue, unsafe.Sizeof(value), format)
	dst = appendUint32(dst, uint32(value))
	return dst
}

func AppendInt64KV(dst []byte, key string, value int64, format string) []byte {
	dst = appendNumberKVMeta(dst, key, intValue, unsafe.Sizeof(value), format)
	dst = appendUint64(dst, uint64(value))
	return dst
}

func AppendUint8KV(dst []byte, key string, value uint8, format string) []byte {
	dst = appendNumberKVMeta(dst, key, uintValue, unsafe.Sizeof(value), format)
	dst = appendUint8(dst, value)
	return dst
}

func AppendUint16KV(dst []byte, key string, value uint16, format string) []byte {
	dst = appendNumberKVMeta(dst, key, uintValue, unsafe.Sizeof(value), format)
	dst = appendUint16(dst, value)
	return dst
}

func AppendUint32KV(dst []byte, key string, value uint32, format string) []byte {
	dst = appendNumberKVMeta(dst, key, uintValue, unsafe.Sizeof(value), format)
	dst = appendUint32(dst, value)
	return dst
}

func AppendUint64KV(dst []byte, key string, value uint64, format string) []byte {
	dst = appendNumberKVMeta(dst, key, uintValue, unsafe.Sizeof(value), format)
	dst = appendUint64(dst, value)
	return dst
}

func AppendUintptrKV(dst []byte, key string, value uintptr, format string) []byte {
	dst = appendNumberKVMeta(dst, key, uintptrValue, unsafe.Sizeof(uint64(value)), format)
	dst = appendUint64(dst, uint64(value))
	return dst
}

func AppendFloat32KV(dst []byte, key string, value float32, format string) []byte {
	dst = appendNumberKVMeta(dst, key, floatValue, unsafe.Sizeof(value), format)
	dst = appendFloat32(dst, value)
	return dst
}

func AppendFloat64KV(dst []byte, key string, value float64, format string) []byte {
	dst = appendNumberKVMeta(dst, key, floatValue, unsafe.Sizeof(value), format)
	dst = appendFloat64(dst, value)
	return dst
}

func AppendComplex64KV(dst []byte, key string, value complex64, format string) []byte {
	dst = appendNumberKVMeta(dst, key, complexValue, unsafe.Sizeof(value), format)
	dst = appendFloat32(dst, real(value))
	dst = appendFloat32(dst, imag(value))
	return dst
}

func AppendComplex128KV(dst []byte, key string, value complex128, format string) []byte {
	dst = appendNumberKVMeta(dst, key, complexValue, unsafe.Sizeof(value), format)
	dst = appendFloat64(dst, real(value))
	dst = appendFloat64(dst, imag(value))
	return dst
}

func AppendStringKV(dst []byte, key, value, format string) []byte {
	dst = appendKVMeta(dst, key, stringValue)
	dst = appendString(dst, format)
	dst = appendString(dst, value)
	return dst
}

func AppendTimeKV(dst []byte, key string, value time.Time, layout string) []byte {
	dst = appendKVMeta(dst, key, timeValue)
	dst = appendString(dst, layout)
	dst = appendUint64(dst, uint64(value.UnixNano()))
	return dst
}

func AppendDurationKV(dst []byte, key string, value time.Duration) []byte {
	dst = appendKVMeta(dst, key, durationValue)
	dst = appendUint64(dst, uint64(value))
	return dst
}

func AppendEnd(dst []byte) []byte {
	dst = appendFieldKind(dst, endIdentity)

	loadLen := len(dst) - sizeOfHeader
	if loadLen > math.MaxUint32 {
		panic(fmt.Sprintf("glog: log buffer is too large, beyond %d bytes", math.MaxUint32))
	}
	appendUint32(dst[sizeOfVersion:sizeOfVersion], uint32(loadLen))

	return dst
}

func appendNumberKVMeta(dst []byte, key string, kind valueKind,
	size uintptr, format string) []byte {

	dst = appendKVMeta(dst, key, kind)
	dst = appendString(dst, format)
	dst = appendNumberSize(dst, size)
	return dst
}

func appendKVMeta(dst []byte, key string, kind valueKind) []byte {
	dst = appendFieldKind(dst, keyValuePairField)
	dst = appendKey(dst, key)
	dst = appendValueKind(dst, kind)
	return dst
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
	buf := make([]byte, unsafe.Sizeof(uint(0)))
	size := putUvarint(buf, uint(len(str)))
	dst = append(dst, byte(size))
	dst = append(dst, buf[:size]...)
	dst = append(dst, str...)
	return dst
}

func appendFieldKind(dst []byte, kind fieldKind) []byte {
	return append(dst, byte(kind))
}

func appendValueKind(dst []byte, kind valueKind) []byte {
	return append(dst, byte(kind))
}

func appendKey(dst []byte, key string) []byte {
	return appendString(dst, key)
}

func appendNumberSize(dst []byte, size uintptr) []byte {
	return append(dst, byte(size))
}

func putUvarint(buf []byte, n uint) int {
	i := 0
	for n > 0 {
		buf[i] = byte(n)
		n >>= 8
		i++
	}
	return i
}
