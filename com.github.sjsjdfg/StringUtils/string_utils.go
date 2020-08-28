package StringUtils

import (
	"strings"
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

func IsAllEmpty(list ...string) bool {
	if len(list) <= 0 {
		return true
	}
	for _, s := range list {
		if IsNotEmpty(s) {
			return false
		}
	}
	return true
}

func IsAnyEmpty(list ...string) bool {
	if len(list) <= 0 {
		return true
	}
	for _, s := range list {
		if IsEmpty(s) {
			return true
		}
	}
	return false
}

func Equal(str1, str2 string) bool {
	if IsAllEmpty(str1, str2) {
		return true
	}
	return str1 == str2
}

func EqualIgnoreCase(str1, str2 string) bool {
	if IsAllEmpty(str1, str2) {
		return true
	}
	return strings.ToLower(str1) == strings.ToLower(str2)
}

func EqualsAny(str string, list ...string) bool {
	if len(list) <= 0 {
		return false
	}
	for _, s := range list {
		if Equal(str, s) {
			return true
		}
	}
	return false
}

func EqualsAnyIgnoreCase(str string, list ...string) bool {
	if len(list) <= 0 {
		return false
	}
	for _, s := range list {
		if EqualIgnoreCase(str, s) {
			return true
		}
	}
	return false
}

func IsDigital(str string) bool {
	if IsBlank(str) {
		return false
	}
	for _, s := range str {
		if !unicode.IsDigit(s) {
			return false
		}
	}
	return true
}

func DefaultIfEmpty(str, defaultStr string) string {
	if IsNotEmpty(str) {
		return str
	}
	return defaultStr
}
