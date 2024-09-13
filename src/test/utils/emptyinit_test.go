package utils

import (
	"github.com/sku0x20/gRunner/src/pkg/utils"
	"testing"
)

func Test_Value(t *testing.T) {
	e := utils.EmptyInit[int]()
	if e != 0 {
		t.Fatalf("e != 0")
	}
}

func Test_Pointer(t *testing.T) {
	e := utils.EmptyInit[*int]()
	*e = 1000
	if *e != 1000 {
		t.Fatalf("*e != 1000")
	}
}

func Test_Slice(t *testing.T) {
	e := utils.EmptyInit[[]int]()
	e = append(e, 12)
	if e[0] != 12 {
		t.Fatalf("e[0] != 12")
	}
}
