package text

import (
	"bytes"

	"github.com/gratonos/glog/pkg/glog/iface"
)

const (
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Reset   = "\033[0m"
)

var levelColors = [...]string{
	iface.Trace: Green,
	iface.Debug: Green,
	iface.Info:  Green,
	iface.Warn:  Yellow,
	iface.Error: Red,
	iface.Fatal: Red,
}

type textDyer struct {
	buf      *bytes.Buffer
	level    iface.Level
	coloring bool
}

func newTextDyer(buf *bytes.Buffer, level iface.Level, coloring bool) *textDyer {
	return &textDyer{
		buf:      buf,
		level:    level,
		coloring: coloring,
	}
}

func (this *textDyer) DyeContent(str string) {
	if this.coloring {
		this.dye(str, levelColors[this.level])
	} else {
		this.Write(str)
	}
}

func (this *textDyer) DyeSymbol(symbol string) {
	if this.coloring {
		this.dye(symbol, Cyan)
	} else {
		this.Write(symbol)
	}
}

func (this *textDyer) DyeKey(key string) {
	if this.coloring {
		this.dye(key, Blue)
	} else {
		this.Write(key)
	}
}

func (this *textDyer) Write(str string) {
	this.buf.WriteString(str)
}

func (this *textDyer) dye(str, color string) {
	this.buf.WriteString(color)
	this.buf.WriteString(str)
	this.buf.WriteString(Reset)
}
