package builtins

import (
	. "gopkg.in/check.v1"
)

//printf must make error on no arguments
func (s TestSuite) TestPrintfFailsWithNoArgument(c *C) {
	c.Check(Printf(), NotNil)
}

// printf must make error on single argument
func (s *TestSuite) TestPrintfFailsWithSingleArgument(c *C) {
	c.Check(Printf("a"), NotNil)
}

func (s *TestSuite) TestPrintfWithValidFormat(c *C) {
	num := "123"
	c.Check(Printf("no a valid format %d", num), IsNil)
}

func (s *TestSuite) TestPrintfNotFailsWithInvalidFormat(c *C) {
	msg := "ok?"
	c.Check(Printf("no a valid format %1", msg), IsNil)
}
