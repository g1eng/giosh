package shell

import (
	"fmt"
	. "gopkg.in/check.v1"
)

func (s *CommandLine) TestDumpErrorsWithBlankArray(c *C) {
	s.error = []error{}
	c.Check(s.DumpErrors(), IsNil)
}

func (s *CommandLine) TestDumpErrorsWithNilArray(c *C) {
	s.error = []error{nil, nil}
	c.Check(s.DumpErrors(), IsNil)
}

func (s *CommandLine) TestDumpErrorWithNotANilArray(c *C) {
	s.error = []error{fmt.Errorf("dummy error")}
	c.Check(s.DumpErrors(), NotNil)
}
