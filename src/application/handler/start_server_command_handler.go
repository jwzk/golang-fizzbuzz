package handler

import (
	"fizzbuzz/application/command"
	"fizzbuzz/driver"
)

func StartServerCommandHandler(command command.StartServerCommand, server driver.Server) error {
	err := server.Start()
	if err != nil {
		return err
	}

	return nil
}
