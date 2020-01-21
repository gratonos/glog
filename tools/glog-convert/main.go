package main

import (
	"flag"
)

func init() {
	initFlags()
	initLog()
	initPool()
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		args = []string{"."}
	}

	for _, arg := range args {
		processPath(arg)
	}

	close(taskChan)
	waitGroup.Wait()
}
