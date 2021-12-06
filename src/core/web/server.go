package web

import (
	"fizzbuzz/application/command"
	"fizzbuzz/application/handler"
	"fizzbuzz/driver"
	"fizzbuzz/infrastructure/server"
	"log"
	"os"
	"strconv"
)

func InitServer(router driver.Router) {
	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	command := command.StartServerCommand{}
	http := server.New(uint16(port), router)
	err = handler.StartServerCommandHandler(command, http)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
