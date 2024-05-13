package main

import (
	"log"

	"github.com/go-resty/resty/v2"
)

func main() {
	client := resty.New()

	response, err := client.R().
		SetPathParams(map[string]string{"postID": "1"}).
		Get("https://jsonplaceholder.typicode.com/posts/{postID}")

	if err != nil {
		log.Println(err)

		return
	}

	log.Println(response)
}
