package api

import "testing"

type TestFunc[E any] func(t *testing.T, extra E)
type SetupFunc[E any] func(t *testing.T, extra E)
type TeardownFunc[E any] func(t *testing.T, extra E)
type ExtraInit[E any] func() E
