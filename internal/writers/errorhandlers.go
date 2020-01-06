package writers

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type ErrorHandler func(time.Time, error)

var (
	handlers = make(map[string]ErrorHandler)
	lock     sync.Mutex
)

func MustRegisterErrorHandler(name string, handler ErrorHandler) {
	if err := RegisterErrorHandler(name, handler); err != nil {
		panic(err)
	}
}

func RegisterErrorHandler(name string, handler ErrorHandler) error {
	if name == "" {
		return errors.New("glog: register error handler: name is empty")
	}

	lock.Lock()
	defer lock.Unlock()

	if handlers[name] != nil {
		return fmt.Errorf("glog: register error handler: '%s' exists", name)
	}

	handlers[name] = handler
	return nil
}

func GetErrorHandler(name string) (ErrorHandler, error) {
	if name == "" {
		return nil, nil
	}

	lock.Lock()
	defer lock.Unlock()

	handler := handlers[name]
	if handler == nil {
		return nil, fmt.Errorf("glog: get error handler: '%s' does not exist", name)
	} else {
		return handler, nil
	}
}
