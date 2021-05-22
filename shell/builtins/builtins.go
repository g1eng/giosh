package builtins

import (
	"fmt"
)

type builtinFunc func(args ...string) error

func echo(desc string) error {
	_, err := fmt.Printf("%s\n", desc)
	return err
}

func printf(desc string) error {
	_, err := fmt.Printf("%s", desc)
	return err
}
