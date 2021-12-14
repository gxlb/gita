package main

import (
	"fmt"
	"gita/exec"
	"gita/gitlib"

	cli "github.com/urfave/cli/v2"
)

var cmdSync = &cli.Command{
	Name:    "sync",
	Aliases: []string{},
	Usage:   "sync branch",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "src",
			Aliases:  []string{"s"},
			Usage:    "source branch.",
			Required: false,
		},
		&cli.StringSliceFlag{
			Name:     "dst",
			Aliases:  []string{"t"},
			Usage:    "destination branches.",
			Required: false,
		},
		&cli.StringSliceFlag{
			Name:     "tx",
			Aliases:  []string{"tX"},
			Usage:    "no use to hide flag t.",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "realDoSync",
			Aliases:  []string{},
			Usage:    "really do sync action",
			Required: false,
		},
	},
	Description: `

`,
	Action: func(c *cli.Context) error {
		src := c.String("src")
		dst := c.StringSlice("dst")
		realDoSync := c.Bool("realDoSync")

		if src == "" {
			gitlib.Fetch()
			return exec.ShellExecute(gitlib.CmdBranchesGraph)
		}
		if len(dst) == 0 {
			return fmt.Errorf("missing dst flag")
		}
		bg, err := gitlib.GetBranchesGraph()
		if err != nil {
			return err
		}
		if err := bg.SyncBranchList(src, dst, realDoSync); err != nil {
			return err
		}
		return nil
	},
}
