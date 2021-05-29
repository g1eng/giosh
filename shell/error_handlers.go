package shell

import (
	"fmt"
	"os"
)

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
	for _, e := range c.error {
		if e != nil {
			if isNotNilArray == nil {
				isNotNilArray = e
			}
			if c.Debug {
				_, _ = fmt.Fprintf(os.Stderr, "%v", e)
			}
		}
	}
	return isNotNilArray
}
