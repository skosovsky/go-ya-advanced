package main

import (
	"compress/zlib"
	"io"
	"log"
	"net/http"
	"strings"
)

type zlibResponseWriter struct {
	http.ResponseWriter
	writer io.Writer
}

func (z zlibResponseWriter) Write(data []byte) (int, error) {
	return z.writer.Write(data) //nolint:wrapcheck // wrap func
}

func zlibHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !isCompressSupport("deflate", r.Header.Values("Accept-Encoding")) || !needCompress(r.Header.Values("Content-Type")) {
			next.ServeHTTP(w, r)

			return
		}

		zlibWriterBestSpeed, err := zlib.NewWriterLevel(w, zlib.BestSpeed)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

			return
		}

		defer func(zlibWriter *zlib.Writer) {
			err = zlibWriter.Close()
			if err != nil {
				log.Println(err)
			}
		}(zlibWriterBestSpeed)

		w.Header().Set("Content-Encoding", "deflate")

		gzipWriter := zlibResponseWriter{
			ResponseWriter: w,
			writer:         zlibWriterBestSpeed,
		}

		next.ServeHTTP(&gzipWriter, r)
	})
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

func defaultHandle(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	_, err := io.WriteString(w, "<html><body>"+strings.Repeat("Hello, world<br>", 20)+"</body></html>") //nolint:mnd // example

	if err != nil {
		log.Println(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultHandle)

	if err := http.ListenAndServe("localhost:8080", zlibHandle(mux)); err != nil { //nolint:gosec // example
		log.Fatal(err)
	}
}
