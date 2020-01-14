package console

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/gratonos/glog/internal/encoding/binary"
	"github.com/gratonos/glog/internal/encoding/text"
	"github.com/gratonos/glog/pkg/glog/iface"
)

type Writer struct {
	config iface.ConsoleWriter
}

func New(config iface.ConsoleWriter) *Writer {
	writer := &Writer{}
	return writer
}

func (this *Writer) Write(log []byte, tm time.Time) {
	var record binary.Record
	if err := binary.ReadRecord(&record, bytes.NewBuffer(log)); err != nil {
		panic(fmt.Sprintf("glog: corrupted log: %v", err))
	}

	_, err := os.Stderr.Write(text.FormatRecord(&record, this.config.Coloring))
	if err != nil && this.config.ErrorHandler != nil {
		this.config.ErrorHandler(tm, err)
	}
}

func (this *Writer) SetConfig(config iface.ConsoleWriter) error {
	this.config = config
	return nil
}
