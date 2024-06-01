package main

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWebhook(t *testing.T) {
	t.Parallel()

	handler := http.HandlerFunc(webhook)
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)

	successBody := `{"response": {"text": "Sorry, I'm noob"}, "version": "1.0"}`

	testCases := []struct {
		name     string
		method   string
		body     string
		wantCode int
		wantBody string
	}{
		{
			name:     "Get method",
			method:   http.MethodGet,
			body:     "",
			wantCode: http.StatusMethodNotAllowed,
			wantBody: "",
		},
		{
			name:     "Put method",
			method:   http.MethodPut,
			body:     "",
			wantCode: http.StatusMethodNotAllowed,
			wantBody: "",
		},
		{
			name:     "Delete method",
			method:   http.MethodDelete,
			body:     "",
			wantCode: http.StatusMethodNotAllowed,
			wantBody: "",
		},
		{
			name:     "Post method without body",
			method:   http.MethodPost,
			body:     "",
			wantCode: http.StatusInternalServerError,
			wantBody: "",
		},
		{
			name:     "Post method unsupported type",
			method:   http.MethodPost,
			body:     `{"request": {"type": "idunno", "command": "do something"}, "version": "1.0"}`,
			wantCode: http.StatusUnprocessableEntity,
			wantBody: "",
		},
		{
			name:     "Post method success",
			method:   http.MethodPost,
			body:     `{"request": {"type": "SimpleUtterance", "command": "sudo do something"}, "version": "1.0"}`,
			wantCode: http.StatusOK,
			wantBody: successBody,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.method, func(t *testing.T) {
			t.Parallel()

			request := resty.New().R()
			request.Method = tt.method
			request.URL = server.URL

			if len(tt.body) > 0 {
				request.SetHeader("Content-Type", "application/json")
				request.SetBody(tt.body)
			}

			response, err := request.Send()
			require.NoError(t, err, "error making request")

			assert.Equal(t, tt.wantCode, response.StatusCode(), "Response code not match")
			if tt.wantBody != "" {
				assert.JSONEq(t, tt.wantBody, string(response.Body()), "Response body not match")
			}
		})
	}
}

func TestGzipCompression(t *testing.T) { //nolint:tparallel // not work
	t.Parallel()

	handler := gzipMiddleware(webhook)

	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)

	requestBody := `{"request": {"type": "SimpleUtterance", "command": "sudo do something"}, "version": "1.0"}`
	successBody := `{"response": {"text": "Sorry, I'm noob"}, "version": "1.0"}`

	t.Run("sends_gzip", func(t *testing.T) { //nolint:paralleltest // not work
		buf := bytes.NewBuffer(nil)
		gzipWriter := gzip.NewWriter(buf)
		_, err := gzipWriter.Write([]byte(requestBody))
		require.NoError(t, err)
		err = gzipWriter.Close()
		require.NoError(t, err)

		r := httptest.NewRequest(http.MethodPost, server.URL, buf)
		r.RequestURI = ""
		r.Header.Set("Content-Encoding", "gzip")
		r.Header.Set("Accept-Encoding", "")

		response, err := http.DefaultClient.Do(r)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode)

		defer response.Body.Close()

		responseBody, err := io.ReadAll(response.Body)
		require.NoError(t, err)
		require.JSONEq(t, successBody, string(responseBody))
	})

	t.Run("accepts_gzip", func(t *testing.T) { //nolint:paralleltest // not work
		buf := bytes.NewBufferString(requestBody)
		request := httptest.NewRequest(http.MethodPost, server.URL, buf)
		request.RequestURI = ""
		request.Header.Set("Accept-Encoding", "gzip")

		response, err := http.DefaultClient.Do(request)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode)

		defer response.Body.Close()

		gzipReader, err := gzip.NewReader(response.Body)
		require.NoError(t, err)

		gzipData, err := io.ReadAll(gzipReader)
		require.NoError(t, err)

		require.JSONEq(t, successBody, string(gzipData))
	})
}
