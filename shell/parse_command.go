package shell

import (
	"io"
	"os/exec"
)

// parseCommand sets exec.Command object for shell.CommandLine struct.
// Parsing a lexical scope, Commandline.Parse finally registers a command
// and its stdin/stdout pipe with this function.
func (c *CommandLine) parseCommand(cmdName string, args []string) {
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
