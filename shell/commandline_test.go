package shell

import (
	"fmt"
	. "gopkg.in/check.v1"
)

func (s *CommandLine) TestTerminateLineWithoutPS(c *C) {
	// fatal error made with
	c.Check(s.TerminateLine(), IsNil)
}

func (s *CommandLine) TestTerminateLineWithPS(c *C) {
	c.Check(s.TerminateLine(false), IsNil)
}

func (s *CommandLine) TestTerminateLineWithError(c *C) {
	e := fmt.Errorf("err")
	s.error = []error{e}
	c.Check(s.TerminateLine(), IsNil)
}
