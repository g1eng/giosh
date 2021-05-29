package builtins

import (
	. "gopkg.in/check.v1"
)

func (s TestSuite) TestPrintfFailsWithNoArgument(c *C) {
	if !c.Check(Printf(), NotNil) {
		c.Fatal("printf must make error on no arguments")
	}
}

func (s *TestSuite) TestPrintfFailsWithSingleArgument(c *C) {
	if !c.Check(Printf("a"), NotNil) {
		c.Fatal("printf must make error on single argument")
	}
}
