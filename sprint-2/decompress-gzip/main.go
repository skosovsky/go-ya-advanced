package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func LengthHandle(w http.ResponseWriter, r *http.Request) {
	var reader io.ReadCloser

	if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		gzipReader, err := gzip.NewReader(r.Body)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

			return
		}

		reader = gzipReader

		defer func(gzipReader *gzip.Reader) {
			err = gzipReader.Close()
			if err != nil {
				log.Println(err)
			}
		}(gzipReader)
	} else {
		reader = r.Body
	}

	body, err := io.ReadAll(reader)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	_, err = fmt.Fprintf(w, "Length: %d", len(body))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}
}

func main() {

}
