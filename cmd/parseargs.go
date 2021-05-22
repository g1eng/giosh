package cmd

import (
	"fmt"
	"github.com/g1eng/giosh/shell"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
)

func parseArgs(cli *cli.Context) (err error) {
	giosh := shell.New()

	if cli.Bool("debug") {
		giosh.SetDebug(true)
	} else {
		giosh.SetDebug(false)
	}

	if cli.NArg() > 0 {
		// read script from command line argument
		var b []byte
		fmt.Println(cli.Args().Get(0))
		if b, err = ioutil.ReadFile(cli.Args().Get(0)); err != nil {
			fmt.Print(err)
			return err
		}
		giosh.SetInput(string(b))
		if cli.Bool("debug") {
			log.Printf("b: string: %s : hex : %x", b, b)
			log.Printf("c.input: %s", giosh.GetInput())
		}
		return giosh.Run()
	} else {
		// read commands from stdin
		giosh.SetInput("")
		return giosh.Run()
	}
}
