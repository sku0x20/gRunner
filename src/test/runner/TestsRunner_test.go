package runner

import (
	"gRunner/src/pkg/runner"
	"testing"
)

func Test_Nop(t *testing.T) {
	r := runner.NewTestsRunner[any](t)
	r.Run()
}

func Test_NoFixtures(tm *testing.T) {
	var t1t *testing.T = nil
	t1 := func(t *testing.T, extra any) {
		t1t = t
	}
	r := runner.NewTestsRunner[any](tm)
	r.Add(t1)
	r.Run()
	if t1t == nil {
		tm.Fatalf("didn't ran the tests")
	}
}

func Test_Fixtures(tm *testing.T) {
	var setupT *testing.T = nil
	var teardownT *testing.T = nil
	var t1t *testing.T = nil
	setup := func(t *testing.T) any {
		setupT = t
		return nil
	}
	teardown := func(t *testing.T, extra any) {
		teardownT = t
	}
	t1 := func(t *testing.T, extra any) {
		t1t = t
	}
	r := runner.NewTestsRunner[any](tm)
	r.Setup(setup)
	r.Add(t1)
	r.Teardown(teardown)
	r.Run()
	if setupT != t1t || teardownT != t1t {
		tm.Fatalf("t different for fixtures")
	}
}

func Test_Extra(tm *testing.T) {
	setup := func(t *testing.T) string {
		return "some value"
	}
	t1 := func(t *testing.T, extra string) {
		if extra != "some value" {
			t.Fatalf("wrong value, expected \"some value\", got \"%s\"", extra)
		}
	}
	teardown := func(t *testing.T, extra string) {
		if extra != "some value" {
			t.Fatalf("wrong value, expected \"some value\", got \"%s\"", extra)
		}
	}
	r := runner.NewTestsRunner[string](tm)
	r.Setup(setup)
	r.Teardown(teardown)
	r.Add(t1)
	r.Run()
}

// teardownCalledEvenIfPanic
// teardownCalledEvenIfFatal

// support for one time setup and teardown

// Usage
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
