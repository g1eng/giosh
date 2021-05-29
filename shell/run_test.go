package shell

import ch "gopkg.in/check.v1"

func (s *CommandLine) TestCommandLine_RunExitWithNil(c *ch.C) {
	sh := New()
	if err := sh.Run(); err != nil {
		c.Fatal(err)
	}
}

func (s *CommandLine) TestCommandLine_RunWithInput(c *ch.C) {
	sh := New()
	sh.input = "ls -l /"
	if err := sh.Run(); err != nil {
		c.Fatal(err)
	}
}
