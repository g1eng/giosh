package builtins

import . "gopkg.in/check.v1"

func (s *TestSuite) TestEchoNoArg(c *C) {
	if !c.Check(echo(), IsNil) {
		c.Fatal("echo with no arguments should exit without failure")
	}
}

func (s *TestSuite) TestEchoSingle(c *C) {
	if !c.Check(echo("ok"), IsNil) {
		c.Fatal("echo with single valid string should exit without failure")
	}
}

func (s *TestSuite) TestEchoMulti(c *C) {
	if !c.Check(echo("this", "is", "ok"), IsNil) {
		c.Fatal("echo with three valid strings should exit without failure")
	}
}

func (s *TestSuite) TestEchoNoNewline(c *C) {
	if !c.Check(echo("-n", "this", "is", "ok"), IsNil) {
		c.Fatal("echo without newline should exit without failure")
	}
}
