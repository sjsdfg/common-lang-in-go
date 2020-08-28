package StringUtils

import (
	"unicode"
)

func IsEmpty(str string) bool {
	return len(str) <= 0
}

func IsNotEmpty(str string) bool {
	return !IsNotEmpty(str)
}

func IsNotBlank(str string) bool {
	if IsEmpty(str) {
		return true
	}
	for _, c := range str {
		if unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

func IsBlank(str string) bool {
	return !IsNotBlank(str)
}
