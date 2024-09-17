package runner

import (
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"strings"
	"testing"
)

func Test_TeardownCalledAfterPanic(t *testing.T) {
	if runReal() {
		TeardownCalledAfterPanic_Real(t)
		return
	} else {
		output := spawnAndRun(funcName())
		lines := splitAndTrimLines(output)
		if !strings.Contains(lines[2], "call panic") {
			t.Fatalf("should print 'call panic'")
		}
		if !strings.Contains(lines[3], "teardown called") {
			t.Fatalf("should print 'teardown called'")
		}
	}
}

func TeardownCalledAfterPanic_Real(tm *testing.T) {
	r := runner.NewTestsRunnerEmptyInit[any](tm)
	r.Teardown(func(t *testing.T, extra any) {
		t.Logf("teardown called")
	})
	r.Add(func(t *testing.T, extra any) {
		t.Logf("call panic")
		panic("test-panic")
	})
	r.Run()
}

func Test_TeardownCalledAfterFatal(t *testing.T) {
	if runReal() {
		TeardownCalledAfterFatal_Real(t)
		return
	} else {
		output := spawnAndRun(funcName())
		lines := splitAndTrimLines(output)
		if !strings.Contains(lines[2], "call fatal") {
			t.Fatalf("should print 'call fatal'")
		}
		if !strings.Contains(lines[3], "teardown called") {
			t.Fatalf("should print 'teardown called'")
		}
	}
}

func TeardownCalledAfterFatal_Real(tm *testing.T) {
	r := runner.NewTestsRunnerEmptyInit[any](tm)
	r.Teardown(func(t *testing.T, extra any) {
		t.Logf("teardown called")
	})
	r.Add(func(t *testing.T, extra any) {
		t.Fatalf("call fatal")
	})
	r.Run()
}
