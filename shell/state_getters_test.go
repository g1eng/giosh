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

func (s *CommandLine) TestBlankLine(c *C) {
	s.lexicalScope = []string{"ls"}
	s.expression = [][]string{{"ls"}}
	c.Log("ls")
	c.Check(s.isBlankLine(), Equals, false)
	c.Log("[nospace]")
	s.lexicalScope = []string{}
	s.expression = [][]string{}
	c.Check(s.isBlankLine(), Equals, true)
	c.Log("[spaceX2]")
	s.lexicalScope = []string{"  "}
	s.expression = [][]string{{}}
	c.Check(s.isBlankLine(), Equals, true)
}
