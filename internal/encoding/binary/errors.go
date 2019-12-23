package binary

import (
	"errors"
	"fmt"
)

const (
	decodingErrPrefix = "glog: decode binary record"
)

type IOError struct {
	Err error
}

func newIOError(err error) *IOError {
	return &IOError{
		Err: err,
	}
}

func (err *IOError) Error() string {
	return fmt.Sprintf("%s: %v", decodingErrPrefix, err.Err)
}

type MagicError struct {
	Magic []byte
}

func newMagicError(magic []byte) *MagicError {
	return &MagicError{
		Magic: magic,
	}
}

func (err *MagicError) Error() string {
	return fmt.Sprintf("%s: unmatched magic %02x", decodingErrPrefix, err.Magic)
}

type VersionError struct {
	Version uint8
}

func newVersionError(version uint8) *VersionError {
	return &VersionError{
		Version: version,
	}
}

func (err *VersionError) Error() string {
	return fmt.Sprintf("%s: unsupported version %d", decodingErrPrefix, err.Version)
}

type FormatError struct {
	Reason string
}

func newFormatError(reason string) *FormatError {
	return &FormatError{
		Reason: reason,
	}
}

func (err *FormatError) Error() string {
	return fmt.Sprintf("%s: %v", decodingErrPrefix, err.Reason)
}

var EOF = errors.New(decodingErrPrefix + ": end of file")
