package runner

import (
	"github.com/sku0x20/gRunner/src/pkg/runner"
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

func Test_TearDownAfterPanic(tm *testing.T) {
	r := runner.NewTestsRunner[any](tm)
	r.Teardown(func(t *testing.T, extra any) {
		recover()
	})
	r.Add(func(t *testing.T, extra any) {
		panic("test-panic")
	})
	r.Run()
}
