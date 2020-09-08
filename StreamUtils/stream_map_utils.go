package StreamUtils

import "reflect"

// extract string field from list
func MapToStringSlice(list interface{}, action func(interface{}) string) []string {
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
		elem := value.Index(i)
		result = append(result, action(elem.Interface()))
	}
	return result
}
