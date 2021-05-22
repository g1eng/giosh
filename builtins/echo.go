package builtins

import (
	"fmt"
	"strings"
)

//echo is `echo` builtin for giosh
func echo(desc ...string) (err error) {
	if len(desc) == 0 {
		println("")
		return nil
	} else if len(desc) >= 2 && desc[0] == "-n" {
		desc = desc[1:]
		post := strings.Join(desc, " ")
		_, err = fmt.Printf("%s", post)
		return err
	} else {
		return err
	}
}
