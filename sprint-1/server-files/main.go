package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static")) // доступны все файлы из указанной папки

	// конкретный файл из папки - localhost:8080/test.txt
	mux.Handle("/", fileServer)

	// конкретный файл из папки - localhost:8080/files/test.txt
	mux.Handle("/files/", http.StripPrefix("/files/", fileServer)) // заменяем путь на files

	// только конкретный файл
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/test.txt")
	})

	err := http.ListenAndServe("localhost:8080", mux) //nolint:gosec // it's learning code
	if err != nil {
		panic(err)
	}
}
