package shell

import "log"

func (c *CommandLine) dumpParserObject(seq int) {
	if len(c.lexicalScope)-1 < seq {
		log.Fatalf("invalid debug target %d\n", seq)
	}
	for j := range c.expression[seq] {
		log.Printf("lexicalScope[%d]: %v", seq, c.lexicalScope[seq])
		log.Printf("expression[%d]: %v", seq, c.expression[seq])
		log.Printf("expression[%d][%d]: %v", seq, j, c.expression[seq][j])
		log.Printf("cmd path %d: %s", seq, c.command[seq].Path)
		log.Printf("args %d: %v", seq, c.command[seq].Args)
	}
}
