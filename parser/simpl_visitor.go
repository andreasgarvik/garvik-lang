// Code generated from Simpl.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Simpl

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by SimplParser.
type SimplVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SimplParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by SimplParser#addExpr.
	VisitAddExpr(ctx *AddExprContext) interface{}

	// Visit a parse tree produced by SimplParser#multExpr.
	VisitMultExpr(ctx *MultExprContext) interface{}

	// Visit a parse tree produced by SimplParser#numExpr.
	VisitNumExpr(ctx *NumExprContext) interface{}

	// Visit a parse tree produced by SimplParser#parenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by SimplParser#letExpr.
	VisitLetExpr(ctx *LetExprContext) interface{}

	// Visit a parse tree produced by SimplParser#idExpr.
	VisitIdExpr(ctx *IdExprContext) interface{}
}
