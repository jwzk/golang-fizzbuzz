package handler

import (
	"errors"
	"fizzbuzz/application/command"
	"fizzbuzz/driver"
	"net/http"
	"testing"
)

func TestGETBestStatisticCommandHandler(t *testing.T) {
	command := command.GETBestStatisticCommand{}
	storeMock := driver.StoreMock{}
	storeErrorMock := driver.StoreErrorMock{}
	storeInvalidMock := driver.StoreInvalidMock{}
	// request := driver.Request{}
	req := http.Request{}

	inputs := []struct {
		store         driver.Store
		expectedValue BestStructRes
		expectedError error
	}{
		{storeMock, BestStructRes{Hits: 4, Parameters: "example-1"}, nil},
		{storeErrorMock, BestStructRes{}, errors.New("store error")},
		{storeInvalidMock, BestStructRes{}, nil},
	}

	for _, input := range inputs {
		result, err := GETBestStatisticCommandHandler(&req, command, input.store)

		if result != input.expectedValue {
			t.Errorf("'%v' should be '%v'", result, input.expectedValue)
		}
		if err == nil && err != input.expectedError {
			t.Errorf("'%v' should be '%v'", err, input.expectedError)
		}
		if err != nil && err.Error() != input.expectedError.Error() {
			t.Errorf("'%v' should be '%v'", err, input.expectedError)
		}
	}
}
