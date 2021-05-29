package shell

import . "gopkg.in/check.v1"

func (s *CommandLine) TestCommandLine_ParseStatement(c *C) {
	sh := New()
	sh.lexicalScope = append(sh.lexicalScope, "ls -l -n")
	if err := sh.parseStatement(); err != nil {
		c.Fatal(err)
	} else if len(sh.expression) == 0 {
		c.Fatal("statement must be set")
	} else if len(sh.expression[0]) == 0 {
		c.Fatal("expression must be set")
	} else if sh.expression[0][0] != "ls" {
		c.Fatal("command must be `ls`")
	} else if sh.expression[0][1] != "-l" {
		c.Fatal("first option is -l")
	} else if sh.expression[0][2] != "-n" {
		c.Fatal("second option is -n")
	}
}

func (s *CommandLine) TestParseStatementWithBlankLine(c *C) {
	sh := New()
	sh.lexicalScope = []string{"", "ls -l"}
	sh.expression = [][]string{}
	sh.expression = append(s.expression, []string{})
	sh.expression = append(s.expression, []string{"ls", "-l"})
	c.Check(sh.parseStatement(), IsNil)
}

func (s *CommandLine) TestParseStatementWithBlankLine2(c *C) {
	sh := New()
	sh.lexicalScope = []string{"\n"}
	c.Check(sh.parseStatement(), IsNil)
}

func (s *CommandLine) TestEvaluateStatement(c *C) {
	sh := New()
	sh.lexicalScope = []string{"ls -l", "grep m"}
	_ = sh.parseStatement()
	sh.evaluateStatement("")
	c.Check(sh.DumpErrors(), IsNil)
}

func (s *CommandLine) TestEvaluateInvalidStatement(c *C) {
	sh := New()
	sh.lexicalScope = []string{"; ls", "grep m"}
	_ = sh.parseStatement()
	sh.evaluateStatement("")
	c.Check(sh.DumpErrors(), NotNil)
}

func (s *CommandLine) TestEvaluateInvalidCommand(c *C) {
	sh := New()
	sh.lexicalScope = []string{"ls", "notavalidcommand-g"}
	_ = sh.parseStatement()
	sh.evaluateStatement("")
	c.Check(sh.DumpErrors(), NotNil)
}
