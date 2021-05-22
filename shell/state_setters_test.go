package shell

import (
	. "gopkg.in/check.v1"
	"testing"
)

func init() {
	Suite(&CommandLine{})
}

func Test(t *testing.T) { TestingT(t) }

func (s *CommandLine) TestCommandLine_SetDebug(c *C) {
	s.SetDebug(true)
	if s.Debug != true {
		c.Fatal("cmd.Debug must be true after setDebug(true)")
	}
}

func (s *CommandLine) TestCommandLine_SetDebugFalse(c *C) {
	s.SetDebug(false)
	if s.Debug != false {
		c.Fatal("cmd.Debug must be false after setDebug(false)")
	}
}

func (s *CommandLine) TestCommandLine_SetInput(c *C) {
	s.SetInput("a")
	if s.input != "a" {
		c.Fatal("CommandLine.input must be set to set argument")
	}
}
