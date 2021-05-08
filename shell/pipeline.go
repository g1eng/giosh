package shell

import (
	"io"
	"log"
	"os/exec"
	"regexp"
)

type PipeIO struct {
	cmd        []*exec.Cmd
	stdin      []io.WriteCloser
	stdout     []io.ReadCloser
	expression [][]string
	result     [][]byte
	error      []error
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

//setExpression sets shell expression with IFS
//this function is applied to single lexicalScope
func (p *PipeIO) setExpression(lex string) {
	expr := regexp.MustCompilePOSIX(" |\\t").Split(lex, -1)
	expr = trimExpression(expr) //trim line-head space characters
	p.expression = append(p.expression, expr)
}

func (p *PipeIO) WriteTo(dest io.WriteCloser, output []byte) {
	_, err := io.WriteString(dest, string(output))
	p.error = append(p.error, err)
}

func (p *PipeIO) Flush() {
	p.cmd = []*exec.Cmd{}
	p.expression = [][]string{}
	p.stdin = []io.WriteCloser{}
	p.stdout = []io.ReadCloser{}
	p.error = []error{}
}
