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

type CommandLine struct {
	lexicalScope []string
	expression   [][]string
	cmd          []*exec.Cmd
	pipe         []Pipe
	stream       StreamIO
	tmpIndex     int
	error        []error
	debug        bool
	input        string
	lineno       int
}

type Pipe struct {
	stdin  io.WriteCloser
	stdout io.ReadCloser
}

// registerCommand sets exec.Command object for shell.CommandLine struct.
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
	c.lineno++
	stdin, err = c.cmd[tmpIndex].StdinPipe()
	c.track(err)
	stdout, err = c.cmd[tmpIndex].StdoutPipe()
	c.track(err)
	c.pipe = append(c.pipe, Pipe{
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

func (c *CommandLine) TerminateLine() (err error) {
	// for bufio.Writer, write PS string
	err = c.DumpErrors()
	if err != nil {
		fmt.Println(err)
	}
	for i := range c.stream.buf.writer {
		if _, err = c.stream.buf.writer[i].Write([]byte(c.GetPsString())); err != nil {
			return err
		} else if err = c.stream.buf.writer[i].Flush(); err != nil {
			return err
		}
	}
	return nil
}

// Refresh clears basic properties for each evaluation of lexicalScopes.
// For bufio.Scanner, this methods clears properties related to every registered
// recognized lexicalScopes at the head of Parse.
func (c *CommandLine) Refresh() {
	c.lexicalScope = []string{}
	c.cmd = []*exec.Cmd{}
	c.expression = [][]string{}
	c.error = []error{}
	c.pipe = []Pipe{}
	c.tmpIndex = 0
}

// Parse parses giosh command line and return output string.
// At now, any lexical scopes are separated with "|" and
// any statements are separated with EOF.
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

// New generates a new shell instance with default parameters.
func New() CommandLine {
	giosh := CommandLine{}
	giosh.lineno = 1
	giosh.stream = StreamIO{
		buf:  SystemIO{},
		file: []*os.File{},
		rest: Rest{},
	}
	giosh.Refresh()
	return giosh
}

// Exec is the entry point for giosh shell interface.
//
// If you simply run it without any additional reader/writer sets,
// the vm starts with default I/O interface (bufio.Scanner and os.Stdout).
func (c *CommandLine) Exec() error {
	var scanner *bufio.Scanner
	fmt.Printf(c.GetPsString())
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
