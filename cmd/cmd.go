package cmd

import (
	"github.com/urfave/cli"
	"os"
)

func Cmd() {
	app := &cli.App{
		Name:    "giosh",
		Version: "v0.1.2",
		Authors: []*cli.Author{
			{
				Name:  "Suzume Nomura",
				Email: "suzume315@g00.g1e.org",
			},
		},
		Copyright: "(c) 2021 G1 Engineering Committers",
		Usage:     "invoke and execute",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "show debug messages",
			},
		},
		HideVersion: true,
		HideHelp:    true,
		Action:      parseArgs,
	}
	_ = app.Run(os.Args)
}
