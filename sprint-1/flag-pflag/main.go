package main

import (
	"flag"
	"log"

	"github.com/spf13/pflag"
)

func main() {
	var ip = pflag.Int("flag", 1234, "help message for flag")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	log.Println(*ip)
}
