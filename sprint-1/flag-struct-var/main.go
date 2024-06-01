package main

import (
	"flag"
	"log"
)

func main() {
	var options struct {
		width int
		thumb bool
	}

	flag.IntVar(&options.width, "width", 640, "width of image")
	flag.BoolVar(&options.thumb, "thumb", true, "create image")

	flag.Parse()

	log.Println(options.width, options.thumb)
}
