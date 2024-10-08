package utils

import (
	"reflect"
	"testing"
)

func EmptyInit[E any](_ *testing.T) E {
	t := reflect.TypeFor[E]()
	switch t.Kind() {
	case reflect.Ptr:
		rType := t.Elem()
		z := reflect.New(rType).Interface()
		return z.(E)
	case reflect.Slice:
		z := reflect.MakeSlice(t, 0, 10).Interface()
		return z.(E)
	case reflect.Map:
		z := reflect.MakeMap(t).Interface()
		return z.(E)
	default:
		return *new(E)
	}
}
