package server

import (
	"fizzbuzz/driver"
	"fmt"
	"net/http"
)

type HTTP struct {
	Port   uint16
	Router driver.Router
}

func New(port uint16, router driver.Router) *HTTP {
	return &HTTP{Port: port, Router: router}
}

func (server HTTP) Start() error {
	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		server.Router(writer, req)
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", server.Port), nil)
	if err != nil {
		return err
	}

	return nil
}
