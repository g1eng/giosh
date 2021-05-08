package shell

import (
	gioParser "github.com/g1eng/giop/core"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var lineNo = 1

// GetPsString prints PS shell description
func GetPsString() string {
	return os.Getenv("USER") + "@G[" + strconv.Itoa(lineNo) + "]> "
}

// Exec is a ParserFunc, which returns the result string of the command execution
func (p *PipeIO) Exec(_ *gioParser.GioParser, s string) (string, error) {
	var (
		cmdName      string
		args         []string
		lexicalScope []string
		cmd          []*exec.Cmd
		cmd1         *exec.Cmd
		cmd2         *exec.Cmd
	)
	p.Flush()
	lexicalScope = strings.Split(s, "|")
	if len(lexicalScope) != 0 {
		p.setExpression(lexicalScope[0])
		cmdName = p.expression[0][0]
	} else {
		return GetPsString(), nil
	}
	if len(p.expression[0][0]) > 1 {
		args = p.expression[0][1:]
		args = trimExpression(args) //trim line-head space characters
		for i := range p.expression[0] {
			log.Printf("expression[0][%d]: %s", i, p.expression[0][i])
		}
		for i := range args {
			log.Printf("args[%d]: %s", i, args[i])
		}
	}
	if args == nil || args[0] == "" {
		cmd1 = exec.Command(cmdName)
		p.cmd = append(cmd, cmd1)
	} else {
		cmd1 = exec.Command(cmdName, args...)
		p.cmd = append(cmd, cmd1)
	}
	p.stdin = append(p.stdin, nil)
	p.stdout = append(p.stdout, nil)

	originOutput, err := cmd1.Output()
	p.error = append(p.error, err)

	if len(lexicalScope) == 1 {
		_, err = io.WriteString(os.Stdout, string(originOutput))
		p.error = append(p.error, err)
		return "", err
	}

	for i := range lexicalScope {
		log.Printf("lexicalScope[%d]: %v", i, lexicalScope[i])
		if i != 0 {
			p.setExpression(lexicalScope[i])
			tmpIndex := len(p.expression) - 1
			cmdName = p.expression[tmpIndex][0]
			args = p.expression[tmpIndex][1:]
			args = trimExpression(args) //trim line-head space characters
			log.Printf("expression[%d]: %v", i, p.expression[i])
			for j := range p.expression[i] {
				log.Printf("expression[%d][%d]: %v", i, j, p.expression[i][j])
			}

			cmd2 = exec.Command("tr", "a-z", "A-Z")
			p.cmd = append(p.cmd, cmd2)
			stdin2, _ := cmd2.StdinPipe()
			p.stdin = append(p.stdin, stdin2)
			_, err = io.WriteString(stdin2, string(originOutput))
			p.error = append(p.error, err)
			if err != nil {
				return "", err
			}
			err = stdin2.Close()

			processOutput, _ := cmd2.Output()
			_, err = io.WriteString(os.Stdout, string(processOutput)+" ")
			p.error = append(p.error, err)
			if err != nil {
				return "", err
			}
		}
	}
	return "", nil
}
