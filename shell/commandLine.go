package shell

import (
	"fmt"
	gioParser "github.com/g1eng/giop/core"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

var line = 1

// validate validates the syntax of given expression.
// It aims to be a syntax checker for giosh.
// FIXME [this is stub]
func validate(cmdLine string) error {
	return nil
}

// PipelineExec is Exec alternative, which parses pipe (|) in a command line
func PipelineExec(g *gioParser.GioParser, cmdLine string) (string, error) {
	var (
		pipe PipeIO
		stdout io.ReadCloser
		stdin io.WriteCloser
		err error
	)

	//set "|" as the expression splitter.
	//any lexical scopes separated with "|".
	regex := regexp.MustCompilePOSIX("\\|")
	lexicalScopes := regex.Split(cmdLine, -1)
	fmt.Println(lexicalScopes)
	for i := range lexicalScopes {
		g.Lex.Delimiter = " "
		//set " " as the word splitter IFS
		//command and arguments are separated with " " each other
		regex = regexp.MustCompilePOSIX(" ")
		expression := regex.Split(lexicalScopes[i], -1)
		for j := range expression {
			if len(expression) > 1 && expression[j] == "" {
				expression = expression[1:]
			}
		}
		log.Printf("i=%d lexicalScopes[%d]=%s", i, i, lexicalScopes[i])
		log.Printf("     expression %d=%s", i, expression)
		log.Printf("pipe=")
		log.Println(pipe)
		log.Printf("pipe.command")
		fmt.Println(pipe.command)
		log.Printf("pipe.stdin")
		fmt.Println(pipe.stdin)
		log.Printf("pipe.stdout")
		log.Println(pipe.stdout)
		log.Printf("pipe.result")
		log.Println(pipe.result)
		cmdName, args := setExpression(expression)


		cmd := exec.Command(cmdName, args...)
		pipe.command = append(pipe.command, cmd)
		stdin, err = cmd.StdinPipe()
		pipe.stdin = append(pipe.stdin, stdin)
		pipe.error = append(pipe.error, err)

		// if there are at least one pipe,
		// then write stdout into stdin pipe of the new expression

		go func(){
			if i > 0 {
				if _, err = io.WriteString(pipe.stdin[i], string(pipe.result[i-1])); err != nil{
					log.Fatal(err)
				}
			}
			var result []byte
			result, err = cmd.Output()
			pipe.result = append(pipe.result, result)
			pipe.error = append(pipe.error, err)
			stdout, err = cmd.StdoutPipe()
			pipe.stdout = append(pipe.stdout, stdout)
			pipe.error = append(pipe.error, err)
		}()
	}
	//displaying error
	for i := range pipe.error {
		if pipe.error[i] != nil {
			err = pipe.error[i]
			log.Println(err)
		}
	}
	//calculate result
	resultArray := pipe.result
	resultLength := len(resultArray)
	println("result length ", resultLength)
	if resultLength == 0 {
		return "", err
	} else {
		return string(resultArray[resultLength-1]), err
	}
}

func PrintPsString() error {
	_, err := fmt.Printf( getPsString() )
	return err
}
func getPsString() string {
	return os.Getenv("USER") + "@G[" + strconv.Itoa(line) + "]> "
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
		return getPsString(), nil
	} else {
		// for command execution
		b, err := exec.Command(cmdName, args...).Output()
		line = line + 1
		if err != nil {
			fmt.Printf("giosh %s\n", err)
		}
		return string(b) + getPsString(), nil
	}
}
