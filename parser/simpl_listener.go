// Code generated from Simpl.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Simpl

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SimplListener is a complete listener for a parse tree produced by SimplParser.
type SimplListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStrExpr is called when entering the strExpr production.
	EnterStrExpr(c *StrExprContext)

	// EnterDotExpr is called when entering the dotExpr production.
	EnterDotExpr(c *DotExprContext)

	// EnterLenExpr is called when entering the lenExpr production.
	EnterLenExpr(c *LenExprContext)

	// EnterIfElseExpr is called when entering the ifElseExpr production.
	EnterIfElseExpr(c *IfElseExprContext)

	// EnterSubExpr is called when entering the subExpr production.
	EnterSubExpr(c *SubExprContext)

	// EnterMultExpr is called when entering the multExpr production.
	EnterMultExpr(c *MultExprContext)

	// EnterCommentExpr is called when entering the commentExpr production.
	EnterCommentExpr(c *CommentExprContext)

	// EnterParenExpr is called when entering the parenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterNumExpr is called when entering the numExpr production.
	EnterNumExpr(c *NumExprContext)

	// EnterAddExpr is called when entering the addExpr production.
	EnterAddExpr(c *AddExprContext)

	// EnterLambdaExpr is called when entering the lambdaExpr production.
	EnterLambdaExpr(c *LambdaExprContext)

	// EnterStructExpr is called when entering the structExpr production.
	EnterStructExpr(c *StructExprContext)

	// EnterLookupExpr is called when entering the lookupExpr production.
	EnterLookupExpr(c *LookupExprContext)

	// EnterDivExpr is called when entering the divExpr production.
	EnterDivExpr(c *DivExprContext)

	// EnterLookupAssignExpr is called when entering the lookupAssignExpr production.
	EnterLookupAssignExpr(c *LookupAssignExprContext)

	// EnterBoolExpr is called when entering the boolExpr production.
	EnterBoolExpr(c *BoolExprContext)

	// EnterCallExpr is called when entering the callExpr production.
	EnterCallExpr(c *CallExprContext)

	// EnterListExpr is called when entering the listExpr production.
	EnterListExpr(c *ListExprContext)

	// EnterAssignExpr is called when entering the assignExpr production.
	EnterAssignExpr(c *AssignExprContext)

	// EnterIdExpr is called when entering the idExpr production.
	EnterIdExpr(c *IdExprContext)

	// EnterEqualExpr is called when entering the equalExpr production.
	EnterEqualExpr(c *EqualExprContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStrExpr is called when exiting the strExpr production.
	ExitStrExpr(c *StrExprContext)

	// ExitDotExpr is called when exiting the dotExpr production.
	ExitDotExpr(c *DotExprContext)

	// ExitLenExpr is called when exiting the lenExpr production.
	ExitLenExpr(c *LenExprContext)

	// ExitIfElseExpr is called when exiting the ifElseExpr production.
	ExitIfElseExpr(c *IfElseExprContext)

	// ExitSubExpr is called when exiting the subExpr production.
	ExitSubExpr(c *SubExprContext)

	// ExitMultExpr is called when exiting the multExpr production.
	ExitMultExpr(c *MultExprContext)

	// ExitCommentExpr is called when exiting the commentExpr production.
	ExitCommentExpr(c *CommentExprContext)

	// ExitParenExpr is called when exiting the parenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitNumExpr is called when exiting the numExpr production.
	ExitNumExpr(c *NumExprContext)

	// ExitAddExpr is called when exiting the addExpr production.
	ExitAddExpr(c *AddExprContext)

	// ExitLambdaExpr is called when exiting the lambdaExpr production.
	ExitLambdaExpr(c *LambdaExprContext)

	// ExitStructExpr is called when exiting the structExpr production.
	ExitStructExpr(c *StructExprContext)

	// ExitLookupExpr is called when exiting the lookupExpr production.
	ExitLookupExpr(c *LookupExprContext)

	// ExitDivExpr is called when exiting the divExpr production.
	ExitDivExpr(c *DivExprContext)

	// ExitLookupAssignExpr is called when exiting the lookupAssignExpr production.
	ExitLookupAssignExpr(c *LookupAssignExprContext)

	// ExitBoolExpr is called when exiting the boolExpr production.
	ExitBoolExpr(c *BoolExprContext)

	// ExitCallExpr is called when exiting the callExpr production.
	ExitCallExpr(c *CallExprContext)

	// ExitListExpr is called when exiting the listExpr production.
	ExitListExpr(c *ListExprContext)

	// ExitAssignExpr is called when exiting the assignExpr production.
	ExitAssignExpr(c *AssignExprContext)

	// ExitIdExpr is called when exiting the idExpr production.
	ExitIdExpr(c *IdExprContext)

	// ExitEqualExpr is called when exiting the equalExpr production.
	ExitEqualExpr(c *EqualExprContext)
}
