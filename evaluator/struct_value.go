package evaluator

import "github.com/golang-collections/collections/stack"

// StructValue is storing fields and an environment
type StructValue struct {
	Env *stack.Stack
}
