package main

import (
	cli "github.com/urfave/cli/v2"
)

var cmdRebase = &cli.Command{
	Name:    "rebase",
	Aliases: []string{"rb"},
	Usage:   "git rebase",
	Description: `
git rebase
`,
	Action: func(c *cli.Context) error {
		return nil
	},
}
