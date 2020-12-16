package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/andreasgarvik/garvik-lang/evaluator"
	"github.com/andreasgarvik/garvik-lang/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func eval(input string) {
	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewGarvikLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewGarvikParser(stream)

	// Finally parse the expression (by walking the tree)
	var visitor evaluator.Evaluator
	p.Program().Accept(&visitor)
}

func main() {
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	eval(string(content))
}
