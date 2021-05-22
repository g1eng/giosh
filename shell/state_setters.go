package shell

//SetDebug is write-only setter method for CommandLine.debug flag.
func (c *CommandLine) SetDebug(flag bool) {
	c.debug = flag
}

// SetInput sets raw input string for c.input.
// (for reading scripts from arguments)
func (c *CommandLine) SetInput(input string) {
	c.input = input
}
