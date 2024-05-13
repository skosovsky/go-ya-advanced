package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func TimerTrace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// перед началом выполнения функции сохраняем текущее время
		start := time.Now()
		// вызываем следующий обработчик
		next.ServeHTTP(w, r)
		// после завершения замеряем время выполнения запроса
		duration := time.Since(start)
		// выводим обрабатываем полученный результат
		log.Printf("%s %s %s", r.Method, r.URL.Path, duration)
	})
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RealIP) // или r.Use(middleware.RealIP, middleware.Logger, middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(TimerTrace)
}
