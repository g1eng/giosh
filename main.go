package main

import (
	"fmt"
	gioParser "github.com/g1eng/giop/core"
	"github.com/g1eng/giosh/shell"
	"github.com/urfave/cli"
	"os"
)

var giosh = gioParser.Init()

func parserRun (c *cli.Context) error {
	giosh.ReadLine()
	shell.PrintPsString()

	pipe := shell.PipeIO{}
	giosh.NewParser(pipe.PocPipe)

	if err := giosh.Parse(); err != nil {
		fmt.Println(err)
		fmt.Println("this is error")
		return err
	}
	return nil
}

func main() {
	(&cli.App{
		Name: "giosh",
		Usage: "giosh: a shell written in go and gioparser",
		Action: parserRun,
	}).Run(os.Args)
}
