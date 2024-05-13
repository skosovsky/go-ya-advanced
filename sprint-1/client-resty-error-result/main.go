package main

import (
	"log"
	"time"

	"github.com/go-resty/resty/v2"
)

type APIError struct {
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Text   string `json:"text"`
}

func main() {
	client := resty.New()

	var responseErr APIError
	var post Post

	_, err := client.R().
		SetError(&responseErr).
		SetResult(&post).
		Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		log.Println(responseErr)
		log.Println(err)

		return
	}

	log.Println(post)
}
