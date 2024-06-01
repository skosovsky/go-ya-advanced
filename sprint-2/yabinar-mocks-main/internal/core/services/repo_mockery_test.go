package services

import (
	"fmt"
	"mocks/internal/core/model"
	mocks "mocks/internal/core/services/internal/mockery"
	"reflect"
	"testing"
)

func TestGetFoobarMockery(t *testing.T) {
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
		i, tc := i, tc

		t.Run(fmt.Sprintf("test case #%d: %s", i, tc.Name), func(t *testing.T) {

			storeMock := mocks.NewMockStore(t)

			storeGet := storeMock.EXPECT().
				GetFoobar(tc.Req).
				Times(1).
				Return(tc.StoreGetResponse, tc.StoreGetError)

			foobarMock := mocks.NewMockFoobar(t)

			if tc.StoreGetError == nil && tc.StoreGetResponse == nil {
				foobarCalculate := foobarMock.EXPECT().
					Calculate(tc.Req).
					NotBefore(storeGet).
					Times(1).
					Return(tc.FoobarCalculateResponse)

				_ = storeMock.EXPECT().
					SetFoobar(tc.Req, tc.FoobarCalculateResponse).
					NotBefore(foobarCalculate).
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
