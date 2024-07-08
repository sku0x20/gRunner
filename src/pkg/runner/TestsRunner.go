package runner

import (
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func NewTestsRunner[E any](t *testing.T) *TestsRunner[E] {
	return &TestsRunner[E]{
		t:        t,
		tests:    make([]TestFunc[E], 0, 10),
		setup:    func(t *testing.T) E { return *new(E) },
		teardown: func(t *testing.T, e E) {},
	}
}

type TestFunc[E any] func(t *testing.T, extra E)
type SetupFunc[E any] func(t *testing.T) E
type TeardownFunc[E any] func(t *testing.T, extra E)

type TestsRunner[E any] struct {
	t        *testing.T
	tests    []TestFunc[E]
	setup    SetupFunc[E]
	teardown TeardownFunc[E]
}

func (r *TestsRunner[E]) Add(f TestFunc[E]) {
	r.tests = append(r.tests, f)
}

func (r *TestsRunner[E]) Run() {
	for _, tf := range r.tests {
		r.t.Run(funcName(tf), func(t *testing.T) {
			extra := r.setup(t)
			tf(t, extra)
			r.teardown(t, extra)
		})
	}
}

func (r *TestsRunner[E]) Setup(f SetupFunc[E]) {
	r.setup = f
}

func (r *TestsRunner[E]) Teardown(f TeardownFunc[E]) {
	r.teardown = f
}

func funcName(f any) string {
	absoluteName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return strings.Split(absoluteName, ".")[1]
}
