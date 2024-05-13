package main

import (
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func main() {
	data := url.Values{}
	data.Set("key1", "value1")
	data.Set("key2", "value2")

	request, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8080/", strings.NewReader(data.Encode())) //nolint:noctx // example
	if err != nil {
		log.Println(err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Content-Length", strconv.Itoa(len(data.Encode())))
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()
}
