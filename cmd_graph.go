package main

import (
	"gita/exec"
	"gita/gitlib"

	cli "github.com/urfave/cli/v2"
)

var cmdGraph = &cli.Command{
	Name:    "graph",
	Aliases: []string{"g"},
	Usage:   "git branch graph.",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "nofetch",
			Aliases: []string{"nf", "n"},
			Usage:   "no featch",
		},
	},
	Description: `
git fetch
git status -b -s
git log --graph --decorate --oneline --simplify-by-decoration --all
`,
	Action: func(c *cli.Context) error {
		if !c.Bool("nofetch") {
			if err := gitlib.Fetch(); err != nil {
				return err
			}
		}
		cmds := []string{
			gitlib.CmdStatus,
			gitlib.CmdBranchesGraph,
		}
		if err := exec.ShellExecuteList(cmds); err != nil {
			return err
		}
		return nil
	},
}
