package main

import (
	"runtime"
	"sync"
	"time"
)

type conversionTask struct {
	Path    string
	ModTime time.Time
}

var (
	taskChan  chan conversionTask
	waitGroup sync.WaitGroup
)

func initPool() {
	taskChan = make(chan conversionTask, runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		waitGroup.Add(1)

		go func() {
			defer waitGroup.Done()

			for {
				if task, ok := <-taskChan; ok {
					processFile(task.Path, task.ModTime)
				} else {
					break
				}
			}
		}()
	}
}
