package handler

import (
	"fizzbuzz/application/command"
	"reflect"
	"testing"
)

func TestGETFizzbuzzCommandHandler(t *testing.T) {
	inputs := []struct {
		int1     int
		int2     int
		limit    int
		str1     string
		str2     string
		expected []string
	}{
		{3, 5, 2, "fizz", "buzz", []string{"1", "2"}},
		{3, 5, 10, "fizz", "buzz", []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz"}},
		{2, 4, 10, "fizz", "buzz", []string{"1", "fizz", "3", "fizz", "5", "fizz", "7", "fizzbuzz", "9", "fizz"}},
		{2, 4, 0, "fizz", "buzz", []string{}},
		{2, 4, -42, "fizz", "buzz", []string{}},
	}

	for _, input := range inputs {
		fizzbuzzCommand := command.GETFizzbuzzCommand{Int1: input.int1, Int2: input.int2, Limit: input.limit, Str1: input.str1, Str2: input.str2}
		result := GETFizzbuzzCommandHandler(fizzbuzzCommand)

		if !reflect.DeepEqual(result, input.expected) {
			t.Errorf("'%v' should be '%v'", result, input.expected)
		}
	}
}
