package main

import (
	"fmt"

	"github.com/andreasgarvik/inf225-lab3-go/evaluator"
	"github.com/andreasgarvik/inf225-lab3-go/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func eval(input string) int {
	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewSimplLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewSimplParser(stream)

	// Finally parse the expression (by walking the tree)
	var visitor evaluator.Evaluator
	result := p.Expr().Accept(&visitor)
	if result != nil {
		return result.(int)
	}
	return 0
}

func main() {
	fmt.Printf("The answer is: %d\n", eval("let f(x) = if x == 0 then 0 else if x == 1 then 1 else 2 in f(2) end"))
}
