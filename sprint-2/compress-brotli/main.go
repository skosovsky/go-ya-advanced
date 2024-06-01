package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/andybalholm/brotli"
)

func BrotliCompress(data []byte) ([]byte, error) {
	var buf bytes.Buffer

	writer := brotli.NewWriterLevel(&buf, brotli.BestCompression)
	_, err := writer.Write(data)
	if err != nil {
		return nil, fmt.Errorf("error writing to brotli: %w", err)
	}

	err = writer.Close()
	if err != nil {
		return nil, fmt.Errorf("error closing writer: %w", err)
	}

	return buf.Bytes(), nil
}

func main() {
	data := []byte("Hello World")
	compressed, err := BrotliCompress(data)
	if err != nil {
		log.Println(err)
	}

	log.Printf("compressed data: %s", compressed)
}
