// Code generated from Garvik.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Garvik

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by GarvikParser.
type GarvikVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GarvikParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by GarvikParser#strExpr.
	VisitStrExpr(ctx *StrExprContext) interface{}

	// Visit a parse tree produced by GarvikParser#dotExpr.
	VisitDotExpr(ctx *DotExprContext) interface{}

	// Visit a parse tree produced by GarvikParser#lenExpr.
	VisitLenExpr(ctx *LenExprContext) interface{}

	// Visit a parse tree produced by GarvikParser#ifElseExpr.
	VisitIfElseExpr(ctx *IfElseExprContext) interface{}

	// Visit a parse tree produced by GarvikParser#subExpr.
	VisitSubExpr(ctx *SubExprContext) interface{}

	// Visit a parse tree produced by GarvikParser#multExpr.
	VisitMultExpr(ctx *MultExprContext) interface{}

	// Visit a parse tree produced by GarvikParser#commentExpr.
	VisitCommentExpr(ctx *CommentExprContext) interface{}

	// Visit a parse tree produced by GarvikParser#parenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by GarvikParser#letExpr.
	VisitLetExpr(ctx *LetExprContext) interface{}

	// Visit a parse tree produced by GarvikParser#numExpr.
	VisitNumExpr(ctx *NumExprContext) interface{}

	// Visit a parse tree produced by GarvikParser#addExpr.
	VisitAddExpr(ctx *AddExprContext) interface{}

	// Visit a parse tree produced by GarvikParser#lambdaExpr.
	VisitLambdaExpr(ctx *LambdaExprContext) interface{}

	// Visit a parse tree produced by GarvikParser#structExpr.
	VisitStructExpr(ctx *StructExprContext) interface{}

	// Visit a parse tree produced by GarvikParser#lookupExpr.
	VisitLookupExpr(ctx *LookupExprContext) interface{}

	// Visit a parse tree produced by GarvikParser#divExpr.
	VisitDivExpr(ctx *DivExprContext) interface{}

	// Visit a parse tree produced by GarvikParser#lookupAssignExpr.
	VisitLookupAssignExpr(ctx *LookupAssignExprContext) interface{}

	// Visit a parse tree produced by GarvikParser#boolExpr.
	VisitBoolExpr(ctx *BoolExprContext) interface{}

	// Visit a parse tree produced by GarvikParser#callExpr.
	VisitCallExpr(ctx *CallExprContext) interface{}

	// Visit a parse tree produced by GarvikParser#listExpr.
	VisitListExpr(ctx *ListExprContext) interface{}

	// Visit a parse tree produced by GarvikParser#assignExpr.
	VisitAssignExpr(ctx *AssignExprContext) interface{}

	// Visit a parse tree produced by GarvikParser#idExpr.
	VisitIdExpr(ctx *IdExprContext) interface{}

	// Visit a parse tree produced by GarvikParser#equalExpr.
	VisitEqualExpr(ctx *EqualExprContext) interface{}
}
