package runner

import (
	"strings"
	"testing"
)

func Test_Process(t *testing.T) {
	if runReal() {
		realTest(t)
		return
	} else {
		output := spawnAndRun(funcName())
		lines := splitAndTrimLines(output)
		//for i, line := range output {
		//	t.Logf("%d %s", i, line)
		//}
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
