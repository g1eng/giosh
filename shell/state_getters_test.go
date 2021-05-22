package shell

import (
	. "gopkg.in/check.v1"
)

func (s *CommandLine) SetUpSuite(_ *C) {
	s.input = "ok"
}

func (s *CommandLine) TestCommandLine_GetInput(c *C) {
	if s.GetInput() != "ok" {
		c.Fatal("not ok")
	}
}
