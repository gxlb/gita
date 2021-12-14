package main

import (
	"gita/gitlib"

	cli "github.com/urfave/cli/v2"
)

var cmdClone = &cli.Command{
	Name:    "clone",
	Aliases: []string{"c"},
	Usage:   "git clone",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "mirror",
			Aliases: []string{"m", "backup", "b"},
			Usage:   "create mirror",
		},
		&cli.StringFlag{
			Name:     "url",
			Aliases:  []string{"u", "r", "s"},
			Usage:    "remote url",
			Required: true,
		},
	},
	Description: `
git clone
`,
	Action: func(c *cli.Context) error {
		return gitlib.Clone(c.String("url"), c.Bool("mirror"))
	},
}
