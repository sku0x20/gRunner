package runner

import (
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"slices"
	"testing"
)

func Test_WithoutTest(t *testing.T) {
	r := runner.NewTestsRunnerEmptyInit[any](t)
	r.Run()
}

func Test_WithoutFixtures(tm *testing.T) {
	var t1t *testing.T = nil
	t1 := func(t *testing.T, extra any) {
		t1t = t
	}
	r := runner.NewTestsRunnerEmptyInit[any](tm)
	r.Add(t1)
	r.Run()
	if t1t == nil {
		tm.Fatalf("didn't ran the tests")
	}
}

func Test_SameT(tm *testing.T) {
	var initT *testing.T = nil
	var setupT *testing.T = nil
	var teardownT *testing.T = nil
	var testT *testing.T = nil
	setup := func(t *testing.T, extra any) {
		setupT = t
	}
	teardown := func(t *testing.T, extra any) {
		teardownT = t
	}
	test := func(t *testing.T, extra any) {
		testT = t
	}
	r := runner.NewTestsRunner[any](tm, func(t *testing.T) any {
		initT = t
		return nil
	})
	r.Setup(setup)
	r.Add(test)
	r.Teardown(teardown)
	r.Run()
	if initT != testT || setupT != testT || teardownT != testT {
		tm.Fatalf("t different for fixtures")
	}
}

func Test_SameExtra(tm *testing.T) {
	setup := func(t *testing.T, extra *string) {
		*extra = "some value"
	}
	t1 := func(t *testing.T, extra *string) {
		if *extra != "some value" {
			t.Fatalf("wrong value, expected \"some value\", got \"%s\"", *extra)
		}
	}
	teardown := func(t *testing.T, extra *string) {
		if *extra != "some value" {
			t.Fatalf("wrong value, expected \"some value\", got \"%s\"", *extra)
		}
	}
	r := runner.NewTestsRunner[*string](tm, func(t *testing.T) *string {
		s := "some value"
		return &s
	})
	r.Setup(setup)
	r.Teardown(teardown)
	r.Add(t1)
	r.Run()
}

func Test_ExtraInit(t *testing.T) {
	r := runner.NewTestsRunner[*string](t, func(t *testing.T) *string {
		s := "some value"
		return &s
	})
	r.Add(func(t *testing.T, extra *string) {
		if *extra != "some value" {
			t.Fatalf("wrong value, expected \"some value\", got \"%s\"", *extra)
		}
	})
	r.Run()
}

func Test_MultipleSetups(t *testing.T) {
	r := runner.NewTestsRunnerEmptyInit[any](t)
	called := make([]string, 0, 2)
	r.Setup(func(t *testing.T, extra any) {
		called = append(called, "s1")
	})
	r.Setup(func(t *testing.T, extra any) {
		called = append(called, "s2")
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

func Test_MultipleTeardowns(t *testing.T) {
	r := runner.NewTestsRunnerEmptyInit[any](t)
	called := make([]string, 0, 2)
	r.Teardown(func(t *testing.T, extra any) {
		called = append(called, "t1")
	})
	r.Add(func(t *testing.T, extra any) {
		t.Log("test called")
	})
	r.Teardown(func(t *testing.T, extra any) {
		called = append(called, "t2")
	})
	r.Run()
	if len(called) != 2 {
		t.Fatalf("wrong number of setups, expected 2, got %d", len(called))
	}
	if !slices.Equal(called, []string{"t1", "t2"}) {
		t.Fatalf("wrong order")
	}
}

func Test_PushTeardown(t *testing.T) {
	r := runner.NewTestsRunnerEmptyInit[any](t)
	called := make([]string, 0, 2)
	r.PushTeardown(func(t *testing.T, extra any) {
		called = append(called, "t3")
	})
	r.PushTeardown(func(t *testing.T, extra any) {
		called = append(called, "t2")
	})
	r.Teardown(func(t *testing.T, extra any) {
		called = append(called, "t1")
	})
	r.Add(func(t *testing.T, extra any) {
		t.Log("test called")
	})
	r.Run()
	if len(called) != 3 {
		t.Fatalf("wrong number of setups, expected 3, got %d", len(called))
	}
	if !slices.Equal(called, []string{"t1", "t2", "t3"}) {
		t.Fatalf("wrong order")
	}
}
