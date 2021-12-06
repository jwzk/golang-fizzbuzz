package controller

import (
	"fizzbuzz/application/command"
	"fizzbuzz/application/handler"
	"fizzbuzz/driver"
	"fizzbuzz/infrastructure/store"
	"fizzbuzz/presentation/transformer"
	"fizzbuzz/presentation/validator"
	"log"
	"net/http"
)

const (
	fizzbuzzPath = "/fizzbuzz"
)

func handleStatistic(req *http.Request, commandStruct interface{}) {
	incrementStatisticCommand := command.IncrementStatisticCommand{Hash: commandStruct}
	redis := store.New()
	err := handler.IncrementStatisticCommandHandler(req, incrementStatisticCommand, redis)
	if err != nil {
		log.Println(err)
	}
}

func getFizzbuzz(writer driver.Writer, req *http.Request) {
	requiredParams := []string{"int1", "int2", "limit", "str1", "str2"}
	fizzbuzzCommand := command.GETFizzbuzzCommand{}
	err := validator.FillQueryCommandData(req, &fizzbuzzCommand, requiredParams)
	if err != nil {
		_ = transformer.SendJson(writer, driver.HTTPBadRequest, err.Error())
		return
	}

	fizzbuzzResult := handler.GETFizzbuzzCommandHandler(fizzbuzzCommand)
	err = transformer.SendJson(writer, driver.HTTPOk, fizzbuzzResult)
	if err != nil {
		return
	}

	handleStatistic(req, fizzbuzzCommand)
}

func FizzbuzzSupport(writer driver.Writer, req *http.Request) {
	if req.URL.Path == fizzbuzzPath {
		if req.Method == driver.GETMethod {
			getFizzbuzz(writer, req)
			return
		}
	}

	_ = transformer.SendJson(writer, driver.HTTPNotFound, driver.NotFoundText)
}
