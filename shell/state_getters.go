package shell

import (
	"errors"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
)

// GetPsString returns PS shell description for bufio.Writer
func (c *CommandLine) GetPsString() string {
	return os.Getenv("USER") + "@G[" + strconv.Itoa(c.lineno) + "]> "
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
