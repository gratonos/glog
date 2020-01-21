package binary

import (
	"fmt"
	"io"
	"time"
)

type ValueKind uint8

const (
	Bool ValueKind = iota
	Byte
	Rune
	Int8
	Int16
	Int32
	Int64
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	String
	Time
	Duration

	valueKindBound
)

func (self ValueKind) Legal() bool {
	return self < valueKindBound
}

var valueReaders = [...]func(io.Reader) (interface{}, error){
	Bool:       readBoolValue,
	Byte:       readByteValue,
	Rune:       readRuneValue,
	Int8:       readInt8Value,
	Int16:      readInt16Value,
	Int32:      readInt32Value,
	Int64:      readInt64Value,
	Uint8:      readUint8Value,
	Uint16:     readUint16Value,
	Uint32:     readUint32Value,
	Uint64:     readUint64Value,
	Uintptr:    readUintptrValue,
	Float32:    readFloat32Value,
	Float64:    readFloat64Value,
	Complex64:  readComplex64Value,
	Complex128: readComplex128Value,
	String:     readStringValue,
	Time:       readTimeValue,
	Duration:   readDurationValue,
}

type Context struct {
	Key    string
	Value  interface{}
	Format string
	Kind   ValueKind
}

func AppendBoolContext(dst []byte, key string, value bool) []byte {
	dst = appendContextMeta(dst, key, Bool)
	dst = appendBool(dst, value)
	return dst
}

func AppendByteContext(dst []byte, key string, value byte) []byte {
	dst = appendContextMeta(dst, key, Byte)
	dst = appendUint8(dst, value)
	return dst
}

func AppendRuneContext(dst []byte, key string, value rune) []byte {
	dst = appendContextMeta(dst, key, Rune)
	dst = appendUint32(dst, uint32(value))
	return dst
}

func AppendInt8Context(dst []byte, key string, value int8) []byte {
	dst = appendContextMeta(dst, key, Int8)
	dst = appendUint8(dst, uint8(value))
	return dst
}

func AppendInt16Context(dst []byte, key string, value int16) []byte {
	dst = appendContextMeta(dst, key, Int16)
	dst = appendUint16(dst, uint16(value))
	return dst
}

func AppendInt32Context(dst []byte, key string, value int32) []byte {
	dst = appendContextMeta(dst, key, Int32)
	dst = appendUint32(dst, uint32(value))
	return dst
}

func AppendInt64Context(dst []byte, key string, value int64) []byte {
	dst = appendContextMeta(dst, key, Int64)
	dst = appendUint64(dst, uint64(value))
	return dst
}

func AppendUint8Context(dst []byte, key string, value uint8) []byte {
	dst = appendContextMeta(dst, key, Uint8)
	dst = appendUint8(dst, value)
	return dst
}

func AppendUint16Context(dst []byte, key string, value uint16) []byte {
	dst = appendContextMeta(dst, key, Uint16)
	dst = appendUint16(dst, value)
	return dst
}

func AppendUint32Context(dst []byte, key string, value uint32) []byte {
	dst = appendContextMeta(dst, key, Uint32)
	dst = appendUint32(dst, value)
	return dst
}

func AppendUint64Context(dst []byte, key string, value uint64) []byte {
	dst = appendContextMeta(dst, key, Uint64)
	dst = appendUint64(dst, value)
	return dst
}

func AppendUintptrContext(dst []byte, key string, value uintptr) []byte {
	dst = appendContextMeta(dst, key, Uintptr)
	dst = appendUint64(dst, uint64(value))
	return dst
}

func AppendFloat32Context(dst []byte, key string, value float32) []byte {
	dst = appendContextMeta(dst, key, Float32)
	dst = appendFloat32(dst, value)
	return dst
}

func AppendFloat64Context(dst []byte, key string, value float64) []byte {
	dst = appendContextMeta(dst, key, Float64)
	dst = appendFloat64(dst, value)
	return dst
}

func AppendComplex64Context(dst []byte, key string, value complex64) []byte {
	dst = appendContextMeta(dst, key, Complex64)
	dst = appendComplex64(dst, value)
	return dst
}

