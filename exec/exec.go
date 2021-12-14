// + build linux|windows

package exec

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
)

func ShellExecuteResult(cmdStr string) (string, error) {
	buf := bytes.NewBuffer(nil)
	err := shellExecute(cmdStr, buf)
	return buf.String(), err
}

func ShellExecute(cmdStr string) error {
	return shellExecute(cmdStr, os.Stdout)
}

func ShellExecuteList(cmds []string) error {
	for i, v := range cmds {
		if err := shellExecute(v, os.Stdout); err != nil {
			return fmt.Errorf("%d/%d error:%s\n", i+1, len(cmds), err.Error())
		}
	}
	return nil
}

func shellExecute(cmdStr string, stdout io.Writer) error {
	cmd := makeCmd(cmdStr)
	if cmd.Stdout == nil {
		cmd.Stdout = stdout
	}
	if cmd.Stderr == nil {
		cmd.Stderr = os.Stderr
	}
	if cmd.Stdin == nil {
		cmd.Stdin = os.Stdin
	}

	return executeWait(cmd)
}

func executeWait(cmd *exec.Cmd) error {
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func executeSync(cmd *exec.Cmd, wg sync.WaitGroup) error {
	if err := cmd.Start(); err != nil {
		return err
	}

	wg.Add(1)
	defer wg.Done()
	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}
