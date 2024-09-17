package spikes

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"
)

// spawning itself
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

func runReal() bool {
	return os.Getenv("REAL") == "1"
}

func spawnAndRun(testFunc string) string {
	cmd := exec.Command(os.Args[0], "-test.run", "^"+testFunc+"$")
	cmd.Env = append(os.Environ(), "REAL=1")
	output, _ := cmd.CombinedOutput()
	return string(output)
}

func funcName() string {
	caller, _, _, _ := runtime.Caller(1)
	fullName := runtime.FuncForPC(caller).Name()
	split := strings.Split(fullName, ".")
	return split[len(split)-1]
}

func splitAndTrimLines(s string) []string {
	split := strings.Split(s, "\n")
	lines := make([]string, 0, len(split))
	for _, line := range split {
		trimmed := strings.TrimSpace(line)
		lines = append(lines, trimmed)
	}
	return lines
}

/*
"C:\Program Files\Go\bin\go.exe" test -c -o C:\Users\siddh\AppData\Local\JetBrains\GoLand2024.2\tmp\GoLand\___fork_test_go.test.exe github.com/sku0x20/gRunner/spikes #gosetup
"C:\Program Files\Go\bin\go.exe" tool test2json -t C:\Users\siddh\AppData\Local\JetBrains\GoLand2024.2\tmp\GoLand\___fork_test_go.test.exe -test.v=test2json -test.paniconexit0 -test.run ^\QTest_Process\E$ #gosetup
*/
