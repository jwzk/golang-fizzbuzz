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
	statisticPath = "/statistic"
)

func getStatistic(writer driver.Writer, req *http.Request) {
	requiredParams := []string{}
	bestStatisticCommand := command.GETBestStatisticCommand{}
	err := validator.FillQueryCommandData(req, &bestStatisticCommand, requiredParams)
	if err != nil {
		_ = transformer.SendJson(writer, driver.HTTPBadRequest, err.Error())
		return
	}

	redis := store.New()
	bestStatisticResult, err := handler.GETBestStatisticCommandHandler(req, bestStatisticCommand, redis)
	if err != nil {
		log.Println(err)
		_ = transformer.SendJson(writer, driver.HTTPInternalServerError, driver.InternalServorErrorText)
		return
	}

	err = transformer.SendJson(writer, driver.HTTPOk, bestStatisticResult)
	if err != nil {
		return
	}
}

func StatisticSupport(writer driver.Writer, req *http.Request) {
	if req.URL.Path == statisticPath {
		if req.Method == driver.GETMethod {
			getStatistic(writer, req)
			return
		}
	}

	_ = transformer.SendJson(writer, driver.HTTPNotFound, driver.NotFoundText)
}
