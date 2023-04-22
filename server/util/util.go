package util

import (
	"github.com/df-mc/dragonfly/server"
	"unicode"
)

var (
	Server *server.Server
)

func IsStringAlphanumeric(str string) bool {
	for _, char := range str {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			return false
		}
	}
	return true
}

func InArray(val string, arr []string) bool {
	for _, item := range arr {
		if item == val {
			return true
		}
	}
	return false
}

func RemoveElementFromArray(arr []string, element string) []string {
	for i, v := range arr {
		if v == element {
			return append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}
