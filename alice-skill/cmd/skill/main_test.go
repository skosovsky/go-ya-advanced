package main

import (
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

	successBody := `{
        "response": {
            "text": "Sorry, I'm noob"
        },
        "version": "1.0"
    }`

	testCases := []struct {
		method   string
		wantCode int
		wantBody string
	}{
		{
			method:   http.MethodGet,
			wantCode: http.StatusMethodNotAllowed,
			wantBody: "",
		},
		{
			method:   http.MethodPut,
			wantCode: http.StatusMethodNotAllowed,
			wantBody: "",
		},
		{
			method:   http.MethodDelete,
			wantCode: http.StatusMethodNotAllowed,
			wantBody: "",
		},
		{
			method:   http.MethodPost,
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

			response, err := request.Send()
			require.NoError(t, err, "error making request")

			assert.Equal(t, tt.wantCode, response.StatusCode(), "Response code not match")
			if tt.wantBody != "" {
				assert.JSONEq(t, tt.wantBody, string(response.Body()), "Response body not match")
			}
		})
	}
}
