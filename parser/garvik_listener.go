// Code generated from Garvik.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Garvik

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GarvikListener is a complete listener for a parse tree produced by GarvikParser.
type GarvikListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStrExpr is called when entering the strExpr production.
	EnterStrExpr(c *StrExprContext)

	// EnterDotExpr is called when entering the dotExpr production.
	EnterDotExpr(c *DotExprContext)

	// EnterLenExpr is called when entering the lenExpr production.
	EnterLenExpr(c *LenExprContext)

	// EnterSubExpr is called when entering the subExpr production.
	EnterSubExpr(c *SubExprContext)

	// EnterParenExpr is called when entering the parenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterLetExpr is called when entering the letExpr production.
	EnterLetExpr(c *LetExprContext)

	// EnterGreaterExpr is called when entering the greaterExpr production.
	EnterGreaterExpr(c *GreaterExprContext)

	// EnterLambdaExpr is called when entering the lambdaExpr production.
	EnterLambdaExpr(c *LambdaExprContext)

	// EnterStructExpr is called when entering the structExpr production.
	EnterStructExpr(c *StructExprContext)

	// EnterDivExpr is called when entering the divExpr production.
	EnterDivExpr(c *DivExprContext)

	// EnterLookupAssignExpr is called when entering the lookupAssignExpr production.
	EnterLookupAssignExpr(c *LookupAssignExprContext)

	// EnterCallExpr is called when entering the callExpr production.
	EnterCallExpr(c *CallExprContext)

	// EnterListExpr is called when entering the listExpr production.
	EnterListExpr(c *ListExprContext)

	// EnterAssignExpr is called when entering the assignExpr production.
	EnterAssignExpr(c *AssignExprContext)

	// EnterIfElseExpr is called when entering the ifElseExpr production.
	EnterIfElseExpr(c *IfElseExprContext)

	// EnterLessExpr is called when entering the lessExpr production.
	EnterLessExpr(c *LessExprContext)

	// EnterDotAssignExpr is called when entering the dotAssignExpr production.
	EnterDotAssignExpr(c *DotAssignExprContext)

	// EnterMultExpr is called when entering the multExpr production.
	EnterMultExpr(c *MultExprContext)

	// EnterFieldExpr is called when entering the fieldExpr production.
	EnterFieldExpr(c *FieldExprContext)

	// EnterCommentExpr is called when entering the commentExpr production.
	EnterCommentExpr(c *CommentExprContext)

	// EnterNumExpr is called when entering the numExpr production.
	EnterNumExpr(c *NumExprContext)

	// EnterAddExpr is called when entering the addExpr production.
	EnterAddExpr(c *AddExprContext)

	// EnterLookupExpr is called when entering the lookupExpr production.
	EnterLookupExpr(c *LookupExprContext)

	// EnterBoolExpr is called when entering the boolExpr production.
	EnterBoolExpr(c *BoolExprContext)

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

	// ExitSubExpr is called when exiting the subExpr production.
	ExitSubExpr(c *SubExprContext)

	// ExitParenExpr is called when exiting the parenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitLetExpr is called when exiting the letExpr production.
	ExitLetExpr(c *LetExprContext)

	// ExitGreaterExpr is called when exiting the greaterExpr production.
	ExitGreaterExpr(c *GreaterExprContext)

	// ExitLambdaExpr is called when exiting the lambdaExpr production.
	ExitLambdaExpr(c *LambdaExprContext)

	// ExitStructExpr is called when exiting the structExpr production.
	ExitStructExpr(c *StructExprContext)

	// ExitDivExpr is called when exiting the divExpr production.
	ExitDivExpr(c *DivExprContext)

	// ExitLookupAssignExpr is called when exiting the lookupAssignExpr production.
	ExitLookupAssignExpr(c *LookupAssignExprContext)

	// ExitCallExpr is called when exiting the callExpr production.
	ExitCallExpr(c *CallExprContext)

	// ExitListExpr is called when exiting the listExpr production.
	ExitListExpr(c *ListExprContext)

	// ExitAssignExpr is called when exiting the assignExpr production.
	ExitAssignExpr(c *AssignExprContext)

	// ExitIfElseExpr is called when exiting the ifElseExpr production.
	ExitIfElseExpr(c *IfElseExprContext)

	// ExitLessExpr is called when exiting the lessExpr production.
	ExitLessExpr(c *LessExprContext)

	// ExitDotAssignExpr is called when exiting the dotAssignExpr production.
	ExitDotAssignExpr(c *DotAssignExprContext)

	// ExitMultExpr is called when exiting the multExpr production.
	ExitMultExpr(c *MultExprContext)

	// ExitFieldExpr is called when exiting the fieldExpr production.
	ExitFieldExpr(c *FieldExprContext)

	// ExitCommentExpr is called when exiting the commentExpr production.
	ExitCommentExpr(c *CommentExprContext)

	// ExitNumExpr is called when exiting the numExpr production.
	ExitNumExpr(c *NumExprContext)

	// ExitAddExpr is called when exiting the addExpr production.
	ExitAddExpr(c *AddExprContext)

	// ExitLookupExpr is called when exiting the lookupExpr production.
	ExitLookupExpr(c *LookupExprContext)

	// ExitBoolExpr is called when exiting the boolExpr production.
	ExitBoolExpr(c *BoolExprContext)

	// ExitIdExpr is called when exiting the idExpr production.
	ExitIdExpr(c *IdExprContext)

	// ExitEqualExpr is called when exiting the equalExpr production.
	ExitEqualExpr(c *EqualExprContext)
}
