package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStatusHandler(t *testing.T) {
	t.Parallel()

	type want struct {
		code        int
		response    string
		contentType string
	}

	testCases := []struct {
		name string
		want want
	}{
		{
			name: "positive test 1",
			want: want{
				code:        200,
				response:    `{"status":"ok"}`,
				contentType: "application/json",
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			request := httptest.NewRequest(http.MethodGet, "/status", nil)

			responseRecorder := httptest.NewRecorder()
			StatusHandler(responseRecorder, request)

			response := responseRecorder.Result()

			assert.Equal(t, tt.want.code, response.StatusCode)

			responseBody, err := io.ReadAll(response.Body)
			defer response.Body.Close()

			require.NoError(t, err)

			assert.JSONEq(t, tt.want.response, string(responseBody))
			assert.Equal(t, tt.want.contentType, response.Header.Get("Content-Type"))
		})
	}
}
