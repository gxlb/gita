package exec

import (
	"os/exec"
)

func makeCmd(cmdStr string) *exec.Cmd {
	cmd := exec.Command("/bin/bash", "-c", cmdStr, "|sh")
	return cmd
}
