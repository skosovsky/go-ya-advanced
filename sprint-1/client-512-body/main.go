package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	const limitBytes = 512

	log.Println(len(bodyLimitReader(limitBytes)))
	log.Println(len(bodyCopyN(limitBytes)))
	log.Println(len(bodyCuttingSlice(limitBytes)))
}

func bodyLimitReader(limit int64) string {
	response, err := http.Get("https://practicum.yandex.ru") //nolint:noctx // example
	if err != nil {
		log.Fatal(err)

		return ""
	}
	defer response.Body.Close()

	limitReader := io.LimitReader(response.Body, limit)

	cuttingBody := new(strings.Builder)

	if _, err = io.Copy(cuttingBody, limitReader); err != nil {
		log.Println(err)
	}

	return cuttingBody.String()
}

func bodyCopyN(limit int64) string {
	response, err := http.Get("https://practicum.yandex.ru") //nolint:noctx // example
	if err != nil {
		log.Println(err)

		return ""
	}
	defer response.Body.Close()

	cuttingBody := new(strings.Builder)

	if _, err = io.CopyN(cuttingBody, response.Body, limit); err != nil {
		log.Println(err)
	}

	return cuttingBody.String()
}

func bodyCuttingSlice(limit int64) string {
	response, err := http.Get("https://practicum.yandex.ru") //nolint:noctx // example
	if err != nil {
		log.Println(err)

		return ""
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)

		return ""
	}

	if len(body) > int(limit) {
		body = body[:limit]
	}

	return string(body)
}
