package shell

import (
	. "gopkg.in/check.v1"
	"testing"
)

func init() {
	Suite(&CommandLine{})
}

func Test(t *testing.T) { TestingT(t) }
