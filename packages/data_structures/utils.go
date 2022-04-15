package data_structures

import (
	"reflect"
)

func Equals[T comparable](obj1 T, obj2 T) bool {
	dataType := reflect.TypeOf(obj1)

	switch dataType.Kind() {
	case reflect.Struct:
		return reflect.DeepEqual(obj1, obj2)
	default:
		return obj1 == obj2
	}
}
