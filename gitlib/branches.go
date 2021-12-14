package gitlib

import (
	"fmt"
	"gita/exec"
	"regexp"
	"strings"
)

var regexpStatusOK = regexp.MustCompile(ExprBranchStatusOK)

func ValidateCheckout() error {
	r, err := exec.ShellExecuteResult(CmdStatus)
	if err != nil {
		return err
	}
	if regexpStatusOK.MatchString(r) {
		return nil
	}
	return fmt.Errorf("branch status deny switch:\n%s", r)
}

func Checkout(branch string, remote bool, check bool) error {
	if check {
		if err := ValidateCheckout(); err != nil {
			return err
		}
	}

	cmd := fmt.Sprintf(CmdCheckoutF, branch)
	if remote {
		cmd = fmt.Sprintf(CmdForceCheckoutRemoteF, branch, branch)
	}
	if err := exec.ShellExecute(cmd); err != nil {
		return err
	}
	return nil
}

func GetCurrentBranch() string {
	r, err := exec.ShellExecuteResult("git branch")
	if err != nil {
		return ""
	}
	bs := strings.Split(r, "\n")
	// local branches: []string{"alpha", "  dev_demo", "  develop", "* master", ""}
	// fmt.Printf("local branches: %#v\n", bs)
	for _, v := range bs {
		if len(v) > 0 && v[0] == '*' {
			return strings.TrimSpace(v[1:])
		}
	}
	return ""
}

func Fetch() error {
	return exec.ShellExecute(CmdFetchPrune)
}