func AppendComplex128Context(dst []byte, key string, value complex128) []byte {
	dst = appendContextMeta(dst, key, Complex128)
	dst = appendComplex128(dst, value)
	return dst
}

func AppendStringContext(dst []byte, key, value string) []byte {
	dst = appendContextMeta(dst, key, String)
	dst = appendString(dst, value)
	return dst
}

func AppendTimeContext(dst []byte, key string, value time.Time) []byte {
	dst = appendContextMeta(dst, key, Time)
	dst = appendUint64(dst, uint64(value.UnixNano()))
	return dst
}

func AppendDurationContext(dst []byte, key string, value time.Duration) []byte {
	dst = appendContextMeta(dst, key, Duration)
	dst = appendUint64(dst, uint64(value))
	return dst
}

func appendContextMeta(dst []byte, key string, kind ValueKind) []byte {
	dst = appendFieldKind(dst, fieldContext)
	dst = appendKey(dst, key)
	dst = appendValueKind(dst, kind)
	return dst
}

func appendKey(dst []byte, key string) []byte {
	return appendShortString(dst, key)
}

func appendValueKind(dst []byte, kind ValueKind) []byte {
	return appendUint8(dst, uint8(kind))
}

func readContextMeta(reader io.Reader) (string, ValueKind, error) {
	key, err := readKey(reader)
	if err != nil {
		return "", valueKindBound, err
	}
	kind, err := readValueKind(reader)
	if err != nil {
		return "", valueKindBound, err
	}
	return key, kind, nil
}

func readKey(reader io.Reader) (string, error) {
	return readShortString(reader)
}

func readValueKind(reader io.Reader) (ValueKind, error) {
	u, err := readUint8(reader)
	if err != nil {
		return valueKindBound, err
	}
	kind := ValueKind(u)
	if !kind.Legal() {
		return valueKindBound, newFormatError(fmt.Sprintf("illegal value kind %d", kind))
	}
	return kind, nil
}

func readBoolValue(reader io.Reader) (interface{}, error) {
	return readBool(reader)
}

func readByteValue(reader io.Reader) (interface{}, error) {
	return readUint8(reader)
}

func readRuneValue(reader io.Reader) (interface{}, error) {
	u, err := readUint32(reader)
	return rune(u), err
}

func readInt8Value(reader io.Reader) (interface{}, error) {
	u, err := readUint8(reader)
	return int8(u), err
}

func readInt16Value(reader io.Reader) (interface{}, error) {
	u, err := readUint16(reader)
	return int16(u), err
}

func readInt32Value(reader io.Reader) (interface{}, error) {
	u, err := readUint32(reader)
	return int32(u), err
}

func readInt64Value(reader io.Reader) (interface{}, error) {
	u, err := readUint64(reader)
	return int64(u), err
}

func readUint8Value(reader io.Reader) (interface{}, error) {
	return readUint8(reader)
}

func readUint16Value(reader io.Reader) (interface{}, error) {
	return readUint16(reader)
}

func readUint32Value(reader io.Reader) (interface{}, error) {
	return readUint32(reader)
}

func readUint64Value(reader io.Reader) (interface{}, error) {
	return readUint64(reader)
}

func readUintptrValue(reader io.Reader) (interface{}, error) {
	u, err := readUint64(reader)
	return uintptr(u), err
}

func readFloat32Value(reader io.Reader) (interface{}, error) {
	return readFloat32(reader)
}

func readFloat64Value(reader io.Reader) (interface{}, error) {
	return readFloat64(reader)
}

func readComplex64Value(reader io.Reader) (interface{}, error) {
	return readComplex64(reader)
}

func readComplex128Value(reader io.Reader) (interface{}, error) {
	return readComplex128(reader)
}

func readStringValue(reader io.Reader) (interface{}, error) {
	return readString(reader)
}

func readTimeValue(reader io.Reader) (interface{}, error) {
	return readTime(reader)
}

func readDurationValue(reader io.Reader) (interface{}, error) {
	u, err := readUint64(reader)
	return time.Duration(u), err
}

func readTime(reader io.Reader) (time.Time, error) {
	u, err := readUint64(reader)
	if err != nil {
		return time.Time{}, err
	}
	nano := int64(u)
	return time.Unix(nano/int64(time.Second), nano%int64(time.Second)), nil
}
