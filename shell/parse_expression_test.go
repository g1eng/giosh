package shell

import (
	. "gopkg.in/check.v1"
)

func (s *CommandLine) TestCommandLine_ParseExpression(c *C) {
	sh := New()
	if len(sh.expression) != 0 {
		c.Fail()
	}
	sh.lexicalScope = append(sh.lexicalScope, "ls -l -n")
	sh.parseExpression(sh.lexicalScope[0])
	if len(sh.expression) != 1 {
		c.Fail()
	} else if sh.expression[0][0] != "ls" {
		c.Fail()
	} else if len(sh.expression[0]) != 3 {
		c.Fatal("expression length must be 3 for `ls -l -n`")
	} else if sh.expression[0][1] != "-l" {
		c.Fail()
	} else if sh.expression[0][2] != "-n" {
		c.Fail()
	}
}

func (s *CommandLine) TestTrimExpressionWithBlankSlice(c *C) {
	out := trimExpression([]string{})
	c.Check(len(out), Equals, 1)
}

func (s *CommandLine) TestCommandLine_TrimExpressionWithBlankPadding(c *C) {
	ex := []string{"\t\t", "  ", "ls", "  ", "-l"}
	out := trimExpression(ex)
	c.Check(len(out), Equals, 2)
}
