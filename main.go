package main

import (
	"fmt"
	"github.com/g1eng/giosh/shell"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
)

func parserRun(cli *cli.Context) (err error) {
	commandLine := shell.CommandLine{}

	if cli.Bool("debug") {
		commandLine.SetDebug(true)
	} else {
		commandLine.SetDebug(false)
	}

	if cli.NArg() > 0 {
		var b []byte
		fmt.Println(cli.Args().Get(0))
		if b, err = ioutil.ReadFile(cli.Args().Get(0)); err != nil {
			fmt.Print(err)
			return err
		}
		commandLine.SetInput(string(b))
		//log.Printf("b: string: %s : hex : %x", b, b)
		//log.Printf("c.input: %s", commandLine.GetInput())
		return commandLine.Exec()
	} else {
		commandLine.SetInput("")
		return commandLine.Exec()
	}
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
