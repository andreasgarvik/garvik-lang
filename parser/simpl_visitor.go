// Code generated from Simpl.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Simpl

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by SimplParser.
type SimplVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SimplParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by SimplParser#strExpr.
	VisitStrExpr(ctx *StrExprContext) interface{}

	// Visit a parse tree produced by SimplParser#dotExpr.
	VisitDotExpr(ctx *DotExprContext) interface{}

	// Visit a parse tree produced by SimplParser#lenExpr.
	VisitLenExpr(ctx *LenExprContext) interface{}

	// Visit a parse tree produced by SimplParser#ifElseExpr.
	VisitIfElseExpr(ctx *IfElseExprContext) interface{}

	// Visit a parse tree produced by SimplParser#subExpr.
	VisitSubExpr(ctx *SubExprContext) interface{}

	// Visit a parse tree produced by SimplParser#multExpr.
	VisitMultExpr(ctx *MultExprContext) interface{}

	// Visit a parse tree produced by SimplParser#commentExpr.
	VisitCommentExpr(ctx *CommentExprContext) interface{}

	// Visit a parse tree produced by SimplParser#parenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by SimplParser#numExpr.
	VisitNumExpr(ctx *NumExprContext) interface{}

	// Visit a parse tree produced by SimplParser#addExpr.
	VisitAddExpr(ctx *AddExprContext) interface{}

	// Visit a parse tree produced by SimplParser#lambdaExpr.
	VisitLambdaExpr(ctx *LambdaExprContext) interface{}

	// Visit a parse tree produced by SimplParser#structExpr.
	VisitStructExpr(ctx *StructExprContext) interface{}

	// Visit a parse tree produced by SimplParser#lookupExpr.
	VisitLookupExpr(ctx *LookupExprContext) interface{}

	// Visit a parse tree produced by SimplParser#divExpr.
	VisitDivExpr(ctx *DivExprContext) interface{}

	// Visit a parse tree produced by SimplParser#lookupAssignExpr.
	VisitLookupAssignExpr(ctx *LookupAssignExprContext) interface{}

	// Visit a parse tree produced by SimplParser#boolExpr.
	VisitBoolExpr(ctx *BoolExprContext) interface{}

	// Visit a parse tree produced by SimplParser#callExpr.
	VisitCallExpr(ctx *CallExprContext) interface{}

	// Visit a parse tree produced by SimplParser#listExpr.
	VisitListExpr(ctx *ListExprContext) interface{}

	// Visit a parse tree produced by SimplParser#assignExpr.
	VisitAssignExpr(ctx *AssignExprContext) interface{}

	// Visit a parse tree produced by SimplParser#idExpr.
	VisitIdExpr(ctx *IdExprContext) interface{}

	// Visit a parse tree produced by SimplParser#equalExpr.
	VisitEqualExpr(ctx *EqualExprContext) interface{}
}
