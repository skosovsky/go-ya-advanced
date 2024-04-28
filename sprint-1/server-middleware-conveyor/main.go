package main

import (
	"log"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func Conveyor(hand http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		hand = middleware(hand)
	}

	return hand
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// тут пишем обработку
		next.ServeHTTP(res, req)
	})
}

func middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// тут пишем обработку
		next.ServeHTTP(res, req)
	})
}

func middleware3(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// тут пишем обработку
		next.ServeHTTP(res, req)
	})
}

func rootHandle(resp http.ResponseWriter, _ *http.Request) {
	_, err := resp.Write([]byte("Hello World"))
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	http.Handle("/", Conveyor(http.HandlerFunc(rootHandle), middleware, middleware2, middleware3))
	log.Fatal(http.ListenAndServe("localhost:8080", nil)) //nolint:gosec // it's learning code
}
