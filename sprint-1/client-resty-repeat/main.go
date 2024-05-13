package main

import (
	"log"
	"time"

	"github.com/go-resty/resty/v2"
)

func main() {
	client := resty.New()

	client.
		SetRetryCount(3).                     //nolint:mnd // example
		SetRetryWaitTime(5 * time.Second).    //nolint:mnd // example
		SetRetryMaxWaitTime(20 * time.Second) //nolint:mnd // example

	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"title":"foo", "body":"bar", "userId": 7}`).
		Post("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		log.Println(err)

		return
	}

	log.Println(response)

	response, err = client.R(). // Другой вариант POST-запроса, при map используется JSON по-умолчанию
					SetBody(map[string]interface{}{"title": "My title", "body": "Content", "userId": 7}). //nolint:mnd // example
					Post("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		log.Println(err)

		return
	}

	log.Println(response)
}
