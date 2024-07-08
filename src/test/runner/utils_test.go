package runner

import (
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func Test_FuncName(t *testing.T) {
	if runner.FuncName(testing.Testing) != "Testing" {
		t.Fatalf("expected 'Testing got %s", runner.FuncName(testing.Testing))
	}
	if runner.FuncName(Test_FuncName) != "Test_FuncName" {
		t.Fatalf("expected 'Test_FuncName got %s", runner.FuncName(Test_FuncName))
	}
}
