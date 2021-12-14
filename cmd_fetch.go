package main

import (
	"gita/exec"
	"gita/gitlib"

	cli "github.com/urfave/cli/v2"
)

var cmdFetch = &cli.Command{
	Name:    "fetch",
	Aliases: []string{"f"},
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "clean",
			Aliases: []string{"c"},
			Usage:   "clean local isolate branches",
		},
	},
	Usage: "git fetch",
	Description: `
git fetch (-p)
git remote update origin --prune
`,
	Action: func(c *cli.Context) error {
		cmd := gitlib.CmdFetch
		if c.Bool("clean") {
			//cmd = "git remote update origin --prune"
			cmd = "git remote prune origin"
		}
		if err := exec.ShellExecute(cmd); err != nil {
			return err
		}
		return nil
	},
}
