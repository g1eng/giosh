package shell

import (
	gioParser "github.com/g1eng/giop/core"
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

func (p *PipeIO) execCommand(cmdName string, args []string) {
	var cmd *exec.Cmd
	if args == nil || args[0] == "" {
		cmd = exec.Command(cmdName)
		p.cmd = append(p.cmd, cmd)
	} else {
		cmd = exec.Command(cmdName, args...)
		p.cmd = append(p.cmd, cmd)
	}
	lineNo++
	p.stdin = append(p.stdin, nil)
	p.stdout = append(p.stdout, nil)
}

// Exec is a ParserFunc, which returns the result string of the command execution
func (p *PipeIO) Exec(_ *gioParser.GioParser, s string) (string, error) {
	var (
		cmdName      string
		args         []string
		lexicalScope []string
	)
	p.Flush()
	lexicalScope = strings.Split(s, "|")
	if len(lexicalScope) != 0 {
		p.setExpression(lexicalScope[0])
	}
	if len(p.expression[0]) == 0 {
		return GetPsString(), nil
	} else if len(p.expression[0][0]) == 1 {
		cmdName = p.expression[0][0]
	} else {
		cmdName = p.expression[0][0]
		args = p.expression[0][1:]
		args = trimExpression(args) //trim line-head space characters
		for i := range p.expression[0] {
			log.Printf("expression[0][%d]: %s", i, p.expression[0][i])
		}
		for i := range args {
			log.Printf("args[%d]: %s", i, args[i])
		}
	}

	p.execCommand(cmdName, args)

	originOutput, err := p.cmd[0].Output()
	p.error = append(p.error, err)

	if len(lexicalScope) == 0 {
		return GetPsString(), nil
	} else if len(lexicalScope) == 1 {
		err = p.WriteTo(os.Stdout, originOutput)
		p.error = append(p.error, err)
		return GetPsString(), err
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

			p.execCommand(cmdName, args)
			cmd2 := exec.Command(cmdName, args...)
			p.cmd = append(p.cmd, cmd2)

			stdin2, _ := cmd2.StdinPipe()
			p.stdin = append(p.stdin, stdin2)
			err = p.WriteTo(stdin2, originOutput)
			p.error = append(p.error, err)
			err = stdin2.Close()

			processOutput, _ := cmd2.Output()
			err = p.WriteTo(os.Stdout, processOutput)
			p.error = append(p.error, err)
		}
	}
	for i := range p.error {
		if p.error[i] != nil {
			return GetPsString(), err
		}
	}
	_ = p.WriteTo(os.Stdout, []byte(GetPsString()))
	return GetPsString(), nil
}
