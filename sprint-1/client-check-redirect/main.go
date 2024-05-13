package main

import (
	"log"
	"net/http"
)

func main() {
	client := http.Client{
		CheckRedirect: func(r *http.Request, _ []*http.Request) error {
			log.Println(r.URL)

			return nil
		},
	}

	response, err := client.Get("https://yandex.ru") //nolint:noctx // example
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()
}
