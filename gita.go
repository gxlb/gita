package main

import (
	"fmt"
	"os"

	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "gita",
		Usage:       "gita is git agent tool",
		Description: "gita is git agent tool",
		Flags:       []cli.Flag{},
		Version:     "v0.1.0-2011-12-28",
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "AllyDale",
				Email: "vipally@gmail.com",
			},
		},
		Commands: []*cli.Command{
			cmdStatus,
			cmdFetch,
			cmdPull,
			cmdGraph,
			cmdBranches,
			cmdDelete,
			cmdClean,
			cmdSync,
			cmdClone,
			cmdGoCi,
			cmdGoCover,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
