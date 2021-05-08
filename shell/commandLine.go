package shell

import (
	"fmt"
	gioParser "github.com/g1eng/giop/core"
	"os"
	"os/exec"
	"strconv"
)

var lineNo = 1

// GetPsString prints PS shell description
func GetPsString() string {
	return os.Getenv("USER") + "@G[" + strconv.Itoa(lineNo) + "]> "
}

func setExpression(expression []string) (string, []string) {
	cmdName := expression[0]
	var args []string
	if len(expression) == 1 { //no arg
		args = []string{""}[:0]
	} else {
		args = expression[1:]
	}
	return cmdName, args
}

// Exec is a ParserFunc, which returns the result string of the command execution
func Exec(g *gioParser.GioParser, cmdLine string) (string, error) {
	//parse expression with default IFS " "
	expr := g.Split(cmdLine, -1)
	//parse given expression ant get a command name and arguments
	cmdName, args := setExpression(expr)

	if cmdName == "" {
		// for single newline
		return GetPsString(), nil
	} else {
		// for command execution
		b, err := exec.Command(cmdName, args...).Output()
		lineNo++
		if err != nil {
			fmt.Printf("giosh %s\n", err)
		}
		return string(b) + GetPsString(), nil
	}
}
