package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	endpoint := "http://localhost:8080/"

	data := url.Values{}

	var long string
	_, err := fmt.Scan(&long)
	if err != nil {
		log.Println(err)

		return
	}

	long = strings.TrimSpace(long)

	data.Set("url", long)
	client := &http.Client{}

	request, err := http.NewRequest(http.MethodPost, endpoint, strings.NewReader(data.Encode())) //nolint:noctx // example
	if err != nil {
		log.Println(err)

		return
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)

		return
	}
	defer response.Body.Close()

	log.Println(response.StatusCode)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)

		return
	}
	log.Println(string(body))
}
