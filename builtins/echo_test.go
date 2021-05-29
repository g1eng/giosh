package builtins

import . "gopkg.in/check.v1"

//echo with no arguments should exit without failure"
func (s *TestSuite) TestEchoNoArg(c *C) {
	c.Check(echo(), IsNil)
}

//echo with single valid string should exit without failure
func (s *TestSuite) TestEchoSingle(c *C) {
	c.Check(echo("ok"), IsNil)
}

//echo with three valid strings should exit without failure
func (s *TestSuite) TestEchoMulti(c *C) {
	c.Check(echo("this", "is", "ok"), IsNil)
}

//echo without newline should exit without failure
func (s *TestSuite) TestEchoNoNewline(c *C) {
	c.Check(echo("-n", "this", "is", "ok"), IsNil)
}
