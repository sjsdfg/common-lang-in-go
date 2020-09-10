package CollectionUtils

import "reflect"

type matchFunc func(i int) bool

func AllMatch(list interface{}, action matchFunc) bool {
	if list == nil || action == nil {
		return false
	}
	value := reflect.ValueOf(list)
	if kind := value.Type().Kind(); kind != reflect.Array && kind != reflect.Slice {
		return false
	}
	for i := 0; i < value.Len(); i++ {
		match := action(i)
		if !match {
			return false
		}
	}
	return true
}

func AnyMatch(list interface{}, action matchFunc) bool {
	if list == nil || action == nil {
		return false
	}
	value := reflect.ValueOf(list)
	if kind := value.Type().Kind(); kind != reflect.Array && kind != reflect.Slice {
		return false
	}
	for i := 0; i < value.Len(); i++ {
		match := action(i)
		if match {
			return true
		}
	}
	return false
}

func NoneMatch(list interface{}, action matchFunc) bool {
	if list == nil || action == nil {
		return false
	}
	value := reflect.ValueOf(list)
	if kind := value.Type().Kind(); kind != reflect.Array && kind != reflect.Slice {
		return false
	}
	for i := 0; i < value.Len(); i++ {
		match := action(i)
		if match {
			return false
		}
	}
	return true
}
