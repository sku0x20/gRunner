package test

import (
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func Test_TestRunner(t *testing.T) {
	r := runner.NewTestsRunner[string](t)
	r.Setup(setup)
	r.Teardown(teardown)
	r.Add(sampleTest)
	r.Run()
}

func setup(t *testing.T) string {
	// do setup
	return "setup-data"
}

func teardown(t *testing.T, extra string) {
	// do teardown
	t.Logf("teardown got extra '%s'", extra)
}

func sampleTest(t *testing.T, extra string) {
	if extra != "setup-data" {
		t.Fatalf("wrong value, expected \"setup-data\", got \"%s\"", extra)
	}
	t.Logf("should print function name '%s' in test runner", "sampleTest")
	t.Logf("sampleTest got extra '%s'", extra)
}
