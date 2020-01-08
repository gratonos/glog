package text

import (
	"bytes"
	"fmt"
	"strconv"
	"time"

	"github.com/gratonos/glog/internal/encoding/binary"
	"github.com/gratonos/glog/pkg/glog/logger/iface"
)

const (
	timeLayout = "2006-01-02 15:04:05.000000"
	separator  = " "
	logMark    = "@@@@@@@@"
)

var levelDesc = [...]string{
	iface.Trace: "TRACE",
	iface.Debug: "DEBUG",
	iface.Info:  "INFO ",
	iface.Warn:  "WARN ",
	iface.Error: "ERROR",
	iface.Fatal: "FATAL",
}

var defaultFormats = [...]string{
	binary.Bool:       "%t",
	binary.Byte:       "%#02x",
	binary.Rune:       "%c",
	binary.Int8:       "%d",
	binary.Int16:      "%d",
	binary.Int32:      "%d",
	binary.Int64:      "%d",
	binary.Uint8:      "%d",
	binary.Uint16:     "%d",
	binary.Uint32:     "%d",
	binary.Uint64:     "%d",
	binary.Uintptr:    "%#x",
	binary.Float32:    "%g",
	binary.Float64:    "%g",
	binary.Complex64:  "%g",
	binary.Complex128: "%g",
	binary.String:     "%s",
	binary.Time:       timeLayout,
	binary.Duration:   "%s",
}

func FormatRecord(record *binary.Record, coloring bool) []byte {
	buf := new(bytes.Buffer)
	dyer := newTextDyer(buf, record.Level, coloring)

	formatTime(dyer, record.Time)
	formatLevel(dyer, record.Level)
	formatMark(dyer, record.Mark)
	formatPkg(dyer, record.Pkg)
	formatFunc(dyer, record.Func)
	formatFile(dyer, record.File)
	formatLine(dyer, record.Line)
	formatMsg(dyer, record.Msg)
	formatContexts(dyer, record.Contexts)

	buf.WriteByte('\n')
	return buf.Bytes()
}

func formatTime(dyer *textDyer, tm time.Time) {
	dyer.DyeContent(tm.Format(timeLayout))
}

func formatLevel(dyer *textDyer, level iface.Level) {
	if !iface.LegalLogLevel(level) {
		panic(fmt.Sprintf("glog: illegal log level: %d", level))
	}
	dyer.Write(separator)
	dyer.DyeContent(levelDesc[level])
}

func formatMark(dyer *textDyer, mark bool) {
	if mark {
		dyer.Write(separator)
		dyer.DyeSymbol(logMark)
	}
}

func formatPkg(dyer *textDyer, pkg string) {
	dyer.Write(separator)
	dyer.DyeContent(pkg)
}

func formatFunc(dyer *textDyer, fn string) {
	if fn != "" {
		dyer.Write(separator)
		dyer.DyeContent(fn)
	}
}

func formatFile(dyer *textDyer, file string) {
	if file != "" {
		dyer.Write(separator)
		dyer.DyeContent(file)
	}
}

func formatLine(dyer *textDyer, line int) {
	if line != 0 {
		dyer.Write(separator)
		dyer.DyeContent(strconv.Itoa(line))
	}
}

func formatMsg(dyer *textDyer, msg string) {
	dyer.Write(separator)
	dyer.DyeSymbol("<")
	dyer.DyeContent(msg)
	dyer.DyeSymbol(">")
}

func formatContexts(dyer *textDyer, contexts []binary.Context) {
	for _, context := range contexts {
		formatContext(dyer, context.Key, formatValue(&context))
	}
}

func formatContext(dyer *textDyer, key, value string) {
	dyer.Write(separator)
	dyer.DyeSymbol("(")
	dyer.DyeKey(key)
	dyer.DyeSymbol(":")
	dyer.Write(separator)
	dyer.DyeContent(value)
	dyer.DyeSymbol(")")
}

func formatValue(context *binary.Context) string {
	kind := context.Kind
	if !kind.Legal() {
		panic(fmt.Sprintf("glog: illegal value kind %d", kind))
	}

	format := context.Format
	if format == "" {
		format = defaultFormats[kind]
	}

	var value string
	if kind == binary.Time {
		tm := context.Value.(time.Time)
		value = tm.Format(format)
	} else {
		value = fmt.Sprintf(format, context.Value)
	}

	return value
}
