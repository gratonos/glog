package console

import (
	"bytes"
	"os"

	"github.com/gratonos/glog/internal/encoding/binary"
	"github.com/gratonos/glog/internal/encoding/text"
	"github.com/gratonos/glog/pkg/glog/logger/iface"
)

type Writer struct {
	config iface.ConsoleConfig
}

func New(config iface.ConsoleConfig) *Writer {
	return &Writer{
		config: config,
	}
}

func (this *Writer) Write(log []byte) {
	var record binary.Record
	binary.ReadRecord(&record, bytes.NewBuffer(log))

	os.Stderr.Write(text.FormatRecord(&record, this.config.Coloring))
}

func (this *Writer) SetConfig(config iface.ConsoleConfig) {
	this.config = config
}
