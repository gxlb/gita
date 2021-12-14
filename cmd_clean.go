package main

import (
	"gita/exec"
	"gita/gitlib"

	cli "github.com/urfave/cli/v2"
)

var cmdClean = &cli.Command{
	Name:    "clean",
	Aliases: []string{},
	Usage:   "clean local isolate branch.",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "view",
			Aliases: []string{"v"},
			Usage:   "view branches only.",
		},
		&cli.BoolFlag{
			Name:    "untrack",
			Aliases: []string{},
			Usage:   "force untrack isolate local branch.",
		},
		&cli.BoolFlag{
			Name:    "forceRemove",
			Aliases: []string{},
			Usage:   "force remove isolate local branch.",
		},
		&cli.BoolFlag{
			Name:    "confirmRemove",
			Aliases: []string{},
			Usage:   "confirm remove isolate local branch. -D",
		},
	},
	Description: `
git getch
git branch -a
git branch --unset-upstream
git branch -d <branch>
`,
	Action: func(c *cli.Context) error {
		if c.Bool("view") {
			gitlib.Fetch()
			return exec.ShellExecute("git branch -a")
		}

		untrack := c.Bool("untrack")
		forceRemove := c.Bool("forceRemove")
		confirmRemove := c.Bool("confirmRemove")
		_, err := gitlib.CleanMissingRemoteBranches(untrack, forceRemove, confirmRemove)
		if err != nil {
			return err
		}

		return nil
	},
}
