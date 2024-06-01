package main

import (
	"flag"
	"log"
	"strings"
)

type Value interface {
	String() string
	Set(string) error
}

type Options struct {
	width   int  //nolint:unused // example
	thumb   bool //nolint:unused // example
	effects []string
}

func (o *Options) String() string {
	return strings.Join(o.effects, ",")
}

func (o *Options) Set(flagValue string) error {
	o.effects = strings.Split(flagValue, ",")

	return nil
}

func main() {
	options := new(Options)
	flag.Var(options, "effects", "Rotation and mirror")
	flag.Parse()

	log.Println(*options)
}
