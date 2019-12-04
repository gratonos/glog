package logger

import (
	"fmt"

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

const (
	floatFormat    = 'f'
	floatPrecision = -1
	integerBase    = 10
)

var levelNames = [...]string{
	iface.Trace: "TRACE",
	iface.Debug: "DEBUG",
	iface.Info:  "INFO ",
	iface.Warn:  "WARN ",
	iface.Error: "ERROR",
	iface.Fatal: "FATAL",
}

var smallIntReps [100]string
var byteReps [256]string

func init() {
	for i := range smallIntReps {
		smallIntReps[i] = fmt.Sprintf("%02d", i)
	}
	for i := range byteReps {
		byteReps[i] = fmt.Sprintf("%02x", i)
	}
}
