package CollectionUtils

import "reflect"

func IsEmpty(collection interface{}) bool {
	if collection == nil {
		return true
	}
	value := reflect.ValueOf(collection)
	kind := value.Type().Kind()
	if kind == reflect.Array || kind == reflect.Slice || kind == reflect.Map {
		return value.Len() <= 0
	}
	return false
}

func IsNotEmpty(collection interface{}) bool {
	return !IsEmpty(collection)
}
