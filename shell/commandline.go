package shell

import (
	"fmt"
	"os/exec"
)

func (c *CommandLine) TerminateLine(withPsString ...bool) (err error) {
	// for bufio.Writer, write PS string
	err = c.DumpErrors()
	if err != nil {
		fmt.Println(err)
	}
	if len(withPsString) != 0 && withPsString[0] == false {
		return nil
	} else {
		_, _ = c.currentWriter.Write([]byte(c.GetPsString()))
		return nil
	}
}

// Refresh clears basic properties for each evaluation of lexicalScopes.
// For bufio.Scanner, this methods clears properties related to every registered
// recognized lexicalScopes at the head of Parse.
func (c *CommandLine) Refresh() {
	c.lexicalScope = []string{}
	c.command = []*exec.Cmd{}
	c.expression = [][]string{}
	c.error = []error{}
	c.pipeSet = []Pipe{}
	c.exprIndex = 0
}
