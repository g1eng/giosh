package shell

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

//parseStatement parses and sets a statement.
//This function is applied to a lexicalScope,
//where each statements are separated with "|".
func (c *CommandLine) parseStatement() error {
	var (
		cmdName string
		args    []string
	)
	for i := range c.lexicalScope {

		c.parseExpression(c.lexicalScope[i])
		if c.isBlankLine() {
			return nil
		} else if len(c.expression[i]) == 0 {
			return fmt.Errorf("pipe error: pipeline #%d is nil\n", i+1)
		}

		// command with no arg
		if len(c.expression[i][0]) == 1 {
			cmdName = c.expression[i][0]
		} else { //command with arguments
			cmdName = c.expression[i][0]
			args = c.expression[i][1:]
			args = trimExpression(args)
		}

		c.parseCommand(cmdName, args)

		if c.Debug {
			c.dumpParserObject(i)
		}

		c.track(c.command[i].Start())
	}
	return nil
}

//evaluateStatement reads a statement from pipeline input.
//This function is applied to each `statement`, which is
//separated with `|`
func (c *CommandLine) evaluateStatement(stmt string) {
	var copySrc io.Reader
	for i := range c.lexicalScope {
		c.exprIndex = i
		if i == 0 {
			copySrc = bytes.NewBufferString(stmt)
		} else {
			copySrc = c.pipeSet[i-1].stdout
		}
		if c.isPipeEnd() {
			c.currentWriter = os.Stdout
		} else {
			c.currentWriter = c.pipeSet[i+1].stdin
		}
		_, err := io.Copy(c.pipeSet[i].stdin, copySrc)
		c.track(err)
		c.track(c.pipeSet[i].stdin.Close())
		_, err = io.Copy(c.currentWriter, c.pipeSet[i].stdout)
		c.track(err)
	}
}
