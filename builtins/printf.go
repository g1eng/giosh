package builtins

import (
	"errors"
	"fmt"
)

//Printf builtin
func Printf(desc ...string) error {
	if len(desc) < 2 {
		return errors.New("too few arguments")
	} else {
		format := desc[0]
		desc = desc[1:]
		_, err := fmt.Printf(format, desc)
		return err
	}
}
