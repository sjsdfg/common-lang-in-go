package CollectionUtils

import (
	"reflect"
	"sort"
)

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

func ForEach(list interface{}, action func(i int)) {
	if list == nil || action == nil {
		return
	}
	value := reflect.ValueOf(list)
	if kind := value.Type().Kind(); kind != reflect.Array && kind != reflect.Slice {
		return
	}
	for i := 0; i < value.Len(); i++ {
		action(i)
	}
}

func Sort(list interface{}, less func(i, j int) bool) {
	sort.SliceStable(list, less)
}

func IndexOf(list interface{}, match func(i int) bool) int {
	if list == nil || match == nil {
		return -1
	}
	value := reflect.ValueOf(list)
	if kind := value.Type().Kind(); kind != reflect.Array && kind != reflect.Slice {
		return -1
	}
	for i := 0; i < value.Len(); i++ {
		if find := match(i); find {
			return i
		}
	}
	return -1
}
