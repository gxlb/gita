package main

import (
	cli "github.com/urfave/cli/v2"
)

var cmdRename = &cli.Command{
	Name:    "rename",
	Aliases: []string{"rn"},
	Usage:   "rename",
	Description: `
`,
	Action: func(c *cli.Context) error {
		return nil
	},
}
