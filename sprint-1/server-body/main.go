package main

import (
	"fmt"
	"log"
	"net/http"
)

func mainPage(res http.ResponseWriter, req *http.Request) {
	body := fmt.Sprintf("Method: %s\r\n", req.Method)
	body += "Header ===\r\n"
	for key, value := range req.Header {
		body += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	body += "Query parameters ===\r\n"
	for key, value := range req.URL.Query() {
		body += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	_, err := res.Write([]byte(body))
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainPage)

	err := http.ListenAndServe("localhost:8080", mux) //nolint:gosec // it's learning code
	if err != nil {
		panic(err)
	}
}
