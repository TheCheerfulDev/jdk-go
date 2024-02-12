package jdkutil

import (
	"strings"
)

func RemoveNewLineFromString(input string) string {
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.ReplaceAll(input, "\r", "")
	return input
}
