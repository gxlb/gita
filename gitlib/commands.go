package gitlib

import (
	"fmt"
)

// git commands
const (
	CmdPull                 = "git pull"
	CmdFetch                = "git fetch"
	CmdFetchPrune           = "git fetch --prune" //"git remote prune origin"
	CmdStatus               = "git status -b -s"
	CmdBranchesGraph        = "git log --graph --decorate --oneline --simplify-by-decoration --all"
	CmdCheckoutF            = "git checkout %s"
	CmdForceCheckoutRemoteF = "git checkout -f --track -B %s remotes/origin/%s --"
	CmdClone                = "git clone"
)

const (
	// master release/v0.2.1
	ExprBranch = `(?:[-\w\./]+)`
)

var (
	// "## master...origin/master\n"
	ExprBranchStatusOK = fmt.Sprintf(`(?:^#{2}\s%s\.{3}%s\n$)`, ExprBranch, ExprBranch)
)
