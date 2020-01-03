package text

import (
	"bytes"

	"github.com/gratonos/glog/pkg/glog/logger/iface"
)

const (
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	reset   = "\033[0m"
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
	mark     bool
	coloring bool
}

func newTextDyer(buf *bytes.Buffer, level iface.Level, mark bool, coloring bool) *textDyer {
	return &textDyer{
		buf:      buf,
		level:    level,
		mark:     mark,
		coloring: coloring,
	}
}

func (this *textDyer) DyeNormal(str string) {
	var begin, end string
	if this.coloring {
		if this.mark {
			begin = magenta
		} else {
			begin = levelColors[this.level]
		}
		end = reset
	}

	this.dye(begin, str, end)
}

func (this *textDyer) DyeLevel(level string) {
	var begin, end string
	if this.coloring {
		begin = levelColors[this.level]
		end = reset
	}

	this.dye(begin, level, end)
}

func (this *textDyer) DyeSymbol(symbol string) {
	var begin, end string
	if this.coloring {
		begin = cyan
		end = reset
	}

	this.dye(begin, symbol, end)
}

func (this *textDyer) DyeKey(key string) {
	var begin, end string
	if this.coloring {
		begin = blue
		end = reset
	}

	this.dye(begin, key, end)
}

func (this *textDyer) Write(str string) {
	this.buf.WriteString(str)
}

func (this *textDyer) dye(begin, str, end string) {
	this.buf.WriteString(begin)
	this.buf.WriteString(str)
	this.buf.WriteString(end)
}
