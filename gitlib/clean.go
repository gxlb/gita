package gitlib

import (
	"fmt"
	"gita/exec"
	"strings"
)

//------------------------------------------------------------------------------

func CleanMissingRemoteBranches(untrack bool, forceRemove bool, confirmRemove bool) ([]string, error) {
	bs, err := GetMissingRemoteBranches()
	if err != nil {
		return nil, err
	}
	if !untrack && !forceRemove {
		fmt.Printf("isolate local branches: %d %v", len(bs), bs)
	}

	if untrack {
		for _, v := range bs {
			cmds := []string{
				fmt.Sprintf("git checkout %s", v),
				"git branch --unset-upstream",
			}
			for _, cmd := range cmds {
				if err := exec.ShellExecute(cmd); err != nil {
					fmt.Println(err)
					//return nil, err
				}
			}
			if err := exec.ShellExecute("git checkout master"); err != nil {
				return bs, err
			}
		}
		if !forceRemove {
			fmt.Printf("untracked pure local branches: %d %v", len(bs), bs)
		}
	}

	if forceRemove {
		if err := exec.ShellExecute("git checkout master"); err != nil {
			return bs, err
		}
		for _, v := range bs {
			cmd := fmt.Sprintf("git branch -d %s", v)
			if confirmRemove {
				cmd = strings.Replace(cmd, "-d", "-D", 1)
			}
			if err := exec.ShellExecute(cmd); err != nil {
				return bs, err
			}
		}
		fmt.Println("force removed local isolate branches:", bs)
	}

	return bs, nil
}

/*
  alpha
  bak/master0.0.1
  dev_demo
  develop
* master
  release/v0.6.1
  remotes/origin/alpha
  remotes/origin/bak/master0.0.1
  remotes/origin/dev_demo
  remotes/origin/develop
  remotes/origin/master
  remotes/origin/release/v0.6.1
  remotes/origin/release/v0.7.0
*/
func GetMissingRemoteBranches() ([]string, error) {
	if err := Fetch(); err != nil {
		return nil, err
	}
	r, err := exec.ShellExecuteResult("git branch -a")
	if err != nil {
		return nil, err
	}

	local := []string{}
	remotes := map[string]struct{}{}
	const remotePrefix = "remotes/origin/"
	bs := strings.Split(r, "\n")
	for _, v := range bs {
		t := v
		if t != "" && t[0] == '*' {
			t = t[1:]
		}
		t = strings.TrimSpace(t)
		if strings.HasPrefix(t, remotePrefix) {
			rb := strings.TrimPrefix(t, remotePrefix)
			remotes[rb] = struct{}{}
		} else {
			if t != "" { // last is ""
				local = append(local, t)
			}
		}
	}
	ret := local[:0]
	for _, v := range local {
		if _, ok := remotes[v]; !ok {
			ret = append(ret, v)
		}
	}
	return ret, nil
}
