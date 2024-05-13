package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("../.."))

	mux.Handle("/golang/", http.StripPrefix("/golang/", fileServer))
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "./main.go")
	})

	err := http.ListenAndServe("localhost:8080", mux) //nolint:gosec // it's learning code
	if err != nil {
		panic(err)
	}
}
