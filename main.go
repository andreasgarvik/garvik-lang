package main

import (
	"fmt"

	"github.com/andreasgarvik/inf122-lab3-go/evaluator"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/andreasgarvik/inf122-lab3-go/parser"
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
	return p.Expr().Accept(&visitor).(int)
}

func main() {
	fmt.Printf("The answer is: %d\n", eval("1 + 2 * 3"))
}
