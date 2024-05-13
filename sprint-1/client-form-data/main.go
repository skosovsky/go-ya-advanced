package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	fileName := "test.txt"

	file, _ := os.Open(fileName)
	defer file.Close()

	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("upload_file", fileName)
	if err != nil {
		log.Println(err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		log.Println(err)
	}
	writer.Close()

	request, err := http.NewRequest(http.MethodPost, "http://localhost:8080/upload", body) //nolint:noctx // example
	if err != nil {
		log.Println(err)
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()
}
