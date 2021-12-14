package main

import (
	"fmt"
	"gita/exec"
	"strings"

	cli "github.com/urfave/cli/v2"
)

var cmdDelete = &cli.Command{
	Name:    "delete",
	Aliases: []string{},
	Usage:   "delete branch.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "confirm",
			Usage:    "confirm operation. true",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "branch",
			Usage:    "branch to delete",
			Required: true,
		},
		&cli.BoolFlag{
			Name:     "remote",
			Usage:    "remove remote branch",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "force",
			Usage:    "force operate",
			Required: false,
		},
	},
	Description: `
git branch -d <branch>
`,
	Action: func(c *cli.Context) error {
		if confirm := c.String("confirm"); confirm != "true" {
			return fmt.Errorf("invalid confirm text %s, expect true", confirm)
		}
		cmd := ""
		branch := c.String("branch")
		switch {
		case c.Bool("remote"):
			cmd = fmt.Sprintf("git branch -r -d origin/%s", branch)
		default:
			cmd = fmt.Sprintf("git branch -d %s", branch)
		}
		if c.Bool("force") {
			cmd = strings.Replace(cmd, "-d", "-D", 1)
		}

		if err := exec.ShellExecute(cmd); err != nil {
			return err
		}
		return nil
	},
}
