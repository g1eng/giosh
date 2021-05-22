package shell

import . "gopkg.in/check.v1"

func (s *CommandLine) TestCommandLine_LexicalScope(c *C) {
	sh := New()
	sh.lexicalScope = append(sh.lexicalScope, "ls -l -n")
	sh.parseLexicalScope()
	if len(sh.lexicalScope) != 1 {
		c.Fail()
	}
}
