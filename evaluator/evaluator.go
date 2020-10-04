package evaluator

import (
	"log"
	"strconv"

	"github.com/andreasgarvik/inf225-lab3-go/parser"
	"github.com/golang-collections/collections/stack"
)

// An Evaluator will evaluate the tree
type Evaluator struct {
	*parser.BaseSimplVisitor
	stack stack.Stack
}

// VisitApplyExpr evaluates a apply expression
func (v *Evaluator) VisitApplyExpr(ctx *parser.ApplyExprContext) interface{} {
	id := ctx.GetId().Accept(v)
	if id != nil {
		fun := id.(FunValue)
		env := make(map[string]interface{})
		param := ctx.GetArg().Accept(v)
		env[fun.Arg] = param
		temp := v.stack
		v.stack = fun.Env
		v.stack.Push(env)
		result := fun.Body.Accept(v)
		v.stack.Pop()
		v.stack = temp
		return result
	}
	return nil
}

// VisitFunExpr evaluates a function expression
func (v *Evaluator) VisitFunExpr(ctx *parser.FunExprContext) interface{} {
	env := make(map[string]interface{})
	id := ctx.GetId().GetText()
	arg := ctx.GetArg().GetText()
	body := ctx.GetBody()
	env[id] = FunValue{arg, body, v.stack}
	v.stack.Push(env)
	result := ctx.GetCall().Accept(v)
	v.stack.Pop()
	return result

}

// VisitMultExpr evaluates a multiplication expression
func (v *Evaluator) VisitMultExpr(ctx *parser.MultExprContext) interface{} {
	left := ctx.GetLeft().Accept(v)
	right := ctx.GetRight().Accept(v)
	if left != nil && right != nil {
		l := left.(int)
		r := right.(int)
		return l * r
	}
	return nil
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
	env := make(map[string]interface{})
	id := ctx.GetId().GetText()
	value := ctx.GetValue().Accept(v)
	env[id] = value
	v.stack.Push(env)
	result := ctx.GetExpression().Accept(v)
	v.stack.Pop()
	return result
}

// VisitAddExpr evaluates a add expression
func (v *Evaluator) VisitAddExpr(ctx *parser.AddExprContext) interface{} {
	left := ctx.GetLeft().Accept(v)
	right := ctx.GetRight().Accept(v)
	if left != nil && right != nil {
		l := left.(int)
		r := right.(int)
		return l + r
	}
	return nil
}

// VisitIdExpr evaluates a ID expression
func (v *Evaluator) VisitIdExpr(ctx *parser.IdExprContext) interface{} {
	value := ctx.GetText()
	if value == "true" {
		return true
	}
	if value == "false" {
		return false
	}

	if v.stack.Len() != 0 {
		s := v.stack.Peek().(map[string]interface{})
		value, ok := s[ctx.ID().GetText()]
		if !ok {
			m := v.stack.Pop()
			value = ctx.Accept(v)
			v.stack.Push(m)
		}
		return value
	}
	return 0
}

// VisitIfElseExpr evaluates a if then else expression
func (v *Evaluator) VisitIfElseExpr(ctx *parser.IfElseExprContext) interface{} {
	con := ctx.GetCon().Accept(v).(bool)
	if con {
		return ctx.GetT().Accept(v)
	}
	return ctx.GetF().Accept(v)
}

// VisitDivExpr evaluates a division expression
func (v *Evaluator) VisitDivExpr(ctx *parser.DivExprContext) interface{} {
	left := ctx.GetLeft().Accept(v)
	right := ctx.GetRight().Accept(v)
	if left != nil && right != nil {
		l := left.(int)
		r := right.(int)
		return l / r
	}
	return nil
}

// VisitSubExpr evaluates a subtraction expression
func (v *Evaluator) VisitSubExpr(ctx *parser.SubExprContext) interface{} {
	left := ctx.GetLeft().Accept(v)
	right := ctx.GetRight().Accept(v)
	if left != nil && right != nil {
		l := left.(int)
		r := right.(int)
		return l - r
	}
	return nil
}

// VisitEqualExpr evaluates a equal expression
func (v *Evaluator) VisitEqualExpr(ctx *parser.EqualExprContext) interface{} {
	left := ctx.GetLeft().Accept(v)
	right := ctx.GetRight().Accept(v)
	if left != nil && right != nil {
		l := left.(int)
		r := right.(int)
		return l == r
	}
	return nil
}
