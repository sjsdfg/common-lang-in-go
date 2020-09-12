package StringUtils

import (
	"github.com/sjsdfg/common-lang-in-go/IntUtils"
	"strings"
	"unicode"
)

func IsEmpty(str string) bool {
	return len(str) <= 0
}

func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
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

func IsAnyNoneEmpty(list ...string) bool {
	if len(list) <= 0 {
		return false
	}
	for _, s := range list {
		if IsNotEmpty(s) {
			return true
		}
	}
	return false
}

func IsBlank(str string) bool {
	return !IsNotBlank(str)
}

func IsNotBlank(str string) bool {
	if IsEmpty(str) {
		return false
	}
	return IsNotEmpty(strings.TrimSpace(str))
}

func IsZero(str string) bool {
	if IsEmpty(str) || str == "0" {
		return true
	}
	return false
}

func IsNotZero(str string) bool {
	return !IsZero(str)
}

func IsAnyZero(list ...string) bool {
	if len(list) <= 0 {
		return true
	}
	for _, s := range list {
		if IsZero(s) {
			return true
		}
	}
	return false
}

func IsAllZero(list ...string) bool {
	if len(list) <= 0 {
		return true
	}
	for _, s := range list {
		if IsNotZero(s) {
			return false
		}
	}
	return true
}

func IsAnyNoneZero(list ...string) bool {
	return !IsAllZero(list...)
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
	return If(IsNotEmpty(str), str, defaultStr)
}

func If(condition bool, ifTrue, ifFalse string) string {
	if condition {
		return ifTrue
	}
	return ifFalse
}

// string truncate [starIndex, endIndex)
// startIndex min is zero
// endIndex max is len(str)
func Truncate(str string, startIndex, endIndex int) string {
	if IsEmpty(str) {
		return Empty
	}
	return str[IntUtils.Max(0, startIndex):IntUtils.Min(len(str), endIndex)]
}

func Ptr(s string) *string {
	return &s
}

func GetPtrValue(s *string) string {
	if s == nil{
		return Empty
	}

	return *s
}

func GetPtrValueWithDefault(s *string, def string) string {
	if s == nil{
		return def
	}

	return *s
}