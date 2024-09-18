package test

import (
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func Test_TestRunner(t *testing.T) {
	r := runner.NewTestsRunner[string](t, func(_ *testing.T) string {
		return "extra init"
	})
	r.Setup(setup)
	r.Teardown(teardown)
	r.Add(sampleTest)
	r.Run()
}

func setup(t *testing.T, extra string) {
	// do setup
	t.Logf("setup got extra '%s'", extra)
}

func teardown(t *testing.T, extra string) {
	// do teardown
	t.Logf("teardown got extra '%s'", extra)
}

func sampleTest(t *testing.T, extra string) {
	if extra != "extra init" {
		t.Fatalf("wrong value, expected \"extra init\", got \"%s\"", extra)
	}
	t.Logf("should print function name '%s' in test runner", "sampleTest")
	t.Logf("sampleTest got extra '%s'", extra)
}
