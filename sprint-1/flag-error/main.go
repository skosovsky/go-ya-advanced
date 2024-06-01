package main

import (
	"flag"
	"log"
)

func main() {
	imfFile := flag.String("file", "", "input image file")
	destDir := flag.String("dest", "./output", "destination directory")
	width := flag.Int("width", 1024, "width of the image")
	isThumb := flag.Bool("thumb", true, "create thumbnail")

	flag.Parse()

	for i, arg := range flag.Args() {
		log.Printf("arg[%d]=%s\n", i, arg)
	}

	log.Println("Image file:", *imfFile)
	log.Println("Destination directory:", *destDir)
	log.Println("Width:", *width)
	log.Println("Thumb:", *isThumb)
}
