package main

import (
	"fmt"
	"log"

	"github.com/gratonos/glog/internal/encoding/text"
)

func initLog() {
	log.SetFlags(0)
}

func infof(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	log.Printf("%s%s: %s%s\n", text.Green, "info", msg, text.Reset)
}

func warnf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	log.Printf("%s%s: %s%s\n", text.Yellow, "warn", msg, text.Reset)
}

func errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	log.Printf("%s%s: %s%s\n", text.Red, "error", msg, text.Reset)
}
