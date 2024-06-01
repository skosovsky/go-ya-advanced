package main

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type CompressWriter struct {
	http.ResponseWriter
	gzipWriter *gzip.Writer
	size       int
}

func NewCompressWriter(writer http.ResponseWriter) *CompressWriter {
	return &CompressWriter{
		ResponseWriter: writer,
		gzipWriter:     gzip.NewWriter(writer),
		size:           0,
	}
}

func (c *CompressWriter) Write(data []byte) (int, error) {
	size, err := c.gzipWriter.Write(data)
	if err != nil {
		return 0, err //nolint:wrapcheck // wrap func
	}

	c.size += size

	const limitNonCompressedData = 1400

	if c.size < limitNonCompressedData {
		c.Header().Del("Content-Encoding")
	}

	if c.size >= limitNonCompressedData {
		c.Header().Set("Content-Encoding", "gzip")
	}

	return size, nil
}

func (c *CompressWriter) WriteHeader(statusCode int) {
	const firstNegativeStatus = 300

	if statusCode < firstNegativeStatus {
		c.Header().Set("Content-Encoding", "gzip")
	}

	if statusCode >= firstNegativeStatus {
		c.Header().Del("Content-Encoding")
	}

	c.ResponseWriter.WriteHeader(statusCode)
}

type CompressReader struct {
	io.ReadCloser
	gzipReader *gzip.Reader
}

func NewCompressReader(reader io.ReadCloser) (*CompressReader, error) {
	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		return nil, fmt.Errorf("could not create gzip reader: %w", err)
	}

	return &CompressReader{
		ReadCloser: reader,
		gzipReader: gzipReader,
	}, nil
}

func (c *CompressReader) Close() error {
	var err error

	if errReader := c.ReadCloser.Close(); errReader != nil {
		err = errors.Join(err, errReader)
	}

	if errGzipReader := c.gzipReader.Close(); err != nil {
		err = errors.Join(err, errGzipReader)
	}

	if err != nil {
		return err
	}

	return nil
}
