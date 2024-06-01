package application

import (
	"fmt"
	"reflect"
	"testing"

	"mocks/internal/core/model"
)

func TestGetFoobar(t *testing.T) {
	cases := []struct {
		name string

		req                  *model.FoobarRequest
		expectedRepoResponse *model.FoobarResponse
		expectedRepoError    error
		expectedResponse     *model.FoobarResponse
		expectedError        error
	}{
		{
			name: "normal",
			req: &model.FoobarRequest{
				N: 3,
			},
			expectedRepoResponse: &model.FoobarResponse{
				Data: []string{"1", "2", "foo"},
			},
			expectedRepoError: nil,
			expectedResponse: &model.FoobarResponse{
				Data: []string{"1", "2", "foo"},
			},
		},
		{
			name:          "nil request",
			req:           nil,
			expectedError: fmt.Errorf("passed foobar request is nil"),
		},
		{
			name: "bad N",
			req: &model.FoobarRequest{
				N: -42,
			},
			expectedError: fmt.Errorf("expected parameter N of the foobar request to be > 0, got -42"),
		},
	}

	for i, tc := range cases {
		i, tc := i, tc

		t.Run(fmt.Sprintf("test case #%d: %s", i, tc.name), func(t *testing.T) {
			mockRepo := newRepoMock(t, tc.req, tc.expectedRepoResponse, tc.expectedError)
			app := NewApplication(mockRepo)

			actualResponse, actualError := app.GetFoobar(tc.req)
			if err := compareErrors(tc.expectedError, actualError); err != nil {
				t.Error(err)
				return
			}
			if !reflect.DeepEqual(tc.expectedResponse, actualResponse) {
				t.Errorf("expected foobar response %v, got %v", tc.expectedResponse, actualResponse)
				return
			}
		})
	}

}

func compareErrors(expected error, actual error) error {
	if expected == nil && actual == nil {
		return nil
	}
	if expected == nil {
		return fmt.Errorf("expected a nil error, got: %v", actual)
	}
	if actual == nil {
		return fmt.Errorf("expected an error: %v, got a nil", expected)
	}
	if expected.Error() != actual.Error() {
		return fmt.Errorf("expected an error %v, got %v", expected, actual)
	}
	return nil
}
