package CollectionUtils

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"
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

// IndexOf find index of list elements
// if not find, return -1
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

func ToString(list interface{}) string {
	if list == nil {
		return ""
	}
	value := reflect.ValueOf(list)
	if kind := value.Type().Kind(); kind != reflect.Array && kind != reflect.Slice {
		return value.String()
	}
	result := make([]string, 0, value.Len())
	for i := 0; i < value.Len(); i++ {
		result = append(result, convToStr(value.Index(i).Interface()))
	}
	return fmt.Sprintf("[%s]", strings.Join(result, ","))
}

func convToStr(value interface{}) string {
	// interface 转 string
	if value == nil {
		return ""
	}

	var key string
	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := jsoniter.Marshal(value)
		key = string(newValue)
	}
	return key
}
