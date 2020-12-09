// Code generated from Simpl.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Simpl

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSimplVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSimplVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitStrExpr(ctx *StrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitDotExpr(ctx *DotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitLenExpr(ctx *LenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitIfElseExpr(ctx *IfElseExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitSubExpr(ctx *SubExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitMultExpr(ctx *MultExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitCommentExpr(ctx *CommentExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitNumExpr(ctx *NumExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitAddExpr(ctx *AddExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitLambdaExpr(ctx *LambdaExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitStructExpr(ctx *StructExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitLookupExpr(ctx *LookupExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitDivExpr(ctx *DivExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitLookupAssignExpr(ctx *LookupAssignExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitBoolExpr(ctx *BoolExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitCallExpr(ctx *CallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitListExpr(ctx *ListExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitAssignExpr(ctx *AssignExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitIdExpr(ctx *IdExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitEqualExpr(ctx *EqualExprContext) interface{} {
	return v.VisitChildren(ctx)
}
