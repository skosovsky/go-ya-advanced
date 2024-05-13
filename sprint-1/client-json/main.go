package main

import (
	"bytes"
	"log"
	"net/http"
)

func main() {
	var body = []byte(`{"message": "hello world"}`)

	request, err := http.NewRequest("POST", "http://localhost:8080/", bytes.NewBuffer(body)) //nolint:noctx,usestdlibvars // example
	if err != nil {
		log.Println(err)
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	contentType := request.Header.Values("Content-Type")
	request.Header.Set("Content-Type-my", "test") //nolint:canonicalheader // test
	log.Println(contentType)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Add("Content-Type", "charset=utf-8")
	contentType = request.Header.Values("Content-Type")
	log.Println(contentType)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()
}
