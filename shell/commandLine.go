package shell

import (
	"errors"
	gioParser "github.com/g1eng/giop/core"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type CommandLine struct {
	cmd        []*exec.Cmd
	expression [][]string
	error      []error
	pipe       []PipeIO
}

type PipeIO struct {
	stdin  io.WriteCloser
	stdout io.ReadCloser
}

var lineNo = 1

// GetPsString prints PS shell description
func GetPsString() string {
	return os.Getenv("USER") + "@G[" + strconv.Itoa(lineNo) + "]> "
}

//setExpression sets shell expression with IFS
//this function is applied to single lexicalScope
func (c *CommandLine) setExpression(lex string) {
	expr := regexp.MustCompilePOSIX("[ \\t]").Split(lex, -1)
	expr = trimExpression(expr) //trim line-head space characters
	c.expression = append(c.expression, expr)
}

// trimExpressionHead trims blank string from the head of the expression array
//It returns processed expression
func trimExpressionHead(expr []string) []string {
	if expr[0] == "" {
		for i := range expr[0] {
			log.Printf("expr[%d]: %x", i, []byte(expr[i]))
		}
		expr = expr[1:]
		expr = trimExpression(expr)
	}
	return expr
}

//trimExpressionTail trims blank string from the end of the expression array.
//It returns processed expression
func trimExpressionTail(expr []string) []string {
	for i := len(expr) - 1; i >= 0; i-- {
		if expr[i] == "" {
			expr = expr[:i]
		} else {
			return expr
		}
	}
	return expr
}

// trimExpression is the wrapper for trimExpressionTail and trimExpressionHead.
// It trims blank string "" from head and end of the given expression
func trimExpression(expr []string) []string {
	if len(expr) == 0 {
		return []string{""}
	}
	expr = trimExpressionHead(expr)
	expr = trimExpressionTail(expr)
	return expr
}

// registerCommand set exec.Command object for shell.CommandLine struct.
func (c *CommandLine) registerCommand(cmdName string, args []string) {
	var cmd *exec.Cmd
	if args == nil || args[0] == "" {
		cmd = exec.Command(cmdName)
		c.cmd = append(c.cmd, cmd)
	} else {
		cmd = exec.Command(cmdName, args...)
		c.cmd = append(c.cmd, cmd)
	}
	lineNo++
	c.pipe = append(c.pipe, PipeIO{})
}

func (c *CommandLine) getCurrentCommand() *exec.Cmd {
	if len(c.cmd) == 0 {
		c.error = append(c.error, errors.New("c.cmd is zero length"))
		return nil
	} else {
		return c.cmd[len(c.cmd)-1]
	}
}

// Exec is a ParserFunc, which returns the result string of the command execution
func (c *CommandLine) Exec(_ *gioParser.GioParser, s string) (string, error) {
	var (
		cmdName      string
		args         []string
		lexicalScope []string
	)
	c.Flush()
	lexicalScope = strings.Split(s, "|")
	if len(lexicalScope) == 0 {
		return GetPsString(), nil
	} else {
		c.setExpression(lexicalScope[0])
	}
	if len(c.expression[0]) == 0 {
		return GetPsString(), nil
	} else if len(c.expression[0][0]) == 1 {
		cmdName = c.expression[0][0]
	} else {
		cmdName = c.expression[0][0]
		args = c.expression[0][1:]
		args = trimExpression(args) //trim line-head space characters
		for i := range c.expression[0] {
			log.Printf("expression[0][%d]: %s", i, c.expression[0][i])
		}
		for i := range args {
			log.Printf("args[%d]: %s", i, args[i])
		}
	}

	c.registerCommand(cmdName, args)

	originOutput, err := c.getCurrentCommand().Output()
	c.error = append(c.error, err)

	if len(lexicalScope) == 1 {
		c.WriteTo(os.Stdout, originOutput)
		return GetPsString(), err
	}

	for i := range lexicalScope {
		log.Printf("lexicalScope[%d]: %v", i, lexicalScope[i])
		if i != 0 {
			c.setExpression(lexicalScope[i])
			tmpIndex := len(c.expression) - 1
			cmdName = c.expression[tmpIndex][0]
			args = c.expression[tmpIndex][1:]
			args = trimExpression(args) //trim line-head space characters
			log.Printf("expression[%d]: %v", i, c.expression[i])
			for j := range c.expression[i] {
				log.Printf("expression[%d][%d]: %v", i, j, c.expression[i][j])
			}

			c.registerCommand(cmdName, args)
			cmd2 := c.getCurrentCommand()

			stdin2, _ := cmd2.StdinPipe()
			c.WriteTo(stdin2, originOutput)
			err = stdin2.Close()

			processOutput, _ := cmd2.Output()
			c.WriteTo(os.Stdout, processOutput)
		}
	}
	for i := range c.error {
		if c.error[i] != nil {
			return GetPsString(), err
		}
	}
	c.WriteTo(os.Stdout, []byte(GetPsString()))

	return GetPsString(), nil
}

func (c *CommandLine) WriteTo(dest io.WriteCloser, output []byte) {
	_, err := io.WriteString(dest, string(output))
	c.error = append(c.error, err)
}

func (c *CommandLine) Flush() {
	c.cmd = []*exec.Cmd{}
	c.expression = [][]string{}
	c.error = []error{}
	c.pipe = []PipeIO{}
}
