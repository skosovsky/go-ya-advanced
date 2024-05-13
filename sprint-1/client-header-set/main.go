package main

import (
	"log"
	"net/http"
)

func main() {
	request, err := http.NewRequest(http.MethodGet, `http://localhost:8080`, nil) //nolint:noctx // example
	if err != nil {
		log.Println(err)
	}

	request.Header.Set("User-Agent", "go-client")
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()
}
