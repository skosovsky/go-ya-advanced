package main

import (
	"flag"
	"os"
)

var (
	flagRunAddr  string //nolint:gochecknoglobals // example
	flagLogLevel string //nolint:gochecknoglobals // example
)

func parseFlags() {
	flag.StringVar(&flagRunAddr, "a", ":8080", "address and port to run server")
	flag.StringVar(&flagLogLevel, "l", "info", "log level")
	flag.Parse()

	if envRunAddr, ok := os.LookupEnv("RUN_ADDR"); ok {
		flagRunAddr = envRunAddr
	}
	if envLogLevel, ok := os.LookupEnv("LOG_LEVEL"); ok {
		flagLogLevel = envLogLevel
	}
}
