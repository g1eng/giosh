package main

import (
	"fmt"
	gioParser "github.com/g1eng/giop/core"
	"github.com/g1eng/giosh/shell"
	"github.com/urfave/cli"
	"os"
)

func parserRun(c *cli.Context) error {
	vm := gioParser.NewParser()
	fmt.Printf(shell.GetPsString())
	vm.ReadLine()

	commandLine := shell.CommandLine{}
	if c.Bool("debug") {
		commandLine.SetDebug(true)
	} else {
		commandLine.SetDebug(false)
	}
	vm.AddFilter(commandLine.Exec)

	if err := vm.Parse(); err != nil {
		if commandLine.DumpErrors() {
			fmt.Println(err)
			fmt.Println("pipe failure")
		}
		return err
	}
	return nil
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
