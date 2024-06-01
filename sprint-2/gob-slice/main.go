package main

import (
	"bytes"
	"encoding/gob"
	"log"
)

func main() {
	data := []byte{12, 255, 129, 2, 1, 2, 255, 130, 0, 1, 12, 0, 0, 17, 255, 130, 0, 2, 6, 72, 101, 108, 108, 111, 44, 5, 119, 111, 114, 108, 100}

	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	result := make([]string, 0)

	if err := decoder.Decode(&result); err != nil {
		log.Fatal(err)
	}

	log.Println(result)
}
