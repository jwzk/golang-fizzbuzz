package handler

import (
	"fizzbuzz/application/command"
	"fmt"
	"strconv"
)

func getSupportedDivider(dividend int, dividers []int) int {
	for _, divider := range dividers {
		if (dividend % divider) == 0 {
			return divider
		}
	}

	return 0
}

func createElement(value int, int1 int, int2 int, str1 string, str2 string) string {
	dividers := []int{int1 * int2, int1, int2}
	supportedDivider := getSupportedDivider(value, dividers)
	switch supportedDivider {
	case int1:
		return str1
	case int2:
		return str2
	case int1 * int2:
		return fmt.Sprintf("%s%s", str1, str2)
	default:
		return strconv.FormatInt(int64(value), 10)
	}
}

func GETFizzbuzzCommandHandler(command command.GETFizzbuzzCommand) []string {
	if command.Limit <= 0 {
		return []string{}
	}

	fizzbuzzElements := make([]string, command.Limit)
	for i := 0; i < command.Limit; i++ {
		fizzbuzzElements[i] = createElement(i+1, command.Int1, command.Int2, command.Str1, command.Str2)
	}

	return fizzbuzzElements
}
