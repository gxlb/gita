package main

import (
	"gita/exec"

	cli "github.com/urfave/cli/v2"
)

var cmdStatus = &cli.Command{
	Name:    "status",
	Aliases: []string{"s"},
	Usage:   "git status with short. ",
	Description: `
git status -b -s
git rev-parse --short HEAD
`,
	Action: func(c *cli.Context) error {
		if err := exec.ShellExecute(`git status -b -s`); err != nil {
			return err
		}
		if err := exec.ShellExecute(`git rev-parse --short HEAD`); err != nil {
			return err
		}
		return nil
	},
}
