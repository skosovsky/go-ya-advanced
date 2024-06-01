package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var version = "0.0.1"

	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s - %s:\n", os.Args[0], version)
		flag.PrintDefaults()
	}

	flag.Parse()
}
