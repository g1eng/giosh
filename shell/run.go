package shell

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

// Run is the entry point for giosh shell interface.
//
// If you simply run it without any additional reader/writer sets,
// the vm starts with default I/O interface (bufio.Scanner and os.Stdout).
func (c *CommandLine) Run() error {
	var scanner *bufio.Scanner
	c.currentWriter = os.Stdout
	fmt.Printf(c.GetPsString())
	c.io.Writer = append(c.io.Writer, bufio.NewWriter(os.Stdout))
	if c.GetInput() == "" {
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		scanner = bufio.NewScanner(bytes.NewBufferString(c.GetInput()))
	}
	for scanner.Scan() {
		c.SetInput(scanner.Text())
		err := c.Parse()
		if err != nil {
			fmt.Print(err)
		}
	}

	return nil
}
