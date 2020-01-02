package binary

import (
	"bufio"
	"bytes"
	"io"
	"time"

	"github.com/gratonos/glog/pkg/glog/logger/iface"
)

type Record struct {
	Time     time.Time
	Pkg      string
	Func     string
	File     string
	Line     int
	Msg      string
	Contexts []Context
	Mark     bool
	Level    iface.Level
}

func ReadRecord(record *Record, reader io.Reader) error {
	if record == nil {
		panic(readingErrPrefix + ": nil record")
	}
	if reader == nil {
		panic(readingErrPrefix + ": nil reader")
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
