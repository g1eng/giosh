package shell

import (
	"fmt"
)

func echo(desc string) error {
	_, err := fmt.Printf("%s\n", desc)
	return err
}

func test(expr string) error {
	return nil
}

func printf(desc string) error {
	_, err := fmt.Printf("%s", desc)
	return err
}