package main

import (
	"net/http"
)

func mainPage(res http.ResponseWriter, _ *http.Request) {
	_, err := res.Write([]byte("Hello World!"))
	if err != nil {
		return
	}
}

func apiPage(res http.ResponseWriter, _ *http.Request) {
	_, err := res.Write([]byte("This is api page."))
	if err != nil {
		return
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", mainPage)
	mux.HandleFunc("/api", apiPage)

	err := http.ListenAndServe("localhost:8080", mux) //nolint:gosec // it's learning code
	if err != nil {
		panic(err)
	}
}
