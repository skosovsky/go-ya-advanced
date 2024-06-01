package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	cnvFLags := flag.NewFlagSet("cnv", flag.ExitOnError)
	filterFlags := flag.NewFlagSet("filter", flag.ExitOnError)

	destDir := cnvFLags.String("dest", "./output", "Destination directory")
	width := cnvFLags.Int("width", 1024, "Width of output image") //nolint:mnd // example
	isThumb := cnvFLags.Bool("thumb", false, "create thumbnail image")

	isGray := filterFlags.Bool("gray", false, "convert to grayscale")
	isSepia := filterFlags.Bool("sepia", false, "convert to sepia")

	if len(os.Args) < 2 { //nolint:mnd // example
		log.Println("set or get subcommand required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "cnv":
		err := cnvFLags.Parse(os.Args[2:])
		if err != nil {
			log.Println(err)

			return
		}
	case "filter":
		err := filterFlags.Parse(os.Args[2:])
		if err != nil {
			log.Println(err)

			return
		}
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if cnvFLags.Parsed() {
		log.Println("cnv flag parsed")
		log.Println("dest:", *destDir)
		log.Println("width:", *width)
		log.Println("isThumb:", *isThumb)
	}

	if filterFlags.Parsed() {
		log.Println("filter flag parsed")
		log.Println("isGray:", *isGray)
		log.Println("isSepia:", *isSepia)
	}
}
