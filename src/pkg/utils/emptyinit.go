package utils

import (
	"reflect"
)

func EmptyInit[E any]() E {
	t := reflect.TypeFor[E]()
	if t.Kind() == reflect.Ptr {
		rType := t.Elem()
		z := reflect.New(rType).Interface()
		return z.(E)
	}
	return *new(E)
}
