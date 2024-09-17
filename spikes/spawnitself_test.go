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
	if os.Getenv("REAL") == "1" {
		realTest(t)
		return
	}
	thisFunc := callerFuncName()
	lines := runTest(thisFunc)
	//for i, line := range lines {
	//	t.Logf("%d %s", i, line)
	//}
	if !strings.Contains(lines[1], "fatal") {
		t.Fatalf("should print fatal")
	}
	if !strings.Contains(lines[2], "defer") {
		t.Fatalf("should print d")
	}
}

func runTest(thisFunc string) []string {
	cmd := exec.Command(os.Args[0], "-test.run", "^"+thisFunc+"$")
	cmd.Env = append(os.Environ(), "REAL=1")
	output, _ := cmd.CombinedOutput()
	lines := getLines(string(output))
	return lines
}

func callerFuncName() string {
	caller, _, _, _ := runtime.Caller(1)
	fullName := runtime.FuncForPC(caller).Name()
	split := strings.Split(fullName, ".")
	return split[len(split)-1]
}

func getLines(s string) []string {
	split := strings.Split(s, "\n")
	lines := make([]string, 0, len(split))
	for _, line := range split {
		trimmed := strings.TrimSpace(line)
		lines = append(lines, trimmed)
	}
	return lines
}

func realTest(t *testing.T) {
	defer func() {
		t.Logf("defer")
	}()
	t.Fatalf("fatal")
}

/*
"C:\Program Files\Go\bin\go.exe" test -c -o C:\Users\siddh\AppData\Local\JetBrains\GoLand2024.2\tmp\GoLand\___fork_test_go.test.exe github.com/sku0x20/gRunner/spikes #gosetup
"C:\Program Files\Go\bin\go.exe" tool test2json -t C:\Users\siddh\AppData\Local\JetBrains\GoLand2024.2\tmp\GoLand\___fork_test_go.test.exe -test.v=test2json -test.paniconexit0 -test.run ^\QTest_Process\E$ #gosetup
*/
