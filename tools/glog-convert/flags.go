package main

import (
	"flag"
	"fmt"
)

var (
	flagColoring bool
)

func initFlags() {
	flag.Usage = usage

	flag.BoolVar(&flagColoring, "color", true, "enable coloring")
}

func usage() {
	name := flag.CommandLine.Name()
	output := flag.CommandLine.Output()

	fmt.Fprintf(output, "Usage: %s [FLAG]... [PATH]...\n", name)
	fmt.Fprintln(output, "Flags:")
	flag.PrintDefaults()
}
