package main

import (
	cli "github.com/urfave/cli/v2"
)

var cmdTag = &cli.Command{
	Name:    "tag",
	Aliases: []string{"t"},
	Usage:   "git tag",
	Description: `
git tag
`,
	Action: func(c *cli.Context) error {
		return nil
	},
}
