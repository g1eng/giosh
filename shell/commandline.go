package shell

import (
	"fmt"
	"os/exec"
)

// TerminateLine makes line termination with printing or not printing PS string
// Also any errors recorded in error stack will be shown by this method.
// FIXME: return value maybe not necessary, plan to remove it.
// Historically this method have a return value with error type.
// This is for DumpError, which have returned error object in old code
func (c *CommandLine) TerminateLine(withPsString ...bool) (err error) {
	// for bufio.Writer, write PS string
	err = c.DumpErrors()
	if err != nil {
		fmt.Println(err)
	}
	if len(withPsString) != 0 {
		if withPsString[0] == false {
			return nil
		}
	}
	_, _ = c.currentWriter.Write([]byte(c.GetPsString()))
	return nil
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
