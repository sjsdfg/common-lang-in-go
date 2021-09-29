package CollectionUtils

import (
	"reflect"
)

// MapToStringSlice extract string field from list
func MapToStringSlice(list interface{}, action func(i int) string) []string {
	if list == nil || action == nil {
		return []string{}
	}
	value := reflect.ValueOf(list)
	if kind := value.Type().Kind(); kind != reflect.Array && kind != reflect.Slice {
		return []string{}
	}
	result := make([]string, 0, value.Len()+value.Len()/2)
	for i := 0; i < value.Len(); i++ {
		result = append(result, action(i))
	}
	return result
}

func MapToStringSliceIgnoreEmpty(list interface{}, action func(i int) string) []string {
	if list == nil || action == nil {
		return []string{}
	}
	value := reflect.ValueOf(list)
	if kind := value.Type().Kind(); kind != reflect.Array && kind != reflect.Slice {
		return []string{}
	}
	result := make([]string, 0, value.Len()+value.Len()/2)
	for i := 0; i < value.Len(); i++ {
		s := action(i)
		if len(s) <= 0 {
			continue
		}
		result = append(result, s)
	}
	return result
}

func MapToStringSliceIgnoreByCondition(list interface{}, action func(i int) string, condition func(s string) bool) []string {
	if list == nil || action == nil {
		return []string{}
	}
	value := reflect.ValueOf(list)
	if kind := value.Type().Kind(); kind != reflect.Array && kind != reflect.Slice {
		return []string{}
	}
	result := make([]string, 0, value.Len()+value.Len()/2)
	for i := 0; i < value.Len(); i++ {
		s := action(i)
		if condition(s) {
			continue
		}
		result = append(result, s)
	}
	return result
}

// extract int field from list
func MapToIntSlice(list interface{}, action func(i int) int) []int {
	if list == nil || action == nil {
		return []int{}
	}
	value := reflect.ValueOf(list)
	if kind := value.Type().Kind(); kind != reflect.Array && kind != reflect.Slice {
		return []int{}
	}
	result := make([]int, 0, value.Len()+value.Len()/2)
	for i := 0; i < value.Len(); i++ {
		result = append(result, action(i))
	}
	return result
}

// extract int64 field from list
func MapToInt64Slice(list interface{}, action func(i int) int64) []int64 {
	if list == nil || action == nil {
		return []int64{}
	}
	value := reflect.ValueOf(list)
	if kind := value.Type().Kind(); kind != reflect.Array && kind != reflect.Slice {
		return []int64{}
	}
	result := make([]int64, 0, value.Len()+value.Len()/2)
	for i := 0; i < value.Len(); i++ {
		result = append(result, action(i))
	}
	return result
}

// extract float64 field from list
func MapToFloat64Slice(list interface{}, action func(i int) float64) []float64 {
	if list == nil || action == nil {
		return []float64{}
	}
	value := reflect.ValueOf(list)
	if kind := value.Type().Kind(); kind != reflect.Array && kind != reflect.Slice {
		return []float64{}
	}
	result := make([]float64, 0, value.Len()+value.Len()/2)
	for i := 0; i < value.Len(); i++ {
		result = append(result, action(i))
	}
	return result
}

// extract float32 field from list
func MapToFloat32Slice(list interface{}, action func(i int) float32) []float32 {
	if list == nil || action == nil {
		return []float32{}
	}
	value := reflect.ValueOf(list)
	if kind := value.Type().Kind(); kind != reflect.Array && kind != reflect.Slice {
		return []float32{}
	}
	result := make([]float32, 0, value.Len()+value.Len()/2)
	for i := 0; i < value.Len(); i++ {
		result = append(result, action(i))
	}
	return result
}
