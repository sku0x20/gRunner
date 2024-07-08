package utils

import (
	"github.com/sku0x20/gRunner/src/pkg/utils"
	"testing"
)

func Test_FuncName(t *testing.T) {
	if utils.FuncName(testing.Testing) != "Testing" {
		t.Fatalf("expected 'Testing got %s", utils.FuncName(testing.Testing))
	}
	if utils.FuncName(Test_FuncName) != "Test_FuncName" {
		t.Fatalf("expected 'Test_FuncName got %s", utils.FuncName(Test_FuncName))
	}
}
