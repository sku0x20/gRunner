package utils

import (
	"github.com/sku0x20/gRunner/src/pkg/utils"
	"testing"
)

func Test_Int(t *testing.T) {
	e := utils.EmptyInit[int]()
	if e != 0 {
		t.Fatalf("e != 0")
	}
}

func Test_PointerInt(t *testing.T) {
	e := utils.EmptyInit[*int]()
	if e == nil {
		t.Fatalf("e == nil")
	}
}

func Test_SliceInt(t *testing.T) {
	e := utils.EmptyInit[[]int]()
	e = append(e, 12)
	if e[0] != 12 {
		t.Fatalf("e == nil")
	}
}
