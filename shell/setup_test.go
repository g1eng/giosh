package shell

import (
	. "gopkg.in/check.v1"
	"os"
	"testing"
)

func init() {
	Suite(&CommandLine{
		currentWriter: os.Stdout,
		error:         []error{},
	})
}

func Test(t *testing.T) { TestingT(t) }
