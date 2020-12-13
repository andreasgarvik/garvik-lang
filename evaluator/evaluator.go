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
	*parser.BaseGarvikVisitor
	stack stack.Stack
}

// VisitProgram evaluates a program
func (v *Evaluator) VisitProgram(ctx *parser.ProgramContext) interface{} {
	for _, expr := range ctx.AllExpr() {
		result := expr.Accept(v)
		if result != nil {
			fun, ok := result.(FunValue)
			if ok {
				fmt.Printf("%s -> %s \n", fun.Param, fun.Body.GetText())
			} else {
				fmt.Println(result)
			}
		}
	}
	return nil
}

// VisitCallExpr evaluates a apply expression
func (v *Evaluator) VisitCallExpr(ctx *parser.CallExprContext) interface{} {
	fun := ctx.GetFun().Accept(v)
	if fun != nil {
		f, ok := fun.(FunValue)
		if ok {
			env := make(map[interface{}]interface{})
			env[ctx.GetFun().GetText()] = f
			env[f.Param] = ctx.GetArg().Accept(v)
			temp := v.stack
			v.stack = f.Env
			v.stack.Push(env)
			result := f.Body.Accept(v)
			v.stack.Pop()
			v.stack = temp
			return result
		}
		return fmt.Sprintf("%s is not a function", ctx.GetFun().GetText())
	}
	return fmt.Sprintf("%s is not defined", ctx.GetFun().GetText())
}

// VisitLetExpr evaluates a let expression
func (v *Evaluator) VisitLetExpr(ctx *parser.LetExprContext) interface{} {
	env := make(map[interface{}]interface{})
	id := ctx.GetId().GetText()
	value := ctx.GetValue().Accept(v)
	env[id] = value
	v.stack.Push(env)
	result := ctx.GetExpression().Accept(v)
	v.stack.Pop()
	return result
}

// VisitMultExpr evaluates a multiplication expression
func (v *Evaluator) VisitMultExpr(ctx *parser.MultExprContext) interface{} {
	left := ctx.GetLeft().Accept(v)
	right := ctx.GetRight().Accept(v)
	if left != nil && right != nil {
		l, ok := left.(int)
		if ok {
			r, ok := right.(int)
			if ok {
				return l * r
			}
			return fmt.Sprintf("%s is not a number", ctx.GetRight().GetText())
		}
		return fmt.Sprintf("%s is not a number", ctx.GetLeft().GetText())
	}
	if left == nil && right != nil {
		return fmt.Sprintf("%s is not defined", ctx.GetLeft().GetText())
	}

	if left != nil && right == nil {
		return fmt.Sprintf("%s is not defined", ctx.GetRight().GetText())
	}
	return fmt.Sprintf("%s and %s is not defined", ctx.GetLeft().GetText(), ctx.GetRight().GetText())
}

// VisitNumExpr evaluates a number expression
func (v *Evaluator) VisitNumExpr(ctx *parser.NumExprContext) interface{} {
	num, err := strconv.Atoi(ctx.NUM().GetText())
	if err != nil {
		log.Fatal("Parsing error")
	}
	return num
}

// VisitStrExpr evaluates a string expression
func (v *Evaluator) VisitStrExpr(ctx *parser.StrExprContext) interface{} {
	str := ctx.STR().GetText()
	r := []rune(str)
	return string(r[1 : len(str)-1])
}

// VisitParenExpr evaluates a parentheses expression
func (v *Evaluator) VisitParenExpr(ctx *parser.ParenExprContext) interface{} {
	return ctx.Expr().Accept(v)
}

// VisitAddExpr evaluates a add expression
func (v *Evaluator) VisitAddExpr(ctx *parser.AddExprContext) interface{} {
	left := ctx.GetLeft().Accept(v)
	right := ctx.GetRight().Accept(v)
	if left != nil && right != nil {
		l, ok := left.(int)
		if ok {
			r, ok := right.(int)
			if ok {
				return l + r
			}
			return fmt.Sprintf("%s is not a number", ctx.GetRight().GetText())
		}
		return fmt.Sprintf("%s is not a number", ctx.GetLeft().GetText())
	}
	if left == nil && right != nil {
		return fmt.Sprintf("%s is not defined", ctx.GetLeft().GetText())
	}

	if left != nil && right == nil {
		return fmt.Sprintf("%s is not defined", ctx.GetRight().GetText())
	}
	return fmt.Sprintf("%s and %s is not defined", ctx.GetLeft().GetText(), ctx.GetRight().GetText())
}

