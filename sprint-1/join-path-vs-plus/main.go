package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	if urlJoinPlus() == urlJoinPath() {
		log.Println("OK")
	}

	log.Println(urlJoinPathResult())
}

func urlJoinPlus() string {
	str := "http://" + "localhost:8080" + "/update" + "/gauge/" + "gauge" + "/"

	return str
}

func urlJoinPath() string {
	str, _ := url.JoinPath("http://localhost:8080", "/update", "/gauge/", "gauge", "/")

	return str
}

func urlJoinSprintf() string {
	str := fmt.Sprintf("http://%s/update/gauge/%v/%v", "localhost:8080", "gauge", "6.54")

	return str
}

func urlJoinPathResult() string {
	str, _ := url.JoinPath("localhost:8080", "update", "gauge", "gauge", "5.67")

	return str
}
