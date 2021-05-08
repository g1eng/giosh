package main

import (
	"fmt"
	gioParser "github.com/g1eng/giop/core"
	"github.com/g1eng/giosh/shell"
	"github.com/urfave/cli"
	"log"
	"os"
)

func parserRun(_ *cli.Context) error {
	vm := gioParser.NewParser()
	fmt.Printf(shell.GetPsString())
	vm.ReadLine()

	commandLine := shell.CommandLine{}
	vm.AddFilter(commandLine.Exec)

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
