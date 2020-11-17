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

func (v *BaseSimplVisitor) VisitApplyExpr(ctx *ApplyExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitFunExpr(ctx *FunExprContext) interface{} {
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

func (v *BaseSimplVisitor) VisitNumExpr(ctx *NumExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitLetExpr(ctx *LetExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitAddExpr(ctx *AddExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitDivExpr(ctx *DivExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitBoolExpr(ctx *BoolExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitLamdaExpr(ctx *LamdaExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitIdExpr(ctx *IdExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimplVisitor) VisitEqualExpr(ctx *EqualExprContext) interface{} {
	return v.VisitChildren(ctx)
}
