package main

import (
	"fmt"
	"gita/exec"
	"gita/gitlib"

	cli "github.com/urfave/cli/v2"
)

var cmdBranches = &cli.Command{
	Name:    "branches",
	Aliases: []string{"b", "bs"},
	Usage:   "show branch list.",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "track",
			Aliases: []string{"t"},
			Usage:   "show track",
		},
		&cli.BoolFlag{
			Name:    "remote",
			Aliases: []string{"r", "a", "all"},
			Usage:   "show remote branch",
		},
		&cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"v"},
			Usage:   "show hash and subject, give twice for upstream branch",
		},
	},
	Description: `
git remote show origin
git branch (-v -a)
`,
	Action: func(c *cli.Context) error {
		cmd := ""
		switch {
		case c.Bool("track"):
			cmd = "git remote show origin"

		default:
			cmd = "git branch"
			if c.Bool("remote") {
				gitlib.Fetch()
				cmd += " -a"
			}
			if c.Bool("verbose") {
				cmd += " -v"
			}
		}
		fmt.Println("CurrentBranch:", gitlib.GetCurrentBranch())
		gitlib.GetBranchesGraph()
		err := exec.ShellExecute(cmd)
		if err != nil {
			return err
		}
		//fmt.Println(GetMissingRemoteBranches())
		return nil
	},
}
