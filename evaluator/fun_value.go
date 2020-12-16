package evaluator

import (
	"github.com/andreasgarvik/garvik-lang/parser"
	"github.com/golang-collections/collections/stack"
)

// FunValue is storing an argument, a body and an environment
type FunValue struct {
	Param interface{}
	Body  parser.IExprContext
	Env   stack.Stack
}
