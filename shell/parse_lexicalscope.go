package shell

import "strings"

// parseLexicalScope sets all lexical scope on the line terminated with EOF.
// At now, giosh does not have multi-line scope for the lexical parsing.
//
// Each lexical scope within a line, is just separated with "|"
func (c *CommandLine) parseLexicalScope() {
	c.lexicalScope = strings.Split(c.GetInput(), "|")
}
