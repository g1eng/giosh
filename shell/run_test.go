package shell

import ch "gopkg.in/check.v1"

func (s *CommandLine) TestCommandLine_Run(c *ch.C) {
	sh := New()
	if err := sh.Run(); err != nil {
		c.Fatal(err)
	}
}
