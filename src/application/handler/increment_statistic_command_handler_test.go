package handler

import (
	"errors"
	"fizzbuzz/application/command"
	"fizzbuzz/driver"
	"math"
	"net/http"
	"testing"
)

func TestIncrementStatisticCommandHandler(t *testing.T) {
	command := command.IncrementStatisticCommand{}
	storeMock := driver.StoreMock{}
	storeErrorMock := driver.StoreErrorMock{}
	req := http.Request{}

	inputs := []struct {
		hash     interface{}
		store    driver.Store
		expected error
	}{
		{"valid-hash", storeMock, nil},
		{"valid-hash", storeErrorMock, errors.New("store error")},
		{math.Inf(1), storeErrorMock, errors.New("json: unsupported value: +Inf")},
	}

	for _, input := range inputs {
		command.Hash = input.hash
		result := IncrementStatisticCommandHandler(&req, command, input.store)
		if result == nil && result != input.expected {
			t.Errorf("'%v' should be '%v'", result, input.expected)
		}
		if result != nil && result.Error() != input.expected.Error() {
			t.Errorf("'%v' should be '%v'", result, input.expected)
		}
	}
}
