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
	lock.Lock()
	defer lock.Unlock()

	log := loggers[name]
	if log == nil {
		log = ilog.New()
		loggers[name] = log
	}

	return logger.New(log, callerPkg())
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