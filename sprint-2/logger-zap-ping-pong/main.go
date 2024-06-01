package main

import (
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

var sugar zap.SugaredLogger //nolint:gochecknoglobals // example

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync() //nolint:errcheck // example

	sugar = *logger.Sugar()

	http.Handle("/ping", WithLogging(pingHandler()))

	addr := "localhost:8080"

	sugar.Infow("Starting server",
		"addr", addr)
	if err = http.ListenAndServe(addr, nil); err != nil { //nolint:gosec // example
		sugar.Fatal(err.Error(),
			"event", "start server failed")
	}
}

func pingHandler() http.Handler { // curl --include localhost:8080/ping
	fn := func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("pong\n"))
	}

	return http.HandlerFunc(fn)
}

func WithLogging(next http.Handler) http.Handler {
	logFunc := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		responseData := &ResponseData{
			status: 0,
			size:   0,
		}

		logWriter := loggingResponseWriter{
			ResponseWriter: w,
			responseData:   responseData,
		}

		next.ServeHTTP(&logWriter, r)

		uri := r.RequestURI
		method := r.Method
		duration := time.Since(start)

		sugar.Infoln(
			"uri", uri,
			"method", method,
			"status", responseData.status,
			"duration", duration,
			"size", responseData.size,
		)
	}

	return http.HandlerFunc(logFunc)
}

type (
	ResponseData struct {
		status int
		size   int
	}

	loggingResponseWriter struct {
		http.ResponseWriter
		responseData *ResponseData
	}
)

func (r *loggingResponseWriter) Write(data []byte) (int, error) {
	size, err := r.ResponseWriter.Write(data)
	if err != nil {
		return 0, fmt.Errorf("failed to write data to response, %w", err)
	}

	r.responseData.size += size

	return size, nil
}

func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode
}
