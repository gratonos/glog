package binary

import (
	"bufio"
	"fmt"
	"io"
	"time"

	"github.com/gratonos/glog/pkg/glog/logger/iface"
)

type fieldKind uint8

const (
	fieldTimestamp fieldKind = iota
	fieldLevel
	fieldPkg
	fieldFunc
	fieldFile
	fieldLine
	fieldMark
	fieldMsg
	fieldContext
	fieldEnd

	fieldKindBound
)

func (self fieldKind) Legal() bool {
	return self < fieldKindBound
}

var binaryMagic = []byte{0x14, 0xf2, 0x79, 0xd3, 0x6b, 0xe7, 0x3d}

const binaryVersion = 0

const (
	sizeOfMagic   = 7
	sizeOfVersion = 1
	sizeOfHeader  = sizeOfMagic + sizeOfVersion
)

var fieldReaders = [...]func(*Record, *bufio.Reader) error{
	fieldTimestamp: readTimestamp,
	fieldLevel:     readLevel,
	fieldPkg:       readPkg,
	fieldFunc:      readFunc,
	fieldFile:      readFile,
	fieldLine:      readLine,
	fieldMark:      readMark,
	fieldMsg:       readMsg,
	fieldContext:   readContext,
}

func AppendBinaryMeta(dst []byte) []byte {
	dst = append(dst, binaryMagic...)
	dst = appendUint8(dst, binaryVersion)
	return dst
}

func ResetBuf(dst []byte) []byte {
	return dst[:sizeOfHeader]
}

func AppendTime(dst []byte, tm time.Time) []byte {
	dst = appendFieldKind(dst, fieldTimestamp)
	dst = appendUint64(dst, uint64(tm.UnixNano()))
	return dst
}

func AppendLevel(dst []byte, level iface.Level) []byte {
	dst = appendFieldKind(dst, fieldLevel)
	dst = appendUint8(dst, uint8(level))
	return dst
}

func AppendPkg(dst []byte, pkg string) []byte {
	dst = appendFieldKind(dst, fieldPkg)
	dst = appendShortString(dst, pkg)
	return dst
}

func AppendFunc(dst []byte, fn string) []byte {
	dst = appendFieldKind(dst, fieldFunc)
	dst = appendShortString(dst, fn)
	return dst
}

func AppendFile(dst []byte, file string) []byte {
	dst = appendFieldKind(dst, fieldFile)
	dst = appendShortString(dst, file)
	return dst
}

func AppendLine(dst []byte, line int) []byte {
	dst = appendFieldKind(dst, fieldLine)
	dst = appendUint32(dst, uint32(line))
	return dst
}

func AppendMark(dst []byte) []byte {
	return appendFieldKind(dst, fieldMark)
}

func AppendMsg(dst []byte, msg string) []byte {
	dst = appendFieldKind(dst, fieldMsg)
	dst = appendString(dst, msg)
	return dst
}

func AppendEnd(dst []byte) []byte {
	return appendFieldKind(dst, fieldEnd)
}

func appendFieldKind(dst []byte, kind fieldKind) []byte {
	return appendUint8(dst, uint8(kind))
}

func readMagic(reader *bufio.Reader) ([]byte, error) {
	magic := make([]byte, sizeOfMagic)
	if err := read(magic, reader); err != nil {
		if ioErr, ok := err.(*IOError); ok && ioErr.Err == io.EOF {
			return nil, EOF
		} else {
			return nil, err
		}
	}
	return magic, nil
}

func readVersion(reader *bufio.Reader) (uint8, error) {
	return readUint8(reader)
}

func readFields(record *Record, reader *bufio.Reader) error {
	for {
		kind, err := readFieldKind(reader)
		if err != nil {
			return err
		}

		if kind == fieldEnd {
			return nil
		}

		err = fieldReaders[kind](record, reader)
		if err != nil {
			return err
		}
	}
}

func readTimestamp(record *Record, reader *bufio.Reader) error {
	tm, err := readTime(reader)
	if err == nil {
		record.Time = tm
	}
	return err
}

func readLevel(record *Record, reader *bufio.Reader) error {
	u, err := readUint8(reader)
	if err != nil {
		return err
	}

	level := iface.Level(u)
	if !iface.LegalLogLevel(level) {
		return newFormatError(fmt.Sprintf("illegal log level: %d", level))
	}

	record.Level = level
	return nil
}

func readPkg(record *Record, reader *bufio.Reader) error {
	pkg, err := readShortString(reader)
	if err == nil {
		record.Pkg = pkg
	}
	return err
}

func readFunc(record *Record, reader *bufio.Reader) error {
	fn, err := readShortString(reader)
	if err == nil {
		record.Func = fn
	}
	return err
}

func readFile(record *Record, reader *bufio.Reader) error {
	file, err := readShortString(reader)
	if err == nil {
		record.File = file
	}
	return err
}

func readLine(record *Record, reader *bufio.Reader) error {
	line, err := readUint32(reader)
	if err == nil {
		record.Line = int(line)
	}
	return err
}

func readMark(record *Record, _ *bufio.Reader) error {
	record.Mark = true
	return nil
}

func readMsg(record *Record, reader *bufio.Reader) error {
	msg, err := readString(reader)
	if err == nil {
		record.Msg = msg
	}
	return err
}

func readContext(record *Record, reader *bufio.Reader) error {
	key, kind, err := readContextMeta(reader)
	if err != nil {
		return err
	}

	value, err := valueReaders[kind](reader)
	if err != nil {
		return err
	}

	context := Context{
		Key:   key,
		Kind:  kind,
		Value: value,
	}
	record.Contexts = append(record.Contexts, context)
	return nil
}

func readFieldKind(reader *bufio.Reader) (fieldKind, error) {
	u, err := readUint8(reader)
	if err != nil {
		return fieldKindBound, err
	}

	kind := fieldKind(u)
	if !kind.Legal() {
		return fieldKindBound, newFormatError(fmt.Sprintf("illegal field kind %d", kind))
	}

	return kind, nil
}
