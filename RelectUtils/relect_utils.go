package RelectUtils

import "reflect"

// Dereference 解除 Type 的指针
func Dereference(t reflect.Type) reflect.Type {
	if t.Kind() != reflect.Ptr {
		return t
	}
	return t.Elem()
}
