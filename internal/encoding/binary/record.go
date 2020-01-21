package binary

import (
	"bytes"
	"io"
	"time"

	"github.com/gratonos/glog/pkg/glog/iface"
)

type Record struct {
	Time     time.Time
	Pkg      string
	File     string
	Line     int
	Msg      string
	Contexts []Context
	Mark     bool
	Level    iface.Level
}

func ReadRecord(record *Record, reader io.Reader) error {
	if record == nil {
		panic(readingErrPrefix + ": record is nil")
	}
	if reader == nil {
		panic(readingErrPrefix + ": reader is nil")
	}

	magic, err := readMagic(reader)
	if err != nil {
		return err
	}
	if !bytes.Equal(magic, binaryMagic) {
		return newMagicError(magic)
	}

	return readRecord(record, reader)
}

func TryReadRecord(record *Record, reader io.Reader) error {
	if record == nil {
		panic(readingErrPrefix + ": record is nil")
	}
	if reader == nil {
		panic(readingErrPrefix + ": reader is nil")
	}

	_, err := searchMagic(reader)
	if err != nil {
		return err
	}

	return readRecord(record, reader)
}

func readRecord(record *Record, reader io.Reader) error {
	version, err := readVersion(reader)
	if err != nil {
		return err
	}
	if version != binaryVersion {
		return newVersionError(version)
	}

	return readFields(record, reader)
}
