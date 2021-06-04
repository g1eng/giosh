package shell

import (
	. "gopkg.in/check.v1"
)

func (s *CommandLine) TestCommandLine_ParseCommandWithArgs(c *C) {
	lexLength := len(s.lexicalScope)
	expLength := len(s.expression)
	s.parseCommand("ls", []string{"-l", "-n"})
	if len(s.error) != 0 {
		c.Fatal("parseCommand must not append new error object")
	} else if len(s.lexicalScope) != lexLength {
		c.Fatal("parseCommand must not change lexical scope length")
	} else if len(s.expression) != expLength {
		c.Fatal("parseCommand must not change expression length")
	} else if len(s.command) == 0 {
		c.Fatal("parseCommand must set a new command field")
	} else if len(s.pipeSet) == 0 {
		c.Fatal("parseCommand must generate a new pipeSet field")
	}
}

func (s *CommandLine) TestCommandLine_ParseCommandWithoutArgs(c *C) {
	s.parseCommand("ls", []string{""})
	if len(s.error) != 0 {
		c.Fatal("parseCommand must not append new error object")
	} else if len(s.command) == 0 {
		c.Fatal("parseCommand must set a new command field")
	}
}
