package main

import (
	"bytes"
	"compress/flate"
	"fmt"
	"log"
)

func Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer

	writer, err := flate.NewWriter(&buf, flate.BestCompression)
	if err != nil {
		return nil, fmt.Errorf("error creating flate writer: %w", err)
	}

	_, err = writer.Write(data)
	if err != nil {
		return nil, fmt.Errorf("error writing data: %w", err)
	}

	err = writer.Close()
	if err != nil {
		return nil, fmt.Errorf("error closing writer: %w", err)
	}

	return buf.Bytes(), nil
}

func Decompress(data []byte) ([]byte, error) {
	reader := flate.NewReader(bytes.NewReader(data))

	var buf bytes.Buffer

	_, err := buf.ReadFrom(reader)
	if err != nil {
		return nil, fmt.Errorf("error decompressing data: %w", err)
	}

	err = reader.Close()
	if err != nil {
		return nil, fmt.Errorf("error closing reader: %w", err)
	}

	return buf.Bytes(), nil
}

func main() {
	data := []byte("Hello World")
	compressed, err := Compress(data)
	if err != nil {
		log.Println(err)
	}

	decompressed, err := Decompress(compressed)
	if err != nil {
		log.Println(err)
	}

	log.Println(string(data))
	log.Println(string(decompressed))
}
