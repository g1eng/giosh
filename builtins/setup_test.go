package builtins

import (
	. "gopkg.in/check.v1"
	"testing"
)

type TestSuite struct{}

func init() {
	Suite(&TestSuite{})
}

func Test(t *testing.T) { TestingT(t) }
