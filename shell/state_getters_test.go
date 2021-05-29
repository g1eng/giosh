package shell

import (
	. "gopkg.in/check.v1"
	"os/exec"
)

func (s *CommandLine) SetUpSuite(_ *C) {
	s.input = "ok"
}

func (s *CommandLine) TestCommandLine_GetInput(c *C) {
	if s.GetInput() != "ok" {
		c.Fatal("not ok")
	}
}

func (s *CommandLine) TestGetCurrentCommand(c *C) {
	c1 := exec.Command("ls", "-l")
	c2 := exec.Command("df")
	s.command = []*exec.Cmd{c1, c2}
	c.Check(s.getCurrentCommand(), Equals, c2)
}

func (s *CommandLine) TestGetCurrentCommandNil(c *C) {
	s.command = []*exec.Cmd{}
	c.Check(s.getCurrentCommand(), IsNil)
}
