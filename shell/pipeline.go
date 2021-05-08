package shell

import (
	gioParser "github.com/g1eng/giop/core"
	"io"
	"os"
	"os/exec"
)

type PipeIO struct {
	command []*exec.Cmd
	stdin   []io.WriteCloser
	stdout  []io.ReadCloser
	result  [][]byte
	error   []error
}

//PocPipe is a simple poc for line conversion with tr command.
//All lowercase characters are converted to uppercase within a pipeline
func (p *PipeIO) PocPipe(_ *gioParser.GioParser, s string) (string, error) {
	cmd := exec.Command("printf", s)
	cmd2 := exec.Command("tr", "a-z", "A-Z")
	stdin2, _ := cmd2.StdinPipe()
	//stdout, _ := cmd2.StdoutPipe()
	//a := "abcJe-"
	b, _ := cmd.Output()
	_, err := io.WriteString(stdin2, string(b))
	if err != nil {
		return "", err
	}
	err = stdin2.Close()

	b2, _ := cmd2.Output()
	_, err = io.WriteString(os.Stdout, string(b2))
	if err != nil {
		return "", err
	}
	return "", nil
}
