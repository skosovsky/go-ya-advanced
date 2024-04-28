package main

import (
	"log"
	"net/http"
	"time"
)

func redirect(resp http.ResponseWriter, req *http.Request) {
	http.Redirect(resp, req, "/", http.StatusMovedPermanently)
}

func longHandler(w http.ResponseWriter, _ *http.Request) {
	const timeout = 10 * time.Second
	time.Sleep(timeout)
	_, err := w.Write([]byte("Request took longer than 3 seconds"))
	if err != nil {
		log.Println(err)
		return
	}
}

func mainPage(res http.ResponseWriter, _ *http.Request) {
	_, err := res.Write([]byte("Hello World!"))
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", mainPage)

	http.HandleFunc("/search/", redirect)

	http.Handle("/dummy/", http.RedirectHandler("/", http.StatusMovedPermanently))

	http.Handle("/404/", http.NotFoundHandler())

	const timeout = 5 * time.Second
	http.Handle("/timeout/", http.TimeoutHandler(http.HandlerFunc(longHandler), timeout, "Timeout occurred"))

	log.Fatal(http.ListenAndServe("localhost:8080", nil)) //nolint:gosec // it's learning code
}
