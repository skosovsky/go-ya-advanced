package services

import (
	"fmt"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"

	"mocks/internal/core/model"
	"mocks/internal/core/services/internal/mocks"
)

//go:generate mockgen -source=./repo.go -destination=internal/mocks/repo_mock.gen.go -package=mocks
// Или: go:generate mockgen -destination=internal/mocks-reflect/repo_mock.gen.go -package=mocks . Store,Foobar

func TestGetFoobar(t *testing.T) {
	cases := []struct {
		Name string

		Req *model.FoobarRequest

		StoreGetResponse *model.FoobarResponse
		StoreGetError    error

		FoobarCalculateResponse *model.FoobarResponse

		StoreSetError error

		ExpectedResponse *model.FoobarResponse
		ExpectedError    error
	}{
		{
			Name: "normal",
			Req: &model.FoobarRequest{
				N: 3,
			},
			StoreGetResponse: nil,
			StoreGetError:    nil,

			FoobarCalculateResponse: &model.FoobarResponse{
				Data: []string{"1", "2", "foo"},
			},

			StoreSetError: nil,

			ExpectedResponse: &model.FoobarResponse{
				Data: []string{"1", "2", "foo"},
			},
			ExpectedError: nil,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("test case #%d: %s", i, tc.Name), func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			storeMock := mocks.NewMockStore(ctrl)
			storeGet := storeMock.EXPECT().
				GetFoobar(tc.Req).
				Times(1).
				Return(tc.StoreGetResponse, tc.StoreGetError)

			foobarMock := mocks.NewMockFoobar(ctrl)
			if tc.StoreGetError == nil && tc.StoreGetResponse == nil {
				foobarCalculate := foobarMock.EXPECT().
					Calculate(tc.Req).
					After(storeGet).
					Times(1).
					Return(tc.FoobarCalculateResponse)

				_ = storeMock.EXPECT().
					SetFoobar(tc.Req, tc.FoobarCalculateResponse).
					After(foobarCalculate).
					Times(1).
					Return(tc.StoreSetError)
			}

			repo := NewRepo(storeMock)
			repo.foobar = foobarMock

			resp, err := repo.GetFoobar(tc.Req)
			if err := compareErrors(tc.ExpectedError, err); err != nil {
				t.Error(err)
				return
			}
			if !reflect.DeepEqual(resp, tc.ExpectedResponse) {
				t.Errorf("expected foobar response %v, got %v", tc.ExpectedResponse, resp)
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