// VisitIdExpr evaluates a ID expression
func (v *Evaluator) VisitIdExpr(ctx *parser.IdExprContext) interface{} {
	if v.stack.Len() != 0 {
		s := v.stack.Peek().(map[interface{}]interface{})
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
	con := ctx.GetCon().Accept(v)
	c, ok := con.(bool)
	if ok {
		if c {
			return ctx.GetT().Accept(v)
		}
		return ctx.GetF().Accept(v)
	}
	return fmt.Sprintf("%s is not a condition", ctx.GetCon().GetText())
}

// VisitDivExpr evaluates a division expression
func (v *Evaluator) VisitDivExpr(ctx *parser.DivExprContext) interface{} {
	left := ctx.GetLeft().Accept(v)
	right := ctx.GetRight().Accept(v)
	if left != nil && right != nil {
		l, ok := left.(int)
		if ok {
			r, ok := right.(int)
			if ok {
				return l / r
			}
			return fmt.Sprintf("%s is not a number", ctx.GetRight().GetText())
		}
		return fmt.Sprintf("%s is not a number", ctx.GetLeft().GetText())
	}
	if left == nil && right != nil {
		return fmt.Sprintf("%s is not defined", ctx.GetLeft().GetText())
	}

	if left != nil && right == nil {
		return fmt.Sprintf("%s is not defined", ctx.GetRight().GetText())
	}
	return fmt.Sprintf("%s and %s is not defined", ctx.GetLeft().GetText(), ctx.GetRight().GetText())
}

// VisitSubExpr evaluates a subtraction expression
func (v *Evaluator) VisitSubExpr(ctx *parser.SubExprContext) interface{} {
	left := ctx.GetLeft().Accept(v)
	right := ctx.GetRight().Accept(v)
	if left != nil && right != nil {
		l, ok := left.(int)
		if ok {
			r, ok := right.(int)
			if ok {
				return l - r
			}
			return fmt.Sprintf("%s is not a number", ctx.GetRight().GetText())
		}
		return fmt.Sprintf("%s is not a number", ctx.GetLeft().GetText())
	}
	if left == nil && right != nil {
		return fmt.Sprintf("%s is not defined", ctx.GetLeft().GetText())
	}

	if left != nil && right == nil {
		return fmt.Sprintf("%s is not defined", ctx.GetRight().GetText())
	}
	return fmt.Sprintf("%s and %s is not defined", ctx.GetLeft().GetText(), ctx.GetRight().GetText())
}

// VisitEqualExpr evaluates a equal expression
func (v *Evaluator) VisitEqualExpr(ctx *parser.EqualExprContext) interface{} {
	left := ctx.GetLeft().Accept(v)
	right := ctx.GetRight().Accept(v)
	if left != nil && right != nil {
		l, ok := left.(int)
		if ok {
			r, ok := right.(int)
			if ok {
				return l == r
			}
			return fmt.Sprintf("%s is not a number", ctx.GetRight().GetText())
		}
		return fmt.Sprintf("%s is not a number", ctx.GetLeft().GetText())
	}
	if left == nil && right != nil {
		return fmt.Sprintf("%s is not defined", ctx.GetLeft().GetText())
	}

	if left != nil && right == nil {
		return fmt.Sprintf("%s is not defined", ctx.GetRight().GetText())
	}
	return fmt.Sprintf("%s and %s is not defined", ctx.GetLeft().GetText(), ctx.GetRight().GetText())
}

// VisitLessExpr evaluates a less then expression
func (v *Evaluator) VisitLessExpr(ctx *parser.LessExprContext) interface{} {
	left := ctx.GetLeft().Accept(v)
	right := ctx.GetRight().Accept(v)
	if left != nil && right != nil {
		l, ok := left.(int)
		if ok {
			r, ok := right.(int)
			if ok {
				return l < r
			}
			return fmt.Sprintf("%s is not a number", ctx.GetRight().GetText())
		}
		return fmt.Sprintf("%s is not a number", ctx.GetLeft().GetText())
	}
	if left == nil && right != nil {
		return fmt.Sprintf("%s is not defined", ctx.GetLeft().GetText())
	}

	if left != nil && right == nil {
		return fmt.Sprintf("%s is not defined", ctx.GetRight().GetText())
	}
	return fmt.Sprintf("%s and %s is not defined", ctx.GetLeft().GetText(), ctx.GetRight().GetText())
}

// VisitGreaterExpr evaluates a greater then expression
func (v *Evaluator) VisitGreaterExpr(ctx *parser.GreaterExprContext) interface{} {
	left := ctx.GetLeft().Accept(v)
	right := ctx.GetRight().Accept(v)
	if left != nil && right != nil {
		l, ok := left.(int)
		if ok {
			r, ok := right.(int)
			if ok {
				return l > r
			}
			return fmt.Sprintf("%s is not a number", ctx.GetRight().GetText())
		}
		return fmt.Sprintf("%s is not a number", ctx.GetLeft().GetText())
	}
	if left == nil && right != nil {
		return fmt.Sprintf("%s is not defined", ctx.GetLeft().GetText())
	}

	if left != nil && right == nil {
		return fmt.Sprintf("%s is not defined", ctx.GetRight().GetText())
	}
	return fmt.Sprintf("%s and %s is not defined", ctx.GetLeft().GetText(), ctx.GetRight().GetText())
}

// VisitLambdaExpr evaluates a lambda expression
func (v *Evaluator) VisitLambdaExpr(ctx *parser.LambdaExprContext) interface{} {
	param := ctx.GetParam().GetText()
	body := ctx.GetBody()
	return FunValue{param, body, v.stack}
}

// VisitAssignExpr evaluates a var expression
func (v *Evaluator) VisitAssignExpr(ctx *parser.AssignExprContext) interface{} {
	env := make(map[interface{}]interface{})
	id := ctx.GetId().GetText()
	value := ctx.GetValue().Accept(v)
	if value != nil {
		env[id] = value
		v.stack.Push(env)
		return nil
	}
	return fmt.Sprintf("%s is not defined", ctx.GetValue().GetText())
}

// VisitCommentExpr skips a comment
func (v *Evaluator) VisitCommentExpr(ctx *parser.CommentExprContext) interface{} {
	return nil
}

// VisitStructExpr evaluates a struct expression
func (v *Evaluator) VisitStructExpr(ctx *parser.StructExprContext) interface{} {
	return ctx.AllExpr()
}

// VisitDotExpr evaluates a dot expression
func (v *Evaluator) VisitDotExpr(ctx *parser.DotExprContext) interface{} {
	id := ctx.GetId().Accept(v)
	if id != nil {
		s, ok := id.([]parser.IExprContext)
		if ok {
			for _, expr := range s {
				expr.Accept(v)
			}
			f := ctx.GetField().Accept(v)
			if f != nil {
				return f
			}
			return fmt.Sprintf("%s is not a field of %s", ctx.GetField().GetText(), ctx.GetId().GetText())
		}
		return fmt.Sprintf("%s is not a struct", ctx.GetId().GetText())
	}
	return fmt.Sprintf("%s is not defined", ctx.GetId().GetText())
}

// VisitListExpr evaluates a list expression
func (v *Evaluator) VisitListExpr(ctx *parser.ListExprContext) interface{} {
	var list []interface{}
	for _, expr := range ctx.AllExpr() {
		list = append(list, expr.Accept(v))
	}
	return list
}

// VisitLookupExpr evaluates a lookup expression
func (v *Evaluator) VisitLookupExpr(ctx *parser.LookupExprContext) interface{} {
	id := ctx.GetId().Accept(v)
	if id != nil {
		list, ok := id.([]interface{})
		if ok {
			key, ok := ctx.GetKey().Accept(v).(int)
			if ok {
				if key < len(list) {
					return list[key]
				}
				return fmt.Sprint("list index out of bounds")
			}
			return fmt.Sprintf("%s is not a number", ctx.GetKey().GetText())
		}
		return fmt.Sprintf("%s is not a list", ctx.GetId().GetText())
	}
	return fmt.Sprintf("%s is not defined", ctx.GetId().GetText())
}

// VisitLenExpr evaluates a lenght of list expression
func (v *Evaluator) VisitLenExpr(ctx *parser.LenExprContext) interface{} {
	id := ctx.GetId().Accept(v)
	if id != nil {
		list, ok := id.([]interface{})
		if ok {
			return len(list)
		}
		return fmt.Sprintf("%s is not a list", ctx.GetId().GetText())
	}
	return fmt.Sprintf("%s is not defined", ctx.GetId().GetText())
}

// VisitLookupAssignExpr evaluates a lookup assignment expression
func (v *Evaluator) VisitLookupAssignExpr(ctx *parser.LookupAssignExprContext) interface{} {
	id := ctx.GetId().Accept(v)
	if id != nil {
		list, ok := id.([]interface{})
		if ok {
			key, ok := ctx.GetKey().Accept(v).(int)
			if ok {
				value := ctx.GetValue().Accept(v)
				if value != nil {
					if key < len(list) {
						list[key] = value
						return list
					}
					return append(list, value)
				}
				return fmt.Sprintf("%s is not defined", ctx.GetValue().GetText())
			}
			return fmt.Sprintf("%s is not a number", ctx.GetKey().GetText())
		}
		return fmt.Sprintf("%s is not a list", ctx.GetId().GetText())
	}
	return fmt.Sprintf("%s is not defined", ctx.GetId().GetText())
}
