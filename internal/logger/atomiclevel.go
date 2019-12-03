package logger

import (
	"sync/atomic"

	"github.com/gratonos/glog/pkg/glog/logger/iface"
)

type atomicLevel struct {
	level iface.Level
}

func (this *atomicLevel) Get() iface.Level {
	return iface.Level(atomic.LoadInt32((*int32)(&this.level)))
}

func (this *atomicLevel) Set(level iface.Level) {
	atomic.StoreInt32((*int32)(&this.level), int32(level))
}
