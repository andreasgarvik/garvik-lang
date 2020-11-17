package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/andreasgarvik/inf225-lab3-go/evaluator"
	"github.com/andreasgarvik/inf225-lab3-go/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func eval(input string) interface{} {
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
	if result == nil {
		log.Fatal("Error validating input")
	}
	return result
}

func main() {
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Println(eval(string(content)))
}
