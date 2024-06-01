package main

import (
	"flag"
	"log"
)

func main() {
	imgFile := flag.String("file", "", "input image file")
	flag.Parse()

	log.Println("Image file:", *imgFile)
}
