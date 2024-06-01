package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var buf bytes.Buffer

	logger := log.New(&buf, "myLog: ", 0)
	logger.Println("Hello World")
	logger.Println("Goodbye World")

	fmt.Println(&buf) //nolint:forbidigo // example
}
