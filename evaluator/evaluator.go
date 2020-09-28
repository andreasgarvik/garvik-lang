package evaluator

import (
	"log"
	"strconv"

	"github.com/andreasgarvik/inf122-lab3-go/parser"
	"github.com/golang-collections/collections/stack"
)

// An Evaluator will evaluate the tree
type Evaluator struct {
	*parser.BaseSimplVisitor
	stack       stack.Stack
	helperStack stack.Stack
}

// VisitMultExpr evaluates a multiplication expression
func (v *Evaluator) VisitMultExpr(ctx *parser.MultExprContext) interface{} {
	left := ctx.GetLeft().Accept(v).(int)
	right := ctx.GetRight().Accept(v).(int)
	return left * right
}

// VisitNumExpr evaluates a number expression
func (v *Evaluator) VisitNumExpr(ctx *parser.NumExprContext) interface{} {
	num, err := strconv.Atoi(ctx.NUM().GetText())
	if err != nil {
		log.Fatal("cannot convert NUM expr to int")
	}
	return num
}

// VisitParenExpr evaluates a parentheses expression
func (v *Evaluator) VisitParenExpr(ctx *parser.ParenExprContext) interface{} {
	return ctx.Expr().Accept(v)
}

// VisitLetExpr evaluates a let expression
func (v *Evaluator) VisitLetExpr(ctx *parser.LetExprContext) interface{} {
	env := make(map[string]int)
	id := ctx.ID().GetText()
	value := ctx.GetValue().Accept(v).(int)
	env[id] = value
	v.stack.Push(env)
	result := ctx.GetExpression().Accept(v).(int)
	v.stack.Pop()
	return result
}

// VisitAddExpr evaluates a add expression
func (v *Evaluator) VisitAddExpr(ctx *parser.AddExprContext) interface{} {
	left := ctx.GetLeft().Accept(v).(int)
	right := ctx.GetRight().Accept(v).(int)
	return left + right
}

// VisitIdExpr evaluates a ID expression
func (v *Evaluator) VisitIdExpr(ctx *parser.IdExprContext) interface{} {
	if v.stack.Len() != 0 {
		s := v.stack.Peek().(map[string]int)
		value, ok := s[ctx.ID().GetText()]
		if !ok {
			v.helperStack.Push(v.stack.Pop())
			value = ctx.Accept(v).(int)
			v.stack.Push(v.helperStack.Pop())
		}
		return value
	}
	return nil
}
