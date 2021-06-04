package shell

import ch "gopkg.in/check.v1"

func (s *CommandLine) TestRunExitWithNil(c *ch.C) {
	sh := New()
	if err := sh.Run(); err != nil {
		c.Fatal(err)
	}
}

func (s *CommandLine) TestRunWithInput(c *ch.C) {
	sh := New()
	sh.input = "ls -l /"
	if err := sh.Run(); err != nil {
		c.Fatal(err)
	}
}

//iss #1
func (s *CommandLine) TestRunWithUnterminatedPipeError(c *ch.C) {
	sh := New()
	sh.input = "ls -l |"
	if err := sh.Run(); err != nil {
		c.Fatal(err)
	}
}
