package CollectionUtils

import "reflect"

func FlatMapToStringSlice(list interface{}, action func(i int) []string) []string {
	if list == nil || action == nil {
		return nil
	}
	value := reflect.ValueOf(list)
	if kind := value.Type().Kind(); kind != reflect.Array && kind != reflect.Slice {
		return nil
	}
	result := make([]string, 0, 3*value.Len()+value.Len()/2)
	for i := 0; i < value.Len(); i++ {
		result = append(result, action(i)...)
	}
	return result
}
