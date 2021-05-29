package shell

import "regexp"

//parseExpression sets shell expression which is separated with IFS.
//This function is applied to single lexicalScope
func (c *CommandLine) parseExpression(lex string) {
	expr := regexp.MustCompilePOSIX("[ \\t]").Split(lex, -1)
	expr = trimExpression(expr) //trim line-head space characters
	c.expression = append(c.expression, expr)
}

// trimExpression is the wrapper for trimExpressionTail and trimExpressionHead.
// It trims blank string "" from head and end of the given expression
func trimExpression(expr []string) []string {
	if len(expr) == 0 {
		return []string{""}
	}
	var newExpr []string
	for i := range expr {
		if !regexp.MustCompile("^[[:blank:]]*$").MatchString(expr[i]) {
			newExpr = append(newExpr, expr[i])
		}
	}
	return newExpr
}
