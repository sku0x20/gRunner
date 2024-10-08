package utils

import (
	"github.com/sku0x20/gRunner/src/pkg/utils"
	"testing"
)

func Test_Value(t *testing.T) {
	e := utils.EmptyInit[int](t)
	if e != 0 {
		t.Fatalf("e != 0")
	}
}

func Test_Pointer(t *testing.T) {
	e := utils.EmptyInit[*int](t)
	*e = 1000
	if *e != 1000 {
		t.Fatalf("*e != 1000")
	}
}

func Test_Slice(t *testing.T) {
	e := utils.EmptyInit[[]int](t)
	e = append(e, 12)
	if e[0] != 12 {
		t.Fatalf("e[0] != 12")
	}
}

func Test_Map(t *testing.T) {
	e := utils.EmptyInit[map[int]int](t)
	e[23] = 32
	if e[23] != 32 {
		t.Fatalf("e[23] != 32")
	}
}
