package main

import (
	"fmt"
	gioParser "github.com/g1eng/giop/core"
	"github.com/g1eng/giosh/shell"
	"github.com/urfave/cli"
	"log"
	"os"
)

var vm = gioParser.NewParser()

func parserRun(_ *cli.Context) error {
	fmt.Printf(shell.GetPsString())
	vm.ReadLine()

	pipe := shell.PipeIO{}
	vm.AddFilter(pipe.PocPipe)

	if err := vm.Parse(); err != nil {
		fmt.Println(err)
		fmt.Println("this is error")
		return err
	}
	return nil
}

func main() {
	err := (&cli.App{
		Name:   "giosh",
		Usage:  "giosh: a shell written in go and gioparser",
		Action: parserRun,
	}).Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
