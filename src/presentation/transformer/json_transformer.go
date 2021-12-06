package transformer

import (
	"encoding/json"
	"errors"
	"fizzbuzz/driver"
	"log"
)

type header struct {
	key   string
	value string
}

func formatToJSON(content interface{}) ([]byte, error) {
	json, err := json.Marshal(content)
	if err != nil {
		return nil, errors.New("invalid content")
	}

	return json, nil
}

func formatResponseToJSON(content interface{}) ([]byte, error) {
	if content == nil {
		return nil, nil
	}

	response, err := formatToJSON(content)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func sendJSONContent(writer driver.Writer, status int, content interface{}) error {
	response, err := formatResponseToJSON(content)
	if err != nil {
		log.Println(err)
		jsonErr, _ := formatToJSON("json format error")
		writer.WriteHeader(driver.HTTPInternalServerError)

		_, err = writer.Write(jsonErr)
		if err != nil {
			log.Println(err)
		}
		return errors.New("send request error")
	}

	writer.WriteHeader(status)

	_, err = writer.Write(response)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(driver.HTTPInternalServerError)
		return errors.New("send request error")
	}

	return nil
}

func setJSONHeaders(writer driver.Writer) {
	jsonHeader := header{
		key:   "Content-Type",
		value: "application/json",
	}

	writer.Header().Set(jsonHeader.key, jsonHeader.value)
}

func SendJson(writer driver.Writer, status int, content interface{}) error {
	setJSONHeaders(writer)
	err := sendJSONContent(writer, status, content)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("HTTP request with response with status %d", status)
	return nil
}
