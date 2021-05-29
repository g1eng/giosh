package shell

import (
	. "gopkg.in/check.v1"
	"os/exec"
)

func (s *CommandLine) TestCommandLine_ParseExpression(c *C) {
	sh := New()
	if len(sh.expression) != 0 {
		c.Fail()
	}
	sh.lexicalScope = append(sh.lexicalScope, "ls -l -n")
	sh.parseExpression(sh.lexicalScope[0])
	if len(sh.expression) != 1 {
		c.Fail()
	} else if sh.expression[0][0] != "ls" {
		c.Fail()
	} else if len(sh.expression[0]) != 3 {
		c.Fatal("expression length must be 3 for `ls -l -n`")
	} else if sh.expression[0][1] != "-l" {
		c.Fail()
	} else if sh.expression[0][2] != "-n" {
		c.Fail()
	}
}

func (s *CommandLine) TestTrimExpressionWithBlankSlice(c *C) {
	out := trimExpression([]string{})
	c.Check(len(out), Equals, 1)
}

func (s *CommandLine) TestCommandLine_TrimExpressionWithBlankPadding(c *C) {
	ex := []string{"\t\t", "  ", "ls", "  ", "-l"}
	out := trimExpression(ex)
	c.Check(len(out), Equals, 2)
}

func (s *CommandLine) TestGetCurrentStdout(c *C) {
	c1 := exec.Command("ls", "-l")
	c2 := exec.Command("grep", "ok")
	c1Stdout, _ := c1.StdoutPipe()
	c2Stdin, _ := c2.StdinPipe()
	pipe := Pipe{
		stdout: c1Stdout,
		stdin:  c2Stdin,
	}
	s.pipeSet = append(s.pipeSet, pipe)
	c.Check(s.getCurrentStdout(), Equals, c1Stdout)
}

func (s *CommandLine) TestGetCurrentStdoutWithoutPipeset(c *C) {
	s.pipeSet = []Pipe{}
	errLen := len(s.error)
	c.Check(s.getCurrentStdout(), IsNil)
	c.Check(len(s.error), Equals, errLen+1)
}

func (s *CommandLine) TestIsPipeEnd(c *C) {
	s.lexicalScope = []string{"ls -l", "grep m"}
	s.exprIndex = 1
	c.Check(s.isPipeEnd(), Equals, true)
}

func (s *CommandLine) TestIsNotPipeEnd(c *C) {
	s.lexicalScope = []string{"ls -l", "grep m", "tr m A"}
	s.exprIndex = 1
	c.Check(s.isPipeEnd(), Equals, false)
}

func (s *CommandLine) TestIsBlankLine(c *C) {
	s.lexicalScope = []string{}
	s.expression = [][]string{}
	c.Check(s.isBlankLine(), Equals, true)
}

func (s *CommandLine) TestIsBlankLine2(c *C) {
	s.lexicalScope = []string{"ls", "-l"}
	s.expression = [][]string{}
	c.Check(s.isBlankLine(), Equals, true)
}

func (s *CommandLine) TestIsBlankLine3(c *C) {
	s.lexicalScope = []string{"ls", "-l"}
	s.expression = [][]string{}
	s.expression = append(s.expression, []string{})
	c.Check(s.isBlankLine(), Equals, true)
}

func (s *CommandLine) TestNonBlankLine(c *C) {
	s.lexicalScope = []string{"ls", "-l"}
	s.expression = [][]string{}
	s.expression = append(s.expression, []string{"ls"})
	c.Check(s.isBlankLine(), Equals, false)
}
