package main

import (
	"gita/exec"
	"os"

	cli "github.com/urfave/cli/v2"
)

var cmdGoCover = &cli.Command{
	Name:    "go-cover",
	Aliases: []string{"cv", "gc"},
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "clean",
			Aliases: []string{"c"},
			Usage:   "clean result",
		},
	},
	Usage: "coverage check",
	Description: `
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
`,
	Action: func(c *cli.Context) error {
		if c.Bool("clean") {
			os.Remove("coverage.out")
			os.Remove("coverage.html")
			return nil
		}

		cmds := []string{
			`go test -coverprofile=coverage.out`,
			`go tool cover -html=coverage.out -o coverage.html`,
		}
		if err := exec.ShellExecuteList(cmds); err != nil {
			return err
		}
		return nil
	},
}
