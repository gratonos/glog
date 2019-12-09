package glog

import (
	"runtime"
	"strings"
	"sync"

	ilogger "github.com/gratonos/glog/internal/logger"
	"github.com/gratonos/glog/pkg/glog/logger"
)

var (
	loggers = map[string]*ilogger.Logger{}
	lock    sync.Mutex
)

func Logger(name string) *logger.Logger {
	return logger.NewLogger(internalLogger(name), callerPkg())
}

func SugaredLogger(name string) *logger.SugaredLogger {
	return logger.NewSugaredLogger(internalLogger(name), callerPkg())
}

func internalLogger(name string) *ilogger.Logger {
	lock.Lock()
	defer lock.Unlock()

	logger := loggers[name]
	if logger == nil {
		logger = ilogger.New()
		loggers[name] = logger
	}

	return logger
}

func callerPkg() string {
	const stackOffset = 2
	pc, _, _, ok := runtime.Caller(stackOffset)
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
