package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, "http://localhost:8080", http.NoBody) //nolint:noctx // example
	if err != nil {
		log.Println(err)
	}

	request.AddCookie(&http.Cookie{
		Name:       "ID",
		Value:      "3675",
		Path:       "",
		Domain:     "",
		Expires:    time.Time{},
		RawExpires: "",
		MaxAge:     0,
		Secure:     false,
		HttpOnly:   false,
		SameSite:   0,
		Raw:        "",
		Unparsed:   nil,
	})
	request.AddCookie(&http.Cookie{
		Name:   "Token",
		Value:  "TEST_TOKEN",
		MaxAge: 360, //nolint:mnd // example
	})

	response, err := client.Do(request)
	if err != nil {
		log.Println(err)

		return
	}
	defer response.Body.Close()
}
