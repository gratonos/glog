package glog

import (
	"github.com/gratonos/glog/pkg/glog/logger/iface"
)

func MustRegisterErrorHandler(name string, handler iface.ErrorHandler) {
	if err := iface.RegisterErrorHandler(name, handler); err != nil {
		panic(err)
	}
}
