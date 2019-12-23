package binary

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"time"
	"unsafe"

	"github.com/gratonos/glog/pkg/glog/logger/iface"
)

type Pair struct {
	Key    string
	Value  interface{}
	Format string
	Kind   ValueKind
}

type Record struct {
	Time  time.Time
	Pkg   string
	Func  string
	File  string
	Line  int
	Msg   string
	Pairs []Pair
	Mark  bool
	Level iface.Level
}

var fieldHandlers = [...]func(*Record, *bufio.Reader) error{
	Timestamp: readTimestamp,
	Level:     readLevel,
	Pkg:       readPkg,
	Func:      readFunc,
	File:      readFile,
	Line:      readLine,
	Mark:      readMark,
	Msg:       readMsg,
	KeyValue:  readKeyValue,
}

var valueHandlers = [...]func(*bufio.Reader) (interface{}, error){
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

func ReadRecord(record *Record, reader io.Reader) error {
	if record == nil {
		panic(decodingErrPrefix + ": nil record")
	}
	if reader == nil {
		panic(decodingErrPrefix + ": nil reader")
	}
	return readRecord(record, bufio.NewReader(reader))
}

func readRecord(record *Record, reader *bufio.Reader) error {
	magic, err := readMagic(reader)
	if err != nil {
		return err
	}
	if !bytes.Equal(magic, binaryMagic) {
		return newMagicError(magic)
	}

	version, err := readVersion(reader)
	if err != nil {
		return err
	}
	if version != binaryVersion {
		return newVersionError(version)
	}

	return readFields(record, reader)
}

func readMagic(reader *bufio.Reader) (magic []byte, err error) {
	magic = make([]byte, sizeOfMagic)
	if err = readBytes(magic, reader); err != nil {
		if ioerr, ok := err.(*IOError); ok && ioerr.Err == io.EOF {
			return nil, EOF
		} else {
			return nil, err
		}
	}
	return magic, nil
}

func readVersion(reader *bufio.Reader) (version uint8, err error) {
	return readUint8(reader)
}

func readFields(record *Record, reader *bufio.Reader) error {
	for {
		u8, err := readUint8(reader)
		if err != nil {
			return err
		}

		kind := FieldKind(u8)
		if kind == End {
			break
		} else if kind > MaxFieldKind {
			return newFormatError(fmt.Sprintf("illegal field kind %d", kind))
		}

		err = fieldHandlers[kind](record, reader)
		if err != nil {
			return err
		}
	}

	return nil
}

func readTimestamp(record *Record, reader *bufio.Reader) error {
	tm, err := readTime(reader)
	if err != nil {
		return err
	}
	record.Time = tm
	return nil
}

func readLevel(record *Record, reader *bufio.Reader) error {
	u8, err := readUint8(reader)
	if err != nil {
		return err
	}
	level := iface.Level(u8)
	if level < iface.Trace || level > iface.Fatal {
		return newFormatError(fmt.Sprintf("illegal log level: %d", level))
	}
	record.Level = level
	return nil
}

func readPkg(record *Record, reader *bufio.Reader) error {
	pkg, err := readShortString(reader)
	if err != nil {
		return err
	}
	record.Pkg = pkg
	return nil
}

func readFunc(record *Record, reader *bufio.Reader) error {
	fn, err := readShortString(reader)
	if err != nil {
		return err
	}
	record.Func = fn
	return nil
}

func readFile(record *Record, reader *bufio.Reader) error {
	file, err := readShortString(reader)
	if err != nil {
		return err
	}
	record.File = file
	return nil
}

func readLine(record *Record, reader *bufio.Reader) error {
	line, err := readUint32(reader)
	if err != nil {
		return err
	}
	record.Line = int(line)
	return nil
}

func readMark(record *Record, _ *bufio.Reader) error {
	record.Mark = true
	return nil
}

func readMsg(record *Record, reader *bufio.Reader) error {
	msg, err := readString(reader)
	if err != nil {
		return err
	}
	record.Msg = msg
	return nil
}

func readKeyValue(record *Record, reader *bufio.Reader) error {
	key, kind, err := readKVMeta(reader)
	if err != nil {
		return err
	}

	pair := Pair{
		Key:  key,
		Kind: kind,
	}
	value, err := valueHandlers[kind](reader)
	if err != nil {
		return err
	}

	pair.Value = value
	record.Pairs = append(record.Pairs, pair)
	return nil
}

func readKVMeta(reader *bufio.Reader) (key string, kind ValueKind, err error) {
	key, err = readKey(reader)
	if err != nil {
		return "", IllegalValueKind, err
	}
	kind, err = readValueKind(reader)
	if err != nil {
		return "", IllegalValueKind, err
	}
	return key, kind, nil
}

func readKey(reader *bufio.Reader) (key string, err error) {
	return readShortString(reader)
}

func readValueKind(reader *bufio.Reader) (kind ValueKind, err error) {
	u8, err := readUint8(reader)
	if err != nil {
		return IllegalValueKind, err
	}
	kind = ValueKind(u8)
	if kind > MaxValueKind {
		return IllegalValueKind, newFormatError(fmt.Sprintf("illegal value kind %d", kind))
	}
	return kind, nil
}

func readBoolValue(reader *bufio.Reader) (value interface{}, err error) {
	return readBool(reader)
}

func readByteValue(reader *bufio.Reader) (value interface{}, err error) {
	return readUint8(reader)
}

func readRuneValue(reader *bufio.Reader) (value interface{}, err error) {
	u32, err := readUint32(reader)
	return rune(u32), err
}

func readInt8Value(reader *bufio.Reader) (value interface{}, err error) {
	u8, err := readUint8(reader)
	return int8(u8), err
}

func readInt16Value(reader *bufio.Reader) (value interface{}, err error) {
	u16, err := readUint16(reader)
	return int16(u16), err
}

func readInt32Value(reader *bufio.Reader) (value interface{}, err error) {
	u32, err := readUint32(reader)
	return int32(u32), err
}

func readInt64Value(reader *bufio.Reader) (value interface{}, err error) {
	u64, err := readUint64(reader)
	return int64(u64), err
}

func readUint8Value(reader *bufio.Reader) (value interface{}, err error) {
	return readUint8(reader)
}

func readUint16Value(reader *bufio.Reader) (value interface{}, err error) {
	return readUint16(reader)
}

func readUint32Value(reader *bufio.Reader) (value interface{}, err error) {
	return readUint32(reader)
}

func readUint64Value(reader *bufio.Reader) (value interface{}, err error) {
	return readUint64(reader)
}

func readUintptrValue(reader *bufio.Reader) (value interface{}, err error) {
	u64, err := readUint64(reader)
	return uintptr(u64), err
}

func readFloat32Value(reader *bufio.Reader) (value interface{}, err error) {
	return readFloat32(reader)
}

func readFloat64Value(reader *bufio.Reader) (value interface{}, err error) {
	return readFloat64(reader)
}

func readComplex64Value(reader *bufio.Reader) (value interface{}, err error) {
	return readComplex64(reader)
}

func readComplex128Value(reader *bufio.Reader) (value interface{}, err error) {
	return readComplex128(reader)
}

func readStringValue(reader *bufio.Reader) (value interface{}, err error) {
	return readString(reader)
}

func readTimeValue(reader *bufio.Reader) (value interface{}, err error) {
	return readTime(reader)
}

func readDurationValue(reader *bufio.Reader) (value interface{}, err error) {
	return readDuration(reader)
}

func readBool(reader *bufio.Reader) (bool, error) {
	u8, err := readUint8(reader)
	if err != nil {
		return false, err
	}
	if u8 > 1 {
		return false, newFormatError(fmt.Sprintf("illegal bool value: %d", u8))
	}
	if u8 == 1 {
		return true, nil
	} else {
		return false, nil
	}
}

func readUint8(reader *bufio.Reader) (u8 uint8, err error) {
	buf := make([]byte, unsafe.Sizeof(u8))
	err = readBytes(buf, reader)
	if err != nil {
		return 0, err
	} else {
		return buf[0], nil
	}
}

func readUint16(reader *bufio.Reader) (u16 uint16, err error) {
	buf := make([]byte, unsafe.Sizeof(u16))
	err = readBytes(buf, reader)
	if err != nil {
		return 0, err
	} else {
		return binary.LittleEndian.Uint16(buf), nil
	}
}

func readUint32(reader *bufio.Reader) (u32 uint32, err error) {
	buf := make([]byte, unsafe.Sizeof(u32))
	err = readBytes(buf, reader)
	if err != nil {
		return 0, err
	} else {
		return binary.LittleEndian.Uint32(buf), nil
	}
}

func readUint64(reader *bufio.Reader) (u64 uint64, err error) {
	buf := make([]byte, unsafe.Sizeof(u64))
	err = readBytes(buf, reader)
	if err != nil {
		return 0, err
	} else {
		return binary.LittleEndian.Uint64(buf), nil
	}
}

func readFloat32(reader *bufio.Reader) (f32 float32, err error) {
	buf := make([]byte, unsafe.Sizeof(f32))
	err = readBytes(buf, reader)
	if err != nil {
		return 0, err
	} else {
		return math.Float32frombits(binary.LittleEndian.Uint32(buf)), nil
	}
}

func readFloat64(reader *bufio.Reader) (f64 float64, err error) {
	buf := make([]byte, unsafe.Sizeof(f64))
	err = readBytes(buf, reader)
	if err != nil {
		return 0, err
	} else {
		return math.Float64frombits(binary.LittleEndian.Uint64(buf)), nil
	}
}

func readComplex64(reader *bufio.Reader) (c64 complex64, err error) {
	r, err := readFloat32(reader)
	if err != nil {
		return 0, err
	}
	i, err := readFloat32(reader)
	if err != nil {
		return 0, err
	}
	return complex(r, i), nil
}

func readComplex128(reader *bufio.Reader) (c128 complex128, err error) {
	r, err := readFloat64(reader)
	if err != nil {
		return 0, err
	}
	i, err := readFloat64(reader)
	if err != nil {
		return 0, err
	}
	return complex(r, i), nil
}

func readString(reader *bufio.Reader) (string, error) {
	size, err := readUint16(reader)
	if err != nil {
		return "", err
	}
	return readStr(reader, uint(size))
}

func readShortString(reader *bufio.Reader) (string, error) {
	size, err := readUint8(reader)
	if err != nil {
		return "", err
	}
	return readStr(reader, uint(size))
}

func readStr(reader *bufio.Reader, size uint) (string, error) {
	buf := make([]byte, size)
	if err := readBytes(buf, reader); err != nil {
		return "", err
	}
	return string(buf), nil
}

func readTime(reader *bufio.Reader) (tm time.Time, err error) {
	u64, err := readUint64(reader)
	if err != nil {
		return time.Time{}, err
	}
	nano := int64(u64)
	return time.Unix(nano/int64(time.Second), nano%int64(time.Second)), nil
}

func readDuration(reader *bufio.Reader) (duration time.Duration, err error) {
	u64, err := readUint64(reader)
	return time.Duration(u64), err
}

func readBytes(buf []byte, reader *bufio.Reader) error {
	_, err := io.ReadFull(reader, buf)
	if err != nil {
		return newIOError(err)
	}
	return nil
}
