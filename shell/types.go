package shell

import (
	sio "github.com/g1eng/giosh/io"
	"io"
	"os/exec"
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
	io            sio.Stream
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

// New generates a new shell instance with default parameters.
func New() CommandLine {
	newShell := CommandLine{}
	newShell.lineno = 1
	newShell.io = sio.Stream{
		Writer: []io.Writer{},
		Reader: []io.Reader{},
	}
	newShell.Refresh()
	return newShell
}
