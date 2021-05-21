package shell

import "regexp"

//setExpression sets shell expression which is separated with IFS.
//This function is applied to single lexicalScope
func (c *CommandLine) setExpression(lex string) {
	expr := regexp.MustCompilePOSIX("[ \\t]").Split(lex, -1)
	expr = trimExpression(expr) //trim line-head space characters
	c.expression = append(c.expression, expr)
}

// trimExpressionHead trims blank string from the head of the expression array
//It returns processed expression
func trimExpressionHead(expr []string) []string {
	if expr[0] == "" {
		expr = expr[1:]
		expr = trimExpression(expr)
	}
	return expr
}

//trimExpressionTail trims blank string from the end of the expression array.
//It returns processed expression
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

// trimExpression is the wrapper for trimExpressionTail and trimExpressionHead.
// It trims blank string "" from head and end of the given expression
func trimExpression(expr []string) []string {
	if len(expr) == 0 {
		return []string{""}
	}
	expr = trimExpressionHead(expr)
	expr = trimExpressionTail(expr)
	return expr
}
