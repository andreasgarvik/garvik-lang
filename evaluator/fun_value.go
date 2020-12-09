package evaluator

import (
	"github.com/andreasgarvik/inf225-lab3-go/parser"
	"github.com/golang-collections/collections/stack"
)

// FunValue is storing an argument, a body and an environment
type FunValue struct {
	Params []string
	Body   parser.IExprContext
	Env    stack.Stack
}
