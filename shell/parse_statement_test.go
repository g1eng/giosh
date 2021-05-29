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

func (s *CommandLine) TestCommandLine_ParseStatementWithBlankLine(c *C) {
	sh := New()
	sh.lexicalScope = []string{}
	if err := sh.parseStatement(); err != nil {
		c.Fatal(err)
	}
}
