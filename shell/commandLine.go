package shell

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
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
	pipe         []PipeIO
	stream       StreamIO
	tmpIndex     int
	error        []error
	debug        bool
	input        string
}

type PipeIO struct {
	stdin  io.WriteCloser
	stdout io.ReadCloser
}

var lineNo = 1

// GetPsString returns PS shell description for bufio.Writer
func GetPsString() string {
	return os.Getenv("USER") + "@G[" + strconv.Itoa(lineNo) + "]> "
}

// trimExpressionHead trims blank string from the head of the expression array
//It returns processed expression
func trimExpressionHead(expr []string) []string {
	if expr[0] == "" {
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

// track is internal error recording function for CommandLine.
// It records error given in argument into CommandLine.error if the argument is not nil
func (c *CommandLine) track(e error) {
	if e != nil {
		c.error = append(c.error, e)
	}
}

// DumpErrors is error reporting function for CommandLine.
// It scans any error object in CommandLine.error array and returns bool if it contains not a nil value
func (c *CommandLine) DumpErrors() (isNotNilArray error) {
	isNotNilArray = nil
	for i := range c.error {
		if c.error[i] != nil {
			if isNotNilArray == nil {
				isNotNilArray = c.error[i]
			}
			if c.debug {
				fmt.Println(c.error[i])
			}
		}
	}
	return isNotNilArray
}

//SetDebug is write-only setter method for CommandLine.debug flag.
func (c *CommandLine) SetDebug(flag bool) {
	c.debug = flag
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
		stdin    io.WriteCloser
		stdout   io.ReadCloser
		err      error
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
	stdin, err = c.cmd[tmpIndex].StdinPipe()
	c.track(err)
	stdout, err = c.cmd[tmpIndex].StdoutPipe()
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

func (c *CommandLine) getCurrentStdout() io.ReadCloser {
	if len(c.pipe) == 0 {
		c.track(errors.New("c.pipe is zero length"))
		return nil
	} else {
		return c.pipe[len(c.pipe)-1].stdout
	}
}

// isPipeEnd detects whether the pipe is end or not and returns bool value
func (c *CommandLine) isPipeEnd() bool {
	if c.tmpIndex == len(c.lexicalScope)-1 && c.debug {
		log.Println("pipe end")
	}
	return c.tmpIndex == len(c.lexicalScope)-1
}

// isBlankLine detects whether the input line is filled with blank or spaces,
// and returns bool value
func (c *CommandLine) isBlankLine() bool {
	if len(c.lexicalScope) == 0 {
		if c.debug {
			log.Println("blank line")
		}
		return true
	} else if len(c.expression) == 0 {
		if c.debug {
			log.Println("blank line")
		}
		return true
	} else if len(c.expression) != 0 && len(c.expression[0]) == 0 {
		if c.debug {
			log.Println("blank line")
		}
		return true
	}
	return false

}

func (c *CommandLine) setStatement() error {
	var (
		cmdName string
		args    []string
	)
	for i := range c.lexicalScope {
		c.tmpIndex = i
		if c.debug {
			log.Printf("lexicalScope[%d]: %v", i, c.lexicalScope[i])
		}

		c.setExpression(c.lexicalScope[i])
		if i == 0 && c.isBlankLine() {
			return c.TerminateLine()
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
		if c.debug {
			log.Printf("expression[%d]: %v", i, c.expression[i])
			for j := range c.expression[i] {
				log.Printf("expression[%d][%d]: %v", i, j, c.expression[i][j])
			}
			log.Printf("cmdName %d: %s", i, cmdName)
			log.Printf("args %d: %v", i, args)
		}

		c.track(c.cmd[i].Start())
	}
	return nil
}

func (c *CommandLine) evaluateStatement(stmt string) error {
	var (
		copySrc  io.Reader
		copyDest io.Writer
	)
	for i := range c.lexicalScope {
		if i == 0 {
			copySrc = bytes.NewBufferString(stmt)
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
	return nil
}

// Parse parses giosh command line and return output string
func (c *CommandLine) Parse() (err error) {
	c.Refresh()
	c.lexicalScope = strings.Split(c.GetInput(), "|")

	if err = c.setStatement(); err != nil {
		return err
	}
	if c.isBlankLine() {
		return c.TerminateLine()
	} else if err = c.evaluateStatement(c.GetInput()); err != nil {
		c.track(err)
	}

	for i := range c.cmd {
		c.track(c.cmd[i].Wait())
	}

	return c.TerminateLine()
}

// SetInput sets raw input string for c.input.
// (for reading scripts from arguments)
func (c *CommandLine) SetInput(input string) {
	c.input = input
}

// GetInput get raw input string from c.input.
// (just for temporary caching for file reading)
func (c *CommandLine) GetInput() string {
	return c.input
}

// Exec is the entry point for giosh shell interface.
//
// If you simply run it without any additional reader/writer sets,
// the vm starts with default I/O interface (bufio.Scanner and os.Stdout).
func (c *CommandLine) Exec() error {
	var scanner *bufio.Scanner
	c.Initialize()
	fmt.Printf(GetPsString())
	c.stream.buf.writer = append(c.stream.buf.writer, bufio.NewWriter(os.Stdout))
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

func (c *CommandLine) TerminateLine() (err error) {
	// for bufio.Writer, write PS string
	err = c.DumpErrors()
	if err != nil {
		fmt.Println(err)
	}
	for i := range c.stream.buf.writer {
		if _, err = c.stream.buf.writer[i].Write([]byte(GetPsString())); err != nil {
			return err
		} else if err = c.stream.buf.writer[i].Flush(); err != nil {
			return err
		}
	}
	return nil
}

func (c *CommandLine) Initialize() {
	c.stream = StreamIO{
		buf:  SystemIO{},
		file: []*os.File{},
		rest: Rest{},
	}
	c.Refresh()
}

// Refresh clears basic properties for each evaluation of lexicalScopes.
// For bufio.Scanner, this methods clears properties related to every registered
// recognized lexicalScopes at the head of Parse.
func (c *CommandLine) Refresh() {
	c.lexicalScope = []string{}
	c.cmd = []*exec.Cmd{}
	c.expression = [][]string{}
	c.error = []error{}
	c.pipe = []PipeIO{}
	c.tmpIndex = 0
}
