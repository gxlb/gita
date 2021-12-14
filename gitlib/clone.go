package gitlib

import (
	"fmt"
	"gita/exec"
)

// Clone execute `git clone` command
func Clone(url string, mirror bool) error {
	m := ""
	if mirror {
		m = " --mirror "
	}
	cmd := fmt.Sprintf("%s %s %s", CmdClone, m, url)

	return exec.ShellExecute(cmd)
}
