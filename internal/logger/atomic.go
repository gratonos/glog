package logger

import (
	"sync/atomic"

	"github.com/gratonos/glog/pkg/glog/logger/iface"
)

type atomicLevel struct {
	level iface.Level
}

func newAtomicLevel(level iface.Level) *atomicLevel {
	atom := new(atomicLevel)
	atom.Set(level)
	return atom
}

func (this *atomicLevel) Get() iface.Level {
	return iface.Level(atomic.LoadInt32((*int32)(&this.level)))
}

func (this *atomicLevel) Set(level iface.Level) {
	atomic.StoreInt32((*int32)(&this.level), int32(level))
}

type atomicBool struct {
	value int32
}

func newAtomicBool(b bool) *atomicBool {
	atom := new(atomicBool)
	atom.Set(b)
	return atom
}

func (this *atomicBool) Get() bool {
	return atomic.LoadInt32(&this.value) == 1
}

func (this *atomicBool) Set(b bool) {
	if b {
		atomic.StoreInt32(&this.value, 1)
	} else {
		atomic.StoreInt32(&this.value, 0)
	}
}
