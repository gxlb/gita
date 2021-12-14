package main

import (
	cli "github.com/urfave/cli/v2"
)

var cmdTrack = &cli.Command{
	Name:    "track",
	Aliases: []string{},
	Usage:   "git track",
	Description: `
`,
	Action: func(c *cli.Context) error {
		return nil
	},
}
