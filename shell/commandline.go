package shell

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

// CommandLine is the main struct for giosh shell instance.
// This struct has any internal state of a shell such as
// lexical scopes, expressions for the evaluation, pipes
// which internal I/O are processed, etc.
type CommandLine struct {
	lexicalScope  []string
	expression    [][]string
	command       []*exec.Cmd
	pipeSet       []Pipe
	io            IOStream
	exprIndex     int
	error         []error
	debug         bool
	input         string
	lineno        int
	currentWriter io.Writer
}

// Pipe is input and output pipe for the shell.
// This is used to store pipe objects derived from
// methods on *exec.Cmd.
type Pipe struct {
	stdin  io.WriteCloser
	stdout io.ReadCloser
}

// registerCommand sets exec.Command object for shell.CommandLine struct.
func (c *CommandLine) registerCommand(cmdName string, args []string) {
	var (
		exprIndex int
		cmd       *exec.Cmd
		stdin     io.WriteCloser
		stdout    io.ReadCloser
		err       error
	)
	exprIndex = len(c.command)
	if args == nil || args[0] == "" {
		cmd = exec.Command(cmdName)
		c.command = append(c.command, cmd)
	} else {
		cmd = exec.Command(cmdName, args...)
		c.command = append(c.command, cmd)
	}
	c.lineno++
	stdin, err = c.command[exprIndex].StdinPipe()
	c.track(err)
	stdout, err = c.command[exprIndex].StdoutPipe()
	c.track(err)
	c.pipeSet = append(c.pipeSet, Pipe{
		stdin:  stdin,
		stdout: stdout,
	})
}

//setStatement parses and sets a statement.
//This function is applied to a lexicalScope,
//where each statements are separated with "|".
func (c *CommandLine) setStatement() error {
	var (
		cmdName string
		args    []string
	)
	for i := range c.lexicalScope {
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

		c.track(c.command[i].Start())
	}
	return nil
}

//evaluateStatement reads a statement from pipeline input.
//This function is applied to each `statement`, which is
//separated with `|`
func (c *CommandLine) evaluateStatement(stmt string) error {
	var copySrc io.Reader
	for i := range c.lexicalScope {
		if i == 0 {
			copySrc = bytes.NewBufferString(stmt)
		} else {
			copySrc = c.pipeSet[i-1].stdout
		}
		if i == len(c.command)-1 {
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
	return nil
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

// setLexicalScope sets all lexical scope on the line terminated with EOF.
// At now giosh does not have multi-line scope for the lexical parsing.
//
// Each lexical scope within a line, is just separated with "|"
func (c *CommandLine) setLexicalScope() {
	c.lexicalScope = strings.Split(c.GetInput(), "|")
}

// Parse parses giosh command line and return output string.
// At now, any lexical scopes are separated with "|" and
// any lines are separated with EOF. (multi-line is not allowed)
func (c *CommandLine) Parse() (err error) {
	c.Refresh()
	c.setLexicalScope()
	if err = c.setStatement(); err != nil {
		return err
	}
	if c.isBlankLine() {
		return c.TerminateLine(false)
	} else if err = c.evaluateStatement(c.GetInput()); err != nil {
		c.track(err)
	}
	for i := range c.command {
		c.track(c.command[i].Wait())
	}

	return c.TerminateLine()
}

// New generates a new shell instance with default parameters.
func New() CommandLine {
	newShell := CommandLine{}
	newShell.lineno = 1
	newShell.io = IOStream{
		writer: []io.Writer{},
		reader: []io.Reader{},
	}
	newShell.Refresh()
	return newShell
}

// Exec is the entry point for giosh shell interface.
//
// If you simply run it without any additional reader/writer sets,
// the vm starts with default I/O interface (bufio.Scanner and os.Stdout).
func (c *CommandLine) Exec() error {
	var scanner *bufio.Scanner
	c.currentWriter = os.Stdout
	fmt.Printf(c.GetPsString())
	c.io.writer = append(c.io.writer, bufio.NewWriter(os.Stdout))
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
