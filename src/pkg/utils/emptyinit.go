package utils

import (
	"reflect"
)

func EmptyInit[E any]() E {
	e := *new(E)
	typeOf := reflect.TypeOf(e)
	if typeOf.Kind() == reflect.Ptr {
		rType := typeOf.Elem()
		z := reflect.New(rType).Interface()
		return z.(E)
	}
	return *new(E)
}
