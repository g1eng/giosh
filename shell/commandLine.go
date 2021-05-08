package shell

import (
	"errors"
	gioParser "github.com/g1eng/giop/core"
	"io"
	//"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type CommandLine struct {
	lexicalScope []string
	expression   [][]string
	cmd          []*exec.Cmd
	Error        []error
	pipe         []PipeIO
	tmpIndex     int
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

// trimExpressionHead trims blank string from the head of the expression array
//It returns processed expression
func trimExpressionHead(expr []string) []string {
	if expr[0] == "" {
		//for i := range expr[0] {
		//	log.Printf("expr[%d]: %x", i, []byte(expr[i]))
		//}
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

// track: error recording function for CommandLine
// It records error given in argument into CommandLine.Error if the argument is not nil
func (c *CommandLine) track(e error) {
	if e != nil {
		c.Error = append(c.Error, e)
	}
}

//setExpression sets shell expression with IFS
//this function is applied to single lexicalScope
func (c *CommandLine) setExpression(lex string) {
	expr := regexp.MustCompilePOSIX("[ \\t]").Split(lex, -1)
	expr = trimExpression(expr) //trim line-head space characters
	c.expression = append(c.expression, expr)
}

// registerCommand set exec.Command object for shell.CommandLine struct.
func (c *CommandLine) registerCommand(cmdName string, args []string) {
	var (
		tmpIndex int
		cmd      *exec.Cmd
	)
	tmpIndex = len(c.cmd)
	if args == nil || args[0] == "" {
		cmd = exec.Command(cmdName)
		c.cmd = append(c.cmd, cmd)
	} else {
		cmd = exec.Command(cmdName, args...)
		c.cmd = append(c.cmd, cmd)
	}
	lineNo++
	stdin, err := c.cmd[tmpIndex].StdinPipe()
	c.track(err)
	stdout, err := c.cmd[tmpIndex].StdoutPipe()
	c.track(err)
	c.pipe = append(c.pipe, PipeIO{
		stdin:  stdin,
		stdout: stdout,
	})
}

// getCurrentCommand returns latest exec.Command registered in CommandLine.cmd
func (c *CommandLine) getCurrentCommand() *exec.Cmd {
	if len(c.cmd) == 0 {
		c.track(errors.New("c.cmd is zero length"))
		return nil
	} else {
		return c.cmd[len(c.cmd)-1]
	}
}

// isPipeEnd detects whether the pipe is end or not and returns bool value
func (c *CommandLine) isPipeEnd() bool {
	if c.tmpIndex == len(c.lexicalScope)-1 {
		//log.Println("pipe end")
	}
	return c.tmpIndex == len(c.lexicalScope)-1
}

// isBlankLine detects whether the input line is filled with blank or spaces,
// and returns bool value
func (c *CommandLine) isBlankLine() bool {
	if len(c.lexicalScope) == 0 {
		//log.Println("blank line")
		return true
	} else if len(c.expression) != 0 && len(c.expression[0]) == 0 {
		//log.Println("blank line")
		return true
	}
	return false
}

// Exec is a ParserFunc, which returns the result string of the command execution
func (c *CommandLine) Exec(_ *gioParser.GioParser, s string) (string, error) {
	var (
		cmdName  string
		args     []string
		copySrc  io.Reader
		copyDest io.Writer
	)
	c.Flush()
	c.lexicalScope = strings.Split(s, "|")

	for i := range c.lexicalScope {
		c.tmpIndex = i
		//log.Printf("lexicalScope[%d]: %v", i, c.lexicalScope[i])

		c.setExpression(c.lexicalScope[i])
		if i == 0 && c.isBlankLine() {
			return GetPsString(), nil
		}

		// command with no arg
		if len(c.expression[i][0]) == 1 {
			cmdName = c.expression[i][0]
		} else { //command with arguments
			cmdName = c.expression[i][0]
			args = c.expression[i][1:]
			args = trimExpression(args)
		}

		c.registerCommand(cmdName, args)

		// debugger
		//log.Printf("expression[%d]: %v", i, c.expression[i])
		//for j := range c.expression[i] {
		//	log.Printf("expression[%d][%d]: %v", i, j, c.expression[i][j])
		//}
		//log.Printf("cmdName %d: %s", i, cmdName)
		//log.Printf("args %d: %v", i, args)

		c.track(c.cmd[i].Start())
	}

	for i := range c.lexicalScope {
		if i == 0 {
			copySrc = os.Stdin
		} else {
			copySrc = c.pipe[i-1].stdout
		}
		if i == len(c.cmd)-1 {
			copyDest = os.Stdout
		} else {
			copyDest = c.pipe[i+1].stdin
		}
		_, err := io.Copy(c.pipe[i].stdin, copySrc)
		c.track(err)
		c.track(c.pipe[i].stdin.Close())
		_, err = io.Copy(copyDest, c.pipe[i].stdout)
		c.track(err)
	}

	for i := range c.cmd {
		c.track(c.cmd[i].Wait())
	}
	defer c.track(c.pipe[0].stdin.Close())
	_, err := c.pipe[0].stdin.Write([]byte("\n"))
	c.track(err)
	c.track(c.pipe[len(c.pipe)-1].stdout.Close())

	return GetPsString(), nil
}

func (c *CommandLine) WriteTo(dest io.WriteCloser, output []byte) {
	_, err := io.WriteString(dest, string(output))
	c.track(err)
}

func (c *CommandLine) Flush() {
	c.cmd = []*exec.Cmd{}
	c.expression = [][]string{}
	c.Error = []error{}
	c.pipe = []PipeIO{}
}
