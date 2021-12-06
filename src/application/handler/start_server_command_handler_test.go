package handler

import (
	"errors"
	"fizzbuzz/application/command"
	"fizzbuzz/driver"
	"testing"
)

func TestStartServerCommandHandler(t *testing.T) {
	command := command.StartServerCommand{}
	serverMock := driver.ServerMock{}
	serverErrorMock := driver.ServerErrorMock{}

	inputs := []struct {
		server   driver.Server
		expected error
	}{
		{serverMock, nil},
		{serverErrorMock, errors.New("server error")},
	}

	for _, input := range inputs {
		result := StartServerCommandHandler(command, input.server)
		if result == nil && result != input.expected {
			t.Errorf("'%v' should be '%v'", result, input.expected)
		}
		if result != nil && result.Error() != input.expected.Error() {
			t.Errorf("'%v' should be '%v'", result, input.expected)
		}
	}
}
