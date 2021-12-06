package app

import (
	"fizzbuzz/driver"
	"fizzbuzz/presentation/controller"
	"fizzbuzz/presentation/transformer"
	"log"
	"net/http"
	"strings"
)

const (
	fizzbuzzResource  = "fizzbuzz"
	statisticResource = "statistic"
	pathDelimiter     = "/"
)

func Router(writer driver.Writer, req *http.Request) {
	log.Printf("HTTP request %s", req.URL.Path)
	resourceName := strings.Split(req.URL.Path, pathDelimiter)[1]

	if resourceName == fizzbuzzResource {
		controller.FizzbuzzSupport(writer, req)
		return
	}

	if resourceName == statisticResource {
		controller.StatisticSupport(writer, req)
		return
	}

	_ = transformer.SendJson(writer, driver.HTTPNotFound, driver.NotFoundText)
}
