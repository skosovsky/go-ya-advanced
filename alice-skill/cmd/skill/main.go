package main

import (
	"alice-skill/internal/logger"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func gzipMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		originalWriter := w

		supportsGzip := isCompressSupport("gzip", r.Header.Values("Accept-Encoding"))
		needCompressWriter := needCompress(r.Header.Values("Content-Type"))
		if supportsGzip && needCompressWriter {
			compressWriter := NewCompressWriter(w)

			originalWriter = compressWriter

			defer func(gzipWriter *gzip.Writer) {
				err := gzipWriter.Close()
				if err != nil {
					logger.Log.Error(err.Error())
				}
			}(compressWriter.gzipWriter)
		}

		contentEncoding := r.Header.Get("Content-Encoding")
		sendsGzip := strings.Contains(contentEncoding, "gzip")
		if sendsGzip {
			compressReader, err := NewCompressReader(r.Body)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

				return
			}

			r.Body = compressReader

			defer func(compressReader *CompressReader) {
				err = compressReader.Close()
				if err != nil {
					logger.Log.Error(err.Error())
				}
			}(compressReader)
		}

		next.ServeHTTP(originalWriter, r)
	}
}

func isCompressSupport(format string, acceptEncodes []string) bool {
	for i := range acceptEncodes {
		if strings.Contains(acceptEncodes[i], format) {
			return true
		}
	}

	return false
}

func needCompress(contentTypes []string) bool {
	acceptContentTypes := map[string]bool{
		"application/javascript": true,
		"application/json":       true,
		"text/css":               true,
		"text/html":              true,
		"text/plain":             true,
		"text/xml":               true,
	}

	var contentTypesSplits []string

	for i := range contentTypes {
		contentTypesSplits = append(contentTypesSplits, strings.Split(contentTypes[i], ";")...)
	}

	for j := range contentTypesSplits {
		if acceptContentTypes[contentTypesSplits[j]] {
			return true
		}
	}

	return false
}

func webhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logger.Log.Debug("got request with bad method",
			zap.String("method", r.Method))

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

		return
	}

	logger.Log.Debug("decoding request")
	var request Request
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		logger.Log.Error("failed to decode request",
			zap.Error(err))

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	if request.Request.Type != TypeSimpleUtterance {
		logger.Log.Debug("unsupported request type",
			zap.String("type", request.Request.Type))

		http.Error(w, "unsupported request type", http.StatusUnprocessableEntity)

		return
	}

	response := Response{
		Response: ResponsePayload{
			Text: "Sorry, I'm noob",
		},
		Version: "1.0",
	}

	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		logger.Log.Error("failed to encode response",
			zap.Error(err))

		return
	}

	logger.Log.Debug("sending HTTP 200 response")
}

func run() error {
	if err := logger.Initialize(flagLogLevel); err != nil {
		return fmt.Errorf("failed to initialize logger: %w", err)
	}

	logger.Log.Info("Running server",
		zap.String("address", flagRunAddr))

	err := http.ListenAndServe(flagRunAddr, logger.RequestLogger(gzipMiddleware(webhook))) //nolint:gosec // it's example
	if err != nil {
		return fmt.Errorf("could not start server: %w", err)
	}

	return nil
}

func main() {
	parseFlags()

	if err := run(); err != nil {
		panic(err)
	}
}
