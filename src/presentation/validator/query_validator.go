package validator

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func setField(command interface{}, name string, value string) error {
	if value == "" {
		return fmt.Errorf("%s field is missing", name)
	}

	commandElem := reflect.ValueOf(command).Elem().Elem()
	fieldValue := reflect.Indirect(commandElem).FieldByName(strings.Title(name))

	if !fieldValue.IsValid() {
		return fmt.Errorf("%s field error", name)
	}

	if !fieldValue.CanSet() {
		return fmt.Errorf("%s field error", name)
	}

	val := reflect.ValueOf(value)
	if fieldValue.Type() != val.Type() {
		if val.Type().String() == "string" && fieldValue.Type().String() == "int" {
			intValue, err := strconv.Atoi(val.String())
			if err != nil {
				return fmt.Errorf("%s field type error", name)
			}
			intVal := reflect.ValueOf(intValue)
			fieldValue.Set(intVal)
			return nil
		}

		// unsupported type should return error
		// add more if necessary
		return fmt.Errorf("%s field error", name)
	}

	fieldValue.Set(val)
	return nil
}

func FillQueryCommandData(req *http.Request, command interface{}, params []string) error {
	for _, param := range params {
		err := setField(&command, param, req.URL.Query().Get(param))
		if err != nil {
			return err
		}
	}

	if len(req.URL.Query()) != len(params) {
		return fmt.Errorf("unexpected field")
	}

	return nil
}
