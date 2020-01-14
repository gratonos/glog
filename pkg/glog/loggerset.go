package glog

import (
	"runtime"
	"strings"
	"sync"

	ilog "github.com/gratonos/glog/internal/logger"
	"github.com/gratonos/glog/pkg/glog/logger"
)

var (
	loggers = map[string]*ilog.Logger{}
	lock    sync.Mutex
)

func Logger(name string) *logger.Logger {
	return logger.NewLogger(internalLogger(name), callerPkg(0+1))
}

func internalLogger(name string) *ilog.Logger {
	lock.Lock()
	defer lock.Unlock()

	logger := loggers[name]
	if logger == nil {
		logger = ilog.New()
		loggers[name] = logger
	}

	return logger
}

func callerPkg(frameSkip int) string {
	pc, _, _, ok := runtime.Caller(frameSkip + 1)
	if ok {
		return pkgName(pc)
	} else {
		return "???"
	}
}

func pkgName(pc uintptr) string {
	name := runtime.FuncForPC(pc).Name()
	lastSlash := strings.LastIndexByte(name, '/')
	nextDot := strings.IndexByte(name[lastSlash+1:], '.')
	if nextDot < 0 {
		return "???"
	}
	return name[:lastSlash+1+nextDot]
}
