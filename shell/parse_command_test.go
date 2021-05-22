package shell

import (
	. "gopkg.in/check.v1"
)

func (s *CommandLine) TestCommandLine_ParseCommand(c *C) {
	s.parseCommand("ls", []string{"-l", "-n"})
	if len(s.error) != 0 {
		c.Fail()
	} else if len(s.lexicalScope) != 0 {
		c.Fail()
	} else if len(s.expression) != 0 {
		c.Fail()
	} else if len(s.command) == 0 {
		c.Fail()
	} else if len(s.pipeSet) == 0 {
		c.Fail()
	}
}
