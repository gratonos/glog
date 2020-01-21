package util

import (
	"bytes"

	"github.com/gratonos/glog/internal/encoding/binary"
	"github.com/gratonos/glog/internal/encoding/text"
)

func BinaryToText(log []byte, coloring bool) ([]byte, error) {
	var record binary.Record
	if err := binary.ReadRecord(&record, bytes.NewBuffer(log)); err != nil {
		return nil, err
	}
	return text.FormatRecord(&record, coloring), nil
}
