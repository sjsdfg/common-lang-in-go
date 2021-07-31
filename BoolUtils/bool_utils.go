package BoolUtils

import "strconv"

func And(list ...bool) bool {
	if len(list) <= 0 {
		return false
	}
	for _, b := range list {
		if IsFalse(b) {
			return false
		}
	}
	return true
}

func Or(list ...bool) bool {
	if len(list) <= 0 {
		return false
	}
	for _, b := range list {
		if b {
			return true
		}
	}
	return false
}

func Xor(list ...bool) bool {
	if len(list) <= 0 {
		return false
	}
	// false if the neutral element of the xor operator
	result := false
	for _, b := range list {
		result = xor(result, b)
	}
	return result
}

// xor(true, true)   = false
// xor(false, false) = false
// xor(true, false)  = true
func xor(a, b bool) bool {
	return a != b
}

func IsFalse(b bool) bool {
	return !b
}

func ValueOf(str string) bool {
	parseBool, err := strconv.ParseBool(str)
	if err != nil {
		return false
	}
	return parseBool
}
