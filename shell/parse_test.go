package shell

import . "gopkg.in/check.v1"

func (s *CommandLine) TestParse1(c *C) {
	s.input = "ls -l"
	c.Check(s.Parse(), IsNil)
}
