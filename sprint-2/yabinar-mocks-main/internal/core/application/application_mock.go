package application

import (
	"testing"

	"mocks/internal/core/model"
)

type repoMock struct {
	t *testing.T

	expectedRequest *model.FoobarRequest

	mockResponse *model.FoobarResponse
	mockError    error
}

func newRepoMock(
	t *testing.T,
	expectedRequest *model.FoobarRequest,
	mockResponse *model.FoobarResponse,
	mockError error,
) *repoMock {
	return &repoMock{
		t:               t,
		expectedRequest: expectedRequest,
		mockResponse:    mockResponse,
		mockError:       mockError,
	}
}

func (r *repoMock) GetFoobar(req *model.FoobarRequest) (*model.FoobarResponse, error) {
	if *req != *r.expectedRequest {
		r.t.Errorf("expected request %v, got: %v", r.expectedRequest, req)
	}
	return r.mockResponse, r.mockError
}
