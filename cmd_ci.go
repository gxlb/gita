package main

import (
	"gita/exec"

	cli "github.com/urfave/cli/v2"
)

var cmdGoCi = &cli.Command{
	Name:    "go-ci",
	Aliases: []string{"ci"},
	Usage:   "ci check",
	Description: `
	golint ./...
	go test -cover ./...
`,
	Action: func(c *cli.Context) error {
		cmds := []string{
			"golint ./...",
			"go test -cover ./...",
		}
		if err := exec.ShellExecuteList(cmds); err != nil {
			return err
		}
		return nil
	},
}
