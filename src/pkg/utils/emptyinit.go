package utils

import (
	"reflect"
)

func EmptyInit[E any]() E {
	t := reflect.TypeFor[E]()
	switch t.Kind() {
	case reflect.Ptr:
		rType := t.Elem()
		z := reflect.New(rType).Interface()
		return z.(E)
	case reflect.Slice:
		z := reflect.MakeSlice(t, 0, 10).Interface()
		return z.(E)
	default:
		return *new(E)
	}
}
