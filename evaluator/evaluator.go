package evaluator

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/andreasgarvik/garvik-lang/parser"
	"github.com/golang-collections/collections/stack"
)

// An Evaluator will evaluate the tree
type Evaluator struct {
	*parser.BaseGarvikVisitor
	stack stack.Stack
	heap  map[interface{}]interface{}
}

// VisitProgram evaluates a program
func (v *Evaluator) VisitProgram(ctx *parser.ProgramContext) interface{} {
	v.heap = make(map[interface{}]interface{})
	for _, expr := range ctx.AllExpr() {
		result := expr.Accept(v)
		if result != nil {
			switch t := result.(type) {
			case FunValue:
				fmt.Printf("%s -> %s \n", t.Param, t.Body.GetText())
			case MethodValue:
				fmt.Printf("%s -> %s \n", t.Fun.Param, t.Fun.Body.GetText())
			case StructValue:
				fmt.Print("{ ")
				for id, field := range t.Env.Peek().(map[interface{}]interface{}) {
					if fun, ok := field.(parser.IExprContext); ok {
						f := fun.Accept(v)
						fun := f.(FunValue)
						fmt.Printf("%s: %s -> %s ", id, fun.Param, fun.Body.GetText())
					} else {
						fmt.Printf("%s: %v ", id, field)
					}
				}
				fmt.Printf("} \n")
			case []interface{}:
				fmt.Print("[")
				for i, e := range t {
					switch lt := e.(type) {
					case FunValue:
						fmt.Printf("%s -> %s,", lt.Param, lt.Body.GetText())
					case StructValue:
						fmt.Print("{ ")
						for id, field := range lt.Env.Peek().(map[interface{}]interface{}) {
							if fun, ok := field.(parser.IExprContext); ok {
								f := fun.Accept(v)
								fun := f.(FunValue)
								fmt.Printf("%s: %s -> %s ", id, fun.Param, fun.Body.GetText())
							} else {
								fmt.Printf("%s: %v ", id, field)
							}
						}
						fmt.Printf("},")
					default:
						if i == len(t)-1 {
							fmt.Printf("%v", e)
						} else {
							fmt.Printf("%v,", e)
						}
					}
				}
				fmt.Println("]")
			default:
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
		switch t := fun.(type) {
		case FunValue:
			env := make(map[interface{}]interface{})
			env[ctx.GetFun().GetText()] = t
			arg := ctx.GetArg().Accept(v)
			env[t.Param] = arg
			temp := v.stack
			v.stack = t.Env
			v.stack.Push(env)
			result := t.Body.Accept(v)
			v.stack.Pop()
			v.stack = temp
			return result

		case MethodValue:
			arg := ctx.GetArg().Accept(v)
			e := t.Struct.Env.Pop()
			env := e.(map[interface{}]interface{})
			env[t.Fun.Param] = arg
			t.Struct.Env.Push(env)
			temp := v.stack
			v.stack = *t.Struct.Env
			result := t.Fun.Body.Accept(v)
			for k, v := range v.stack.Peek().(map[interface{}]interface{}) {
				se := t.Struct.Env.Pop().(map[interface{}]interface{})
				se[k] = v
				t.Struct.Env.Push(se)
			}
			v.stack = temp
			e = t.Struct.Env.Pop()
			env = e.(map[interface{}]interface{})
			delete(env, t.Fun.Param)
			t.Struct.Env.Push(env)
			return result
		default:
			return fmt.Sprintf("%s is not a function or method", ctx.GetFun().GetText())
		}
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
		if l, ok := left.(int); ok {
			if r, ok := right.(int); ok {
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
		if l, ok := left.(int); ok {
			if r, ok := right.(int); ok {
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
	id := ctx.ID().GetText()
	if hv, ok := v.heap[id]; ok {
		return hv
	}

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
	return fmt.Sprintf("%s is not defined", id)
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
	if c, ok := con.(bool); ok {
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
		if l, ok := left.(int); ok {
			if r, ok := right.(int); ok {
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
		if l, ok := left.(int); ok {
			if r, ok := right.(int); ok {
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
		if l, ok := left.(int); ok {
			if r, ok := right.(int); ok {
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
		if l, ok := left.(int); ok {
			if r, ok := right.(int); ok {
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
		if l, ok := left.(int); ok {
			if r, ok := right.(int); ok {
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
	id := ctx.GetId().GetText()
	value := ctx.GetValue().Accept(v)
	if value != nil {
		if s, ok := value.(StructValue); ok {
			v.heap[id] = s
			return nil
		}
		env := make(map[interface{}]interface{})
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
	fields := make(map[interface{}]interface{})
	for _, expr := range ctx.AllExpr() {
		e := expr.Accept(v)
		if field, ok := e.(FieldValue); ok {
			fields[field.ID] = field.Value
		} else {
			fmt.Printf("%s is not a field \n", expr.GetText())
		}
	}
	env := stack.New()
	env.Push(fields)
	return StructValue{env}
}

// VisitFieldExpr evaluates a field expression
func (v *Evaluator) VisitFieldExpr(ctx *parser.FieldExprContext) interface{} {
	id := ctx.GetId().GetText()
	field := ctx.GetValue().Accept(v)
	if _, ok := field.(FunValue); ok {
		return FieldValue{id, ctx.GetValue()}
	}
	return FieldValue{id, field}
}

// VisitDotExpr evaluates a dot expression
func (v *Evaluator) VisitDotExpr(ctx *parser.DotExprContext) interface{} {
	id := ctx.GetId().Accept(v)
	if id != nil {
		if s, ok := id.(StructValue); ok {
			f := ctx.GetField().GetText()
			env := s.Env.Peek()
			if env == nil {
				return nil
			}
			if value, ok := env.(map[interface{}]interface{})[f]; ok {
				c := string(f[0])
				if strings.ToUpper(c) == c {
					if fun, ok := value.(parser.IExprContext); ok {
						f := fun.Accept(v)
						return MethodValue{s, f.(FunValue)}
					}
					return value
				}
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
		if list, ok := id.([]interface{}); ok {
			if key, ok := ctx.GetKey().Accept(v).(int); ok {
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
		if list, ok := id.([]interface{}); ok {
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
		if list, ok := id.([]interface{}); ok {
			if key, ok := ctx.GetKey().Accept(v).(int); ok {
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

// VisitDotAssignExpr evaluates a dot lookup assignment expression
func (v *Evaluator) VisitDotAssignExpr(ctx *parser.DotAssignExprContext) interface{} {
	id := ctx.GetId().Accept(v)
	if id != nil {
		if s, ok := id.(StructValue); ok {
			f := ctx.GetField().GetText()
			key := ctx.GetId().GetText()
			value := ctx.GetValue().Accept(v)
			env := s.Env.Pop()
			env.(map[interface{}]interface{})[f] = value
			s.Env.Push(env)
			v.heap[key] = s
			return nil
		}
		return fmt.Sprintf("%s is not a struct", ctx.GetId().GetText())
	}
	return fmt.Sprintf("%s is not defined", ctx.GetId().GetText())
}
