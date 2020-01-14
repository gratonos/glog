package iface

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type ErrorHandler func(time.Time, error)

type errorHandlerMap struct {
	handlers map[string]ErrorHandler
	lock     sync.Mutex
}

var errorHandlers = &errorHandlerMap{
	handlers: make(map[string]ErrorHandler),
}

func (this *errorHandlerMap) Register(name string, handler ErrorHandler) error {
	if name == "" {
		return errors.New("glog: register error handler: name is empty")
	}
	if handler == nil {
		return errors.New("glog: register error handler: handler is nil")
	}

	this.lock.Lock()
	defer this.lock.Unlock()

	if this.handlers[name] != nil {
		return fmt.Errorf("glog: register error handler: '%s' exists", name)
	}

	this.handlers[name] = handler
	return nil
}

func (this *errorHandlerMap) Get(name string) (ErrorHandler, error) {
	if name == "" {
		return nil, nil
	}

	this.lock.Lock()
	defer this.lock.Unlock()

	handler := this.handlers[name]
	if handler == nil {
		return nil, fmt.Errorf("glog: get error handler: '%s' does not exist", name)
	} else {
		return handler, nil
	}
}

func RegisterErrorHandler(name string, handler ErrorHandler) error {
	return errorHandlers.Register(name, handler)
}

func GetErrorHandler(name string) (ErrorHandler, error) {
	return errorHandlers.Get(name)
}
