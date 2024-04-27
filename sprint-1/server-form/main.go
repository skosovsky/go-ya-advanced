package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func mainPage(res http.ResponseWriter, req *http.Request) {
	body := fmt.Sprintf("Method: %s\r\n", req.Method)
	body += "Header ===\r\n"
	for key, value := range req.Header {
		body += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	body += "Query parameters from req.URL.Query() ===\r\n"

	for key, value := range req.URL.Query() {
		body += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	bodyNew, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(bodyNew)) // empty for GET

	if err = req.ParseForm(); err != nil {
		_, err = res.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
			return
		}
	}

	body += "Query parameters from req.Form ===\r\n"

	for key, value := range req.Form {
		body += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	_, err = res.Write([]byte(body))
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
