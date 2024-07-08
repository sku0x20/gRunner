package utils

import (
	"reflect"
	"runtime"
	"strings"
)

func FuncName(f any) string {
	absoluteName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	split := strings.Split(absoluteName, ".")
	return split[len(split)-1]
}
