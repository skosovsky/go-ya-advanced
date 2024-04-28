package main

import "net/http"

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// тут пишем обработку
		// например, разрешаем запросы cross-domain
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(res, req)
	})
}

func rootHandler(res http.ResponseWriter, _ *http.Request) {
	_, err := res.Write([]byte("Hello World"))
	if err != nil {
		return
	}
}

func main() {
	http.Handle("/", middleware(http.HandlerFunc(rootHandler)))

	err := http.ListenAndServe("localhost:8080", nil) //nolint:gosec // it's learning code
	if err != nil {
		panic(err)
	}
}
