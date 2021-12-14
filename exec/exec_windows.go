package exec

import (
	"bytes"
	"fmt"
	"os/exec"
)

func makeCmd(cmdStr string) *exec.Cmd {
	fmt.Println(">>", cmdStr)
	cmd := exec.Command("cmd", "/C", cmdStr)
	cmd.Stdin = bytes.NewBuffer([]byte(cmdStr))
	return cmd
}
