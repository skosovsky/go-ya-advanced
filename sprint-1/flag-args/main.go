package main

import (
	"log"
	"os"
)

func main() {
	log.Printf("Command %v\n", os.Args[0])

	for i, arg := range os.Args[1:] {
		log.Println(i+1, arg)
	}
}
