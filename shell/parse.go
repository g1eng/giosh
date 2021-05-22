package shell

// Parse parses giosh command line and return output string.
// At now, any lexical scopes are separated with "|" and
// any lines are separated with EOF. (multi-line is not allowed)
func (c *CommandLine) Parse() (err error) {
	c.Refresh()
	c.parseLexicalScope()
	if err = c.parseStatement(); err != nil {
		return err
	}
	if c.isBlankLine() {
		return c.TerminateLine(false)
	} else if err = c.evaluateStatement(c.GetInput()); err != nil {
		c.track(err)
	}
	for i := range c.command {
		c.track(c.command[i].Wait())
	}

	return c.TerminateLine()
}
