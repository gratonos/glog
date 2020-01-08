package logger

import (
	"fmt"
	"sync"
	"time"

	"github.com/gratonos/glog/internal/writers/console"
	"github.com/gratonos/glog/pkg/glog/logger/iface"
)

type Logger struct {
	consoleWriter *console.Writer

	config iface.Config
	level  *atomicLevel
	srcPos *atomicBool

	lock sync.Mutex
}

func New() *Logger {
	config := iface.DefaultConfig()
	return &Logger{
		consoleWriter: console.New(config.ConsoleConfig),
		level:         newAtomicLevel(config.Level),
		srcPos:        newAtomicBool(config.SrcPos),
		config:        config,
	}
}

func (this *Logger) Level() iface.Level {
	return this.level.Get()
}

func (this *Logger) SrcPos() bool {
	return this.srcPos.Get()
}

func (this *Logger) Config() iface.Config {
	this.lock.Lock()
	defer this.lock.Unlock()

	return this.config
}

func (this *Logger) SetConfig(config iface.Config) error {
	this.lock.Lock()
	defer this.lock.Unlock()

	return this.setConfig(config)
}

func (this *Logger) UpdateConfig(updater func(config iface.Config) iface.Config) error {
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

	if this.config.ConsoleWriter {
		this.consoleWriter.Write(log, tm)
	}

	this.lock.Unlock()

	done()
}

func (this *Logger) setConfig(config iface.Config) error {
	level := config.Level
	if !iface.LegalLoggerLevel(level) {
		return fmt.Errorf("glog: set config: illegal logger level: %d", level)
	}

	if err := this.consoleWriter.SetConfig(config.ConsoleConfig); err != nil {
		return fmt.Errorf("glog: set config: invalid config for console writer: %v", err)
	}

	this.config = config
	this.level.Set(level)
	this.srcPos.Set(config.SrcPos)

	return nil
}
