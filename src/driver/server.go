package driver

import (
	"errors"
	"net/http"
)

const (
	HTTPOk                  = http.StatusOK
	HTTPBadRequest          = http.StatusBadRequest
	HTTPNotFound            = http.StatusNotFound
	HTTPInternalServerError = http.StatusInternalServerError
	BadRequestText          = "Bad Request"
	NotFoundText            = "Not Found"
	InternalServorErrorText = "Internal Server Error"
	GETMethod               = "GET"
	POSTMethod              = "POST"
)

// type Request http.Request

type Writer interface {
	Header() http.Header
	Write(p []byte) (n int, err error)
	WriteHeader(statusCode int)
}

type Router func(Writer, *http.Request)

type Server interface {
	Start() error
}

// Mock Part

type ServerMock struct {
}

func (server ServerMock) Start() error {
	return nil
}

type ServerErrorMock struct {
}

func (server ServerErrorMock) Start() error {
	return errors.New("server error")
}
