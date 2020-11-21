package evaluator

import (
	"fmt"
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

// VisitProgram evaluates a program
func (v *Evaluator) VisitProgram(ctx *parser.ProgramContext) interface{} {
	fmt.Println("program")
	for _, expr := range ctx.AllExpr() {
		result := expr.Accept(v)
		if result != nil {
			fmt.Println(result)
		}
	}
	return nil
}

// VisitCallExpr evaluates a apply expression
func (v *Evaluator) VisitCallExpr(ctx *parser.CallExprContext) interface{} {
	fmt.Println("call")
	fun := ctx.GetFun().Accept(v)
	if fun != nil {
		f := fun.(FunValue)
		env := make(map[string]interface{})
		p := ctx.GetArg().Accept(v)
		env[ctx.GetFun().GetText()] = f
		env[f.Arg.(string)] = p
		temp := v.stack
		v.stack = f.Env
		v.stack.Push(env)
		result := f.Body.Accept(v)
		v.stack.Pop()
		v.stack = temp
		return result
	}
	return nil
}

// VisitMultExpr evaluates a multiplication expression
func (v *Evaluator) VisitMultExpr(ctx *parser.MultExprContext) interface{} {
	fmt.Println("mult")
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
	fmt.Println("num")
	num, err := strconv.Atoi(ctx.NUM().GetText())
	if err != nil {
		log.Fatal("Parsing error")
	}
	return num
}

// VisitStrExpr evaluates a string expression
func (v *Evaluator) VisitStrExpr(ctx *parser.StrExprContext) interface{} {
	fmt.Println("str")
	str := ctx.STR().GetText()
	r := []rune(str)
	return string(r[1 : len(str)-1])
}

// VisitParenExpr evaluates a parentheses expression
func (v *Evaluator) VisitParenExpr(ctx *parser.ParenExprContext) interface{} {
	fmt.Println("paren")
	return ctx.Expr().Accept(v)
}

// VisitAddExpr evaluates a add expression
func (v *Evaluator) VisitAddExpr(ctx *parser.AddExprContext) interface{} {
	fmt.Println("add")
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
	fmt.Println("id")
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
	return nil
}

// VisitBoolExpr evaluates a boolean expression
func (v *Evaluator) VisitBoolExpr(ctx *parser.BoolExprContext) interface{} {
	fmt.Println("bool")
	b := ctx.GetText()
	if b == "true" {
		return true
	}
	if b == "false" {
		return false
	}
	return nil
}

// VisitIfElseExpr evaluates an if then else expression
func (v *Evaluator) VisitIfElseExpr(ctx *parser.IfElseExprContext) interface{} {
	fmt.Println("ifelse")
	con := ctx.GetCon().Accept(v)
	if con.(bool) {
		return ctx.GetT().Accept(v)
	}
	return ctx.GetF().Accept(v)
}

// VisitDivExpr evaluates a division expression
func (v *Evaluator) VisitDivExpr(ctx *parser.DivExprContext) interface{} {
	fmt.Println("div")
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
	fmt.Println("sub")
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
	fmt.Println("equal")
	left := ctx.GetLeft().Accept(v)
	right := ctx.GetRight().Accept(v)
	if left != nil && right != nil {
		l := left.(int)
		r := right.(int)
		return l == r
	}
	return nil
}

// VisitLambdaExpr evaluates a lambda expression
func (v *Evaluator) VisitLambdaExpr(ctx *parser.LambdaExprContext) interface{} {
	fmt.Println("lamdba")
	arg := ctx.GetArg().GetText()
	body := ctx.GetBody()
	return FunValue{arg, body, v.stack}
}

// VisitVarExpr evaluates a var expression
func (v *Evaluator) VisitVarExpr(ctx *parser.VarExprContext) interface{} {
	fmt.Println("var")
	env := make(map[string]interface{})
	id := ctx.GetId().GetText()
	value := ctx.GetValue().Accept(v)
	env[id] = value
	v.stack.Push(env)
	return nil
}

// VisitPrintExpr evaluates a print expression
func (v *Evaluator) VisitPrintExpr(ctx *parser.PrintExprContext) interface{} {
	fmt.Println("print")
	return ctx.Expr().Accept(v)
}
