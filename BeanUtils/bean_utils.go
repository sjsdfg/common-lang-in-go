package BeanUtils

import "reflect"

func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	vi := reflect.ValueOf(i)
	return vi.Kind() == reflect.Ptr && vi.IsNil()
}

func IsAnyNil(list ...interface{}) bool {
	if len(list) <= 0 {
		return true
	}
	for _, element := range list {
		if IsNil(element) {
			return true
		}
	}
	return false
}

func IsNonNil(list ...interface{}) bool {
	return !IsAnyNil(list)
}
