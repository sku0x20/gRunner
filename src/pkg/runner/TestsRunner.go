package runner

import (
	. "github.com/sku0x20/gRunner/src/pkg/api"
	"github.com/sku0x20/gRunner/src/pkg/utils"
	"testing"
)

func NewTestsRunner[E any](
	t *testing.T,
	extraFunc ExtraInit[E],
) *TestsRunner[E] {
	return &TestsRunner[E]{
		t:                t,
		tests:            make([]TestFunc[E], 0, 10),
		setups:           make([]SetupFunc[E], 0, 10),
		teardowns:        make([]TeardownFunc[E], 0, 10),
		teardownsInverse: make([]TeardownFunc[E], 0, 10),
		extraFunc:        extraFunc,
	}
}

type TestsRunner[E any] struct {
	t                *testing.T
	tests            []TestFunc[E]
	setups           []SetupFunc[E]
	teardowns        []TeardownFunc[E]
	teardownsInverse []TeardownFunc[E]
	extraFunc        ExtraInit[E]
}

func (r *TestsRunner[E]) Add(f TestFunc[E]) {
	r.tests = append(r.tests, f)
}

func (r *TestsRunner[E]) Run() {
	for _, tf := range r.tests {
		r.t.Run(funcName(tf), func(t *testing.T) {
			extra := r.extraFunc()
			r.runSetups(t, extra)
			defer r.runTeardowns(t, extra)
			tf(t, extra)
		})
	}
}

func (r *TestsRunner[E]) runSetups(t *testing.T, extra E) {
	for _, setup := range r.setups {
		setup(t, extra)
	}
}

func (r *TestsRunner[E]) runTeardowns(t *testing.T, extra E) {
	for _, teardown := range r.teardowns {
		teardown(t, extra)
	}
	//for _, teardownInverse := range r.teardownsInverse {
	//	teardownInverse(t, extra)
	//}
}

func (r *TestsRunner[E]) Setup(f SetupFunc[E]) {
	r.setups = append(r.setups, f)
}

func (r *TestsRunner[E]) Teardown(f TeardownFunc[E]) {
	r.teardowns = append(r.teardowns, f)
}

// todo: rename to pushTeardown; stack/lifo
func (r *TestsRunner[E]) TeardownInverse(f TeardownFunc[E]) {
	r.teardownsInverse = append(r.teardownsInverse, f)
}

func funcName(f any) string {
	return utils.FuncName(f)
}
