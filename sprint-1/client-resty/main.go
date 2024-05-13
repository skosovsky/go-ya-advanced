package main

import (
	"log"

	"github.com/go-resty/resty/v2"
)

func main() {
	client := resty.New()

	response, err := client.R().
		SetAuthToken("Bearer TOKEN").
		Get("https://google.com")

	if err != nil {
		log.Println(err)

		return
	}

	log.Println(response.StatusCode())
	log.Println(response.Status())
	log.Println(response.Header())
	log.Println(response.Time())
	log.Println(response.ReceivedAt())
	log.Println(response.RawBody())
	log.Println(response)
}
