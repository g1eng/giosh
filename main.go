package main

import (
	"fmt"
	"github.com/g1eng/giosh/shell"
	"github.com/urfave/cli"
	"os"
)

func parserRun(cli *cli.Context) error {
	fmt.Printf(shell.GetPsString())
	commandLine := shell.CommandLine{}

	if cli.Bool("debug") {
		commandLine.SetDebug(true)
	} else {
		commandLine.SetDebug(false)
	}

	return commandLine.Exec()
}

func main() {
	app := &cli.App{
		Name:    "giosh",
		Version: "v0.1.1",
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
		Action:      parserRun,
	}
	_ = app.Run(os.Args)
}
