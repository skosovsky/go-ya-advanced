package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func WriteHandleWriteString(w http.ResponseWriter, _ *http.Request) {
	_, err := io.WriteString(w, "test string")
	if err != nil {
		return
	}
}

func WriteHandleFprint(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprint(w, "test string")
	if err != nil {
		return
	}
}

func WriteHandleWrite(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("test string"))
	if err != nil {
		return
	}
}

func main() {
	log.Println()
}
