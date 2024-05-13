package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserViewHandler(t *testing.T) {
	t.Parallel()

	type want struct {
		request     string
		code        int
		response    string
		contentType string
	}

	testCases := []struct {
		name string
		want want
	}{
		{
			name: "positive test user1",
			want: want{
				request:     "/users?user_id=user1",
				code:        200,
				response:    `{"firstName":"Misha", "id":"user1", "lastName":"Popov"}`,
				contentType: "application/json",
			},
		},
		{
			name: "negative test user not found",
			want: want{
				request:     "/users?user_id=user3",
				code:        404,
				response:    `user not found`,
				contentType: "text/plain; charset=utf-8",
			},
		},
	}

	users := make(map[string]User)
	user1 := User{
		ID:        "user1",
		FirstName: "Misha",
		LastName:  "Popov",
	}
	user2 := User{
		ID:        "user2",
		FirstName: "Sasha",
		LastName:  "Popov",
	}
	users["user1"] = user1
	users["user2"] = user2

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			request := httptest.NewRequest(http.MethodGet, tt.want.request, nil)
			responseRecorder := httptest.NewRecorder()
			UserViewHandler(users).ServeHTTP(responseRecorder, request)

			response := responseRecorder.Result()
			assert.Equal(t, tt.want.code, response.StatusCode)

			responseBody, err := io.ReadAll(response.Body)
			require.NoError(t, err)

			err = response.Body.Close()
			require.NoError(t, err)

			if response.Header.Get("Content-Type") == "application/json" {
				assert.JSONEq(t, tt.want.response, string(responseBody))
			}

			assert.Equal(t, tt.want.contentType, response.Header.Get("Content-Type"))
		})
	}
}

func TestUserViewHandlerByExample(t *testing.T) {
	t.Parallel()

	type want struct {
		contentType string
		statusCode  int
		user        User
	}
	testCases := []struct {
		name    string
		request string
		users   map[string]User
		want    want
	}{
		{
			name: "simple test #1",
			users: map[string]User{
				"id1": {
					ID:        "id1",
					FirstName: "Misha",
					LastName:  "Popov",
				},
			},
			want: want{
				contentType: "application/json",
				statusCode:  200,
				user: User{ID: "id1",
					FirstName: "Misha",
					LastName:  "Popov",
				},
			},
			request: "/users?user_id=id1",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			request := httptest.NewRequest(http.MethodGet, tt.request, nil)
			responseRecorder := httptest.NewRecorder()
			handler := UserViewHandler(tt.users)
			handler.ServeHTTP(responseRecorder, request)

			response := responseRecorder.Result()

			assert.Equal(t, tt.want.statusCode, response.StatusCode)
			assert.Equal(t, tt.want.contentType, response.Header.Get("Content-Type"))

			responseBody, err := io.ReadAll(response.Body)
			require.NoError(t, err)
			err = response.Body.Close()
			require.NoError(t, err)

			var user User
			err = json.Unmarshal(responseBody, &user)
			require.NoError(t, err)

			assert.Equal(t, tt.want.user, user)
		})
	}
}
