package main

import (
	"fmt"

	"rule-engine/parser"
	"rule-engine/visitor"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	is := antlr.NewInputStream("a.b > -1 and 1 > -1")
	//fmt.Println(calc("1+2+(3+4)*(5-10)"))
	//
	// Create the Lexer
	lexer := parser.NewRuleLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewRuleParser(stream)
	tree := p.Logical()

	v := visitor.RuleVisitor{
		Items: `{"a":{"b":10}}`,
	}
	res := v.Visit(tree)
	float, ok := res.(bool)
	fmt.Println(float, ok)
}
