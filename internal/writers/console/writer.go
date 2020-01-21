package console

import (
	"fmt"
	"os"
	"time"

	"github.com/gratonos/glog/internal/util"
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
	text, err := util.BinaryToText(log, this.config.TextConfig.Coloring)
	if err != nil {
		panic(fmt.Sprintf("glog: corrupted log: %v", err))
	}

	_, err = os.Stderr.Write(text)
	if err != nil && this.config.ErrorHandler != nil {
		this.config.ErrorHandler(tm, err)
	}
}

func (this *Writer) SetConfig(config iface.ConsoleWriter) error {
	this.config = config
	return nil
}
