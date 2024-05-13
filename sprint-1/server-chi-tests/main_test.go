package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testRequest(t *testing.T, testServer *httptest.Server, method, path string) (*http.Response, string) { // internal func
	t.Helper()

	request, err := http.NewRequest(method, testServer.URL+path, nil)
	require.NoError(t, err)

	response, err := testServer.Client().Do(request)
	require.NoError(t, err)
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	return response, string(responseBody)
}

func TestRouter(t *testing.T) {
	t.Parallel()

	testServer := httptest.NewServer(CarRouter())
	t.Cleanup(testServer.Close)

	testCases := []struct {
		url    string
		want   string
		status int
	}{
		{"/cars/renault/Logan", "Renault Logan", http.StatusOK},
		{"/cars/audi/a4", "Audi A4", http.StatusOK},
		{"/cars/audi/a6", "unknown model: audi a6\n", http.StatusNotFound}, // проверим на ошибочный запрос
		{"/cars/BMW/M5", "BMW M5", http.StatusOK},
		{"/cars/bmw/X6", "BMW X6", http.StatusOK},
		{"/cars/Vw/Passat", "VW Passat", http.StatusOK},
	}

	for _, tt := range testCases {
		response, responseBody := testRequest(t, testServer, http.MethodGet, tt.url)
		assert.Equal(t, tt.status, response.StatusCode)
		assert.Equal(t, tt.want, responseBody)
	}
}
