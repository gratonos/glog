package console

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/gratonos/glog/internal/encoding/binary"
	"github.com/gratonos/glog/internal/encoding/text"
	"github.com/gratonos/glog/internal/writers"
	"github.com/gratonos/glog/pkg/glog/logger/iface"
)

type Writer struct {
	errorHandler writers.ErrorHandler
	config       iface.ConsoleConfig
}

func New(config iface.ConsoleConfig) *Writer {
	writer := &Writer{}
	if err := writer.SetConfig(config); err != nil {
		panic(fmt.Sprintf("glog: invalid default config for console writer: %v", err))
	}
	return writer
}

func (this *Writer) Write(log []byte, tm time.Time) {
	var record binary.Record
	if err := binary.ReadRecord(&record, bytes.NewBuffer(log)); err != nil {
		panic(fmt.Sprintf("glog: corrupted log: %v", err))
	}

	_, err := os.Stderr.Write(text.FormatRecord(&record, this.config.Coloring))
	if err != nil && this.errorHandler != nil {
		this.errorHandler(tm, err)
	}
}

func (this *Writer) SetConfig(config iface.ConsoleConfig) error {
	errorHandler, err := writers.GetErrorHandler(config.ErrorHandler)
	if err != nil {
		return err
	}
	this.errorHandler = errorHandler

	this.config = config
	return nil
}
