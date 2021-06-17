package StringUtils

import (
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/sjsdfg/common-lang-in-go/IntUtils"
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

func EqualsNone(str string, list ...string) bool {
	return !EqualsAny(str, list...)
}

func EqualsNoneIgnoreCase(str string, list ...string) bool {
	return !EqualsAnyIgnoreCase(str, list...)
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

func IsNotDigital(str string) bool {
	return !IsDigital(str)
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

func Distinct(list ...string) []string {
	if len(list) <= 0 {
		return nil
	}
	result := make([]string, 0, len(list))
	distinctMap := make(map[string]struct{}, len(list))

	for _, s := range list {
		if _, found := distinctMap[s]; !found {
			distinctMap[s] = struct{}{}
			result = append(result, s)
		}
	}
	return result
}

func ToInt64(s string) int64 {
	if IsZero(s) {
		return 0
	}
	result, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return result
}

// 摘自 https://github.com/thinkeridea/go-extend/blob/main/exstrings/strings.go
// go-extend 追求高性能的 go 拓展库
func Reverse(s string) string {
	var start, size, end int
	buf := make([]byte, len(s))
	for end < len(s) {
		_, size = utf8.DecodeRuneInString(s[start:])
		end = start + size
		copy(buf[len(buf)-end:], s[start:end])
		start = end
	}
	return string(buf)
}

func ReplaceHolder(str string, holderMap map[string]string) string {
	if len(holderMap) <= 0 {
		return str
	}
	for holder, value := range holderMap {
		str = strings.ReplaceAll(str, holder, value)
	}
	return str
}
