package logger

import (
	"fmt"
	"sync"
	"time"

	"github.com/gratonos/glog/internal/writers/console"
	"github.com/gratonos/glog/internal/writers/file"
	"github.com/gratonos/glog/pkg/glog/iface"
)

type Logger struct {
	consoleWriter *console.Writer
	fileWriter    *file.Writer

	config   iface.Logger
	level    *atomicLevel
	fileLine *atomicBool

	lock sync.Mutex
}

func New() *Logger {
	config := iface.Logger{
		Level:    iface.Trace,
		FileLine: true,
		ConsoleWriter: iface.ConsoleWriter{
			Coloring: true,
			Enable:   true,
		},
	}

	consoleWriter := new(console.Writer)
	if err := consoleWriter.SetConfig(config.ConsoleWriter); err != nil {
		panic(fmt.Sprintf("glog: invalid default config for console writer: %v", err))
	}

	return &Logger{
		consoleWriter: consoleWriter,
		fileWriter:    new(file.Writer),
		config:        config,
		level:         newAtomicLevel(config.Level),
		fileLine:      newAtomicBool(config.FileLine),
	}
}

func (this *Logger) Level() iface.Level {
	return this.level.Get()
}

func (this *Logger) FileLine() bool {
	return this.fileLine.Get()
}

func (this *Logger) Config() iface.Logger {
	this.lock.Lock()
	defer this.lock.Unlock()

	return this.config
}

func (this *Logger) SetConfig(config iface.Logger) error {
	this.lock.Lock()
	defer this.lock.Unlock()

	return this.setConfig(config)
}

func (this *Logger) UpdateConfig(updater func(config iface.Logger) iface.Logger) error {
	if updater == nil {
		panic("glog: update config: updater is nil")
	}

	this.lock.Lock()
	defer this.lock.Unlock()

	return this.setConfig(updater(this.config))
}

func (this *Logger) Commit(emit func(time.Time) []byte, done func()) {
	this.lock.Lock()

	tm := time.Now()
	log := emit(tm)

	if this.config.ConsoleWriter.Enable {
		this.consoleWriter.Write(log, tm)
	}
	if this.config.FileWriter.Enable {
		this.fileWriter.Write(log, tm)
	}

	this.lock.Unlock()

	done()
}

func (this *Logger) setConfig(config iface.Logger) error {
	level := config.Level
	if !level.LegalForLogger() {
		return fmt.Errorf("glog: set config: illegal logger level: %d", level)
	}

	if err := this.consoleWriter.SetConfig(config.ConsoleWriter); err != nil {
		return fmt.Errorf("glog: set config: invalid config for console writer: %v", err)
	}

	if err := this.fileWriter.SetConfig(config.FileWriter); err != nil {
		return fmt.Errorf("glog: set config: invalid config for file writer: %v", err)
	}

	this.config = config
	this.level.Set(level)
	this.fileLine.Set(config.FileLine)

	return nil
}
