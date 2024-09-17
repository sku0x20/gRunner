package runner

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

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
