package shell

import "fmt"

// track is internal error recording function for CommandLine.
// It records error given in argument into CommandLine.error if the argument is not nil
func (c *CommandLine) track(e error) {
	if e != nil {
		c.error = append(c.error, e)
	}
}

// DumpErrors is error reporting function for CommandLine.
// It scans any error object in CommandLine.error array and returns bool if it contains not a nil value
func (c *CommandLine) DumpErrors() (isNotNilArray error) {
	isNotNilArray = nil
	for i := range c.error {
		if c.error[i] != nil {
			if isNotNilArray == nil {
				isNotNilArray = c.error[i]
			}
			if c.debug {
				fmt.Println(c.error[i])
			}
		}
	}
	return isNotNilArray
}
