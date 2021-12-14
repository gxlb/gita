package exec

import (
	"testing"
)

func TestExec(t *testing.T) {
	ShellExecute("echo hello")
	ShellExecute("set path")
}
