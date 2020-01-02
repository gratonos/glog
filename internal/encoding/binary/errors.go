package binary

import (
	"errors"
	"fmt"
)

const (
	readingErrPrefix = "glog: read binary record"
)

type IOError struct {
	Err error
}

func newIOError(err error) *IOError {
	return &IOError{
		Err: err,
	}
}

func (this *IOError) Error() string {
	return fmt.Sprintf("%s: %v", readingErrPrefix, this.Err)
}

type MagicError struct {
	Magic []byte
}

func newMagicError(magic []byte) *MagicError {
	return &MagicError{
		Magic: magic,
	}
}

func (this *MagicError) Error() string {
	return fmt.Sprintf("%s: unmatched magic %02x", readingErrPrefix, this.Magic)
}

type VersionError struct {
	Version uint8
}

func newVersionError(version uint8) *VersionError {
	return &VersionError{
		Version: version,
	}
}

func (this *VersionError) Error() string {
	return fmt.Sprintf("%s: unsupported version %d", readingErrPrefix, this.Version)
}

type FormatError struct {
	Reason string
}

func newFormatError(reason string) *FormatError {
	return &FormatError{
		Reason: reason,
	}
}

func (this *FormatError) Error() string {
	return fmt.Sprintf("%s: %s", readingErrPrefix, this.Reason)
}

var EOF = errors.New(readingErrPrefix + ": end of file")
