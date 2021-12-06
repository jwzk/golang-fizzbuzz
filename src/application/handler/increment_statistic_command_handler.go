package handler

import (
	"encoding/json"
	"fizzbuzz/application/command"
	"fizzbuzz/driver"
	"net/http"
)

func IncrementStatisticCommandHandler(req *http.Request, command command.IncrementStatisticCommand, store driver.Store) error {
	hash, err := json.Marshal(command.Hash)
	if err != nil {
		return err
	}

	storeKey := string(hash)
	err = store.Increment(req.Context(), storeKey)
	if err != nil {
		return err
	}

	return nil
}
