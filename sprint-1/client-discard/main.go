package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	response, err := http.Get("https://yandex.ru") //nolint:noctx // example
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()

	_, err = io.Copy(io.Discard, response.Body)
	if err != nil {
		log.Println(err)
	}
}
