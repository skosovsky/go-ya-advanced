package main

import (
	"flag"
	"log"
	"strings"
)

func main() {
	var effects []string
	flag.Func("effects", "Rotation and mirror", func(flagValue string) error {
		effects = strings.Split(flagValue, " ")

		return nil
	})

	flag.Parse()
	log.Println(effects, len(effects))
}
