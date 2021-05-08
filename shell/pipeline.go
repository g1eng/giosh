package shell

import (
	gioParser "github.com/g1eng/giop/core"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type PipeIO struct {
	command []*exec.Cmd
	stdin   []io.WriteCloser
	stdout  []io.ReadCloser
	result  [][]byte
	error   []error
}

func trimExpressionHead(expr []string) []string {
	if expr[0] == "" {
		for i := range expr[0] {
			log.Printf("expr[%d]: %x", i, []byte(expr[i]))
		}
		expr = expr[1:]
		expr = trimExpression(expr)
	}
	return expr
}
func trimExpressionTail(expr []string) []string {
	for i := len(expr) - 1; i >= 0; i-- {
		if expr[i] == "" {
			expr = expr[:i]
		} else {
			return expr
		}
	}
	return expr
}
func trimExpression(expr []string) []string {
	if len(expr) == 0 {
		return []string{""}
	}
	expr = trimExpressionHead(expr)
	expr = trimExpressionTail(expr)
	return expr
}

//PocPipeParser is a simple poc for line conversion with tr command.
//All lowercase characters for the first command output are converted to uppercase within a pipeline
func (_ *PipeIO) PocPipeParser(_ *gioParser.GioParser, s string) (string, error) {
	var (
		expression   [][]string
		cmdName      string
		args         []string
		lexicalScope []string
		cmd          []*exec.Cmd
	)
	lexicalScope = strings.Split(s, "|")
	if len(lexicalScope) != 0 {
		cmdLine := regexp.MustCompilePOSIX(" |\\t").Split(lexicalScope[0], -1)
		cmdLine = trimExpression(cmdLine) //trim line-head space characters
		expression = append(expression, cmdLine)
		cmdName = expression[0][0]
	} else {
		return GetPsString(), nil
	}
	if len(expression[0][0]) > 1 {
		args = expression[0][1:]
		args = trimExpression(args) //trim line-head space characters
		for i := range expression[0] {
			log.Printf("expression[0][%d]: %x", i, []byte(expression[0][i]))
		}
		for i := range args {
			log.Printf("args[%d]: %s", i, args[i])
		}
	}
	if args == nil || args[0] == "" {
		cmd = append(cmd, exec.Command(cmdName))
	} else {
		cmd = append(cmd, exec.Command(cmdName, args...))
	}

	cmd2 := exec.Command("tr", "a-z", "A-Z")
	stdin2, _ := cmd2.StdinPipe()
	//stdout, _ := cmd2.StdoutPipe()
	//a := "abcJe-"
	b, _ := cmd[0].Output()
	_, err := io.WriteString(stdin2, string(b))
	if err != nil {
		return "", err
	}
	err = stdin2.Close()

	b2, _ := cmd2.Output()
	_, err = io.WriteString(os.Stdout, string(b2)+" ")
	if err != nil {
		return "", err
	}
	return "", nil
}
