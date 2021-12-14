package main

import (
	"gita/exec"

	cli "github.com/urfave/cli/v2"
)

var cmdCheckout = &cli.Command{
	Name:    "pull",
	Aliases: []string{"p"},
	Usage:   "git pull",
	Description: `
git pull
`,
	Action: func(c *cli.Context) error {
		err := exec.ShellExecute("git pull")
		return err
	},
}
