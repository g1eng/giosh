package builtins

import (
	"github.com/g1eng/giosh/builtins"
	"testing"
)

func TestPrintfFailsWithNoArgument(t *testing.T) {
	if err := builtins.Printf(); err == nil {
		t.Fatal("printf must make error on no arguments")
	}
}

func TestPrintfFailsWithSingleArgument(t *testing.T) {
	if err := builtins.Printf("a"); err == nil {
		t.Fatal("printf must make error on single argument")
	}
}
