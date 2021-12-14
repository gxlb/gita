package main

import (
	"gita/exec"
	"gita/gitlib"

	cli "github.com/urfave/cli/v2"
)

var cmdPull = &cli.Command{
	Name:    "pull",
	Aliases: []string{"p"},
	Usage:   "git pull",
	Description: `
git pull
`,
	Action: func(c *cli.Context) error {
		err := exec.ShellExecute(gitlib.CmdPull)
		return err
	},
}
