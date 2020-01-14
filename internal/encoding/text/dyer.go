package text

import (
	"bytes"

	"github.com/gratonos/glog/pkg/glog/iface"
)

const (
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	cyan   = "\033[36m"
	reset  = "\033[0m"
)

var levelColors = [...]string{
	iface.Trace: green,
	iface.Debug: green,
	iface.Info:  green,
	iface.Warn:  yellow,
	iface.Error: red,
	iface.Fatal: red,
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
		this.dye(symbol, cyan)
	} else {
		this.Write(symbol)
	}
}

func (this *textDyer) DyeKey(key string) {
	if this.coloring {
		this.dye(key, blue)
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
	this.buf.WriteString(reset)
}
