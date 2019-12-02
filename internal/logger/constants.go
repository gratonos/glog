package logger

import (
	"github.com/gratonos/glog/pkg/glog/logger/iface"
)

const (
	logBufLen = 512
)

const (
	timeHolder  = "00:00:00.000000"
	hourBegin   = 0
	hourEnd     = 2
	minuteBegin = 3
	minuteEnd   = 5
	secondBegin = 6
	secondEnd   = 8
	microBegin  = 9
	microEnd    = len(timeHolder)
)

var levelNames = []string{
	iface.Trace: "TRACE",
	iface.Debug: "DEBUG",
	iface.Info:  "INFO ",
	iface.Warn:  "WARN ",
	iface.Error: "ERROR",
	iface.Fatal: "FATAL",
}
