package builtins

import (
	"testing"
)

func TestPrintfFailsWithNoArgument(t *testing.T) {
	if err := Printf(); err == nil {
		t.Fatal("printf must make error on no arguments")
	}
}

func TestPrintfFailsWithSingleArgument(t *testing.T) {
	if err := Printf("a"); err == nil {
		t.Fatal("printf must make error on single argument")
	}
}
