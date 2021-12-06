package handler

import (
	"fizzbuzz/application/command"
	"fizzbuzz/driver"
	"net/http"
	"strconv"
	"strings"
)

type BestStructRes struct {
	Hits       int    `json:"hits"`
	Parameters string `json:"parameters"`
}

func GETBestStatisticCommandHandler(req *http.Request, command command.GETBestStatisticCommand, store driver.Store) (BestStructRes, error) {
	keys, err := store.GetKeys(req.Context(), "*")
	if err != nil {
		return BestStructRes{}, err
	}

	counter := 0
	bestResult := ""
	for _, key := range keys {
		data, err := strconv.Atoi(store.GetDataFromKey(req.Context(), key))
		if err != nil {
			break
		}

		if data > counter {
			counter = data
			bestResult = key
		}
	}

	replacer := strings.NewReplacer(`"`, "", "{", "", "}", "")

	return BestStructRes{counter, replacer.Replace(bestResult)}, nil
}
