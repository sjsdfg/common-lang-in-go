package CollectionUtils

import "reflect"

// extract string field from list
func MapToStringSlice(list interface{}, action func(index int) string) []string {
	if list == nil || action == nil {
		return []string{}
	}
	value := reflect.ValueOf(list)
	kind := value.Type().Kind()
	if kind != reflect.Array && kind != reflect.Slice {
		return []string{}
	}
	result := make([]string, 0, value.Len())
	for i := 0; i < value.Len(); i++ {
		result = append(result, action(i))
	}
	return result
}

// extract int field from list
func MapToIntSlice(list interface{}, action func(interface{}) int) []int {
	if list == nil || action == nil {
		return []int{}
	}
	value := reflect.ValueOf(list)
	kind := value.Type().Kind()
	if kind != reflect.Array && kind != reflect.Slice {
		return []int{}
	}
	result := make([]int, 0, value.Len())
	for i := 0; i < value.Len(); i++ {
		elem := value.Index(i)
		result = append(result, action(elem.Interface()))
	}
	return result
}

// extract int64 field from list
func MapToInt64Slice(list interface{}, action func(interface{}) int64) []int64 {
	if list == nil || action == nil {
		return []int64{}
	}
	value := reflect.ValueOf(list)
	kind := value.Type().Kind()
	if kind != reflect.Array && kind != reflect.Slice {
		return []int64{}
	}
	result := make([]int64, 0, value.Len())
	for i := 0; i < value.Len(); i++ {
		elem := value.Index(i)
		result = append(result, action(elem.Interface()))
	}
	return result
}

// extract float64 field from list
func MapToFloat64Slice(list interface{}, action func(interface{}) float64) []float64 {
	if list == nil || action == nil {
		return []float64{}
	}
	value := reflect.ValueOf(list)
	kind := value.Type().Kind()
	if kind != reflect.Array && kind != reflect.Slice {
		return []float64{}
	}
	result := make([]float64, 0, value.Len())
	for i := 0; i < value.Len(); i++ {
		elem := value.Index(i)
		result = append(result, action(elem.Interface()))
	}
	return result
}

// extract float32 field from list
func MapToFloat32Slice(list interface{}, action func(interface{}) float32) []float32 {
	if list == nil || action == nil {
		return []float32{}
	}
	value := reflect.ValueOf(list)
	kind := value.Type().Kind()
	if kind != reflect.Array && kind != reflect.Slice {
		return []float32{}
	}
	result := make([]float32, 0, value.Len())
	for i := 0; i < value.Len(); i++ {
		elem := value.Index(i)
		result = append(result, action(elem.Interface()))
	}
	return result
}
