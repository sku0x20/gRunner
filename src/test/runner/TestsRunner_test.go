package runner

import (
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"slices"
	"testing"
)

func Test_WithoutTest(t *testing.T) {
	r := runner.NewTestsRunner[any](t)
	r.Run()
}

func Test_WithoutFixtures(tm *testing.T) {
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

func Test_FixturesSameT(tm *testing.T) {
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

func Test_SameExtra(tm *testing.T) {
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

func Test_TeardownCalledAfterPanic(tm *testing.T) {
	r := runner.NewTestsRunner[any](tm)
	r.Teardown(func(t *testing.T, extra any) {
		recover()
	})
	r.Add(func(t *testing.T, extra any) {
		panic("test-panic")
	})
	r.Run()
}

func Test_TeardownCalledAfterFatal(tm *testing.T) {
	tm.Skip() // has to be tested manually; check if teardown is called
	r := runner.NewTestsRunner[any](tm)
	r.Teardown(func(t *testing.T, extra any) {
		t.Logf("teardown-called")
	})
	r.Add(func(t *testing.T, extra any) {
		t.Fatalf("test-fatal")
	})
	r.Run()
	if !tm.Failed() {
		tm.Fatalf("should have failed!")
	}
}

func Test_MultipleSetups(t *testing.T) {
	r := runner.NewTestsRunner[any](t)
	called := make([]string, 0, 2)
	r.Setup(func(t *testing.T) any {
		called = append(called, "s1")
		return nil
	})
	r.Setup(func(t *testing.T) any {
		called = append(called, "s2")
		return nil
	})
	r.Add(func(t *testing.T, extra any) {
		t.Log("test called")
	})
	r.Run()
	if len(called) != 2 {
		t.Fatalf("wrong number of setups, expected 2, got %d", len(called))
	}
	if !slices.Equal(called, []string{"s1", "s2"}) {
		t.Fatalf("wrong order")
	}
}
