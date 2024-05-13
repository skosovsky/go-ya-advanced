package main

import (
	"log"
	"net/http"
	"net/http/cookiejar"
)

func main() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Jar: jar,
	}

	request, err := http.NewRequest(http.MethodGet, "https://www.google.com", http.NoBody) //nolint:noctx // example
	if err != nil {
		log.Fatal(err)
	}

	cookie := &http.Cookie{
		Name:   "Token",
		Value:  "TEST_TOKEN",
		MaxAge: 300, //nolint:mnd // example
	}

	request.AddCookie(cookie)
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
}
