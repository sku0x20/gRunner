package runner

import (
	"strings"
	"testing"
)

func Test_SpawnItself(t *testing.T) {
	if runReal() {
		realTest(t)
		return
	} else {
		output := spawnAndRun(funcName())
		lines := splitAndTrimLines(output)
		if !strings.Contains(lines[1], "fatal") {
			t.Fatalf("should print fatal")
		}
		if !strings.Contains(lines[2], "defer") {
			t.Fatalf("should print d")
		}
	}
}

func realTest(t *testing.T) {
	defer func() {
		t.Logf("defer")
	}()
	t.Fatalf("fatal")
}
