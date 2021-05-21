package shell

//SetDebug is write-only setter method for CommandLine.debug flag.
func (c *CommandLine) SetDebug(flag bool) {
	c.debug = flag
}
