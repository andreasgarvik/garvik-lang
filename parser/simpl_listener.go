// Code generated from Simpl.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Simpl

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SimplListener is a complete listener for a parse tree produced by SimplParser.
type SimplListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterApplyExpr is called when entering the applyExpr production.
	EnterApplyExpr(c *ApplyExprContext)

	// EnterFunExpr is called when entering the funExpr production.
	EnterFunExpr(c *FunExprContext)

	// EnterAddExpr is called when entering the addExpr production.
	EnterAddExpr(c *AddExprContext)

	// EnterIfElseExpr is called when entering the ifElseExpr production.
	EnterIfElseExpr(c *IfElseExprContext)

	// EnterDivExpr is called when entering the divExpr production.
	EnterDivExpr(c *DivExprContext)

	// EnterSubExpr is called when entering the subExpr production.
	EnterSubExpr(c *SubExprContext)

	// EnterMultExpr is called when entering the multExpr production.
	EnterMultExpr(c *MultExprContext)

	// EnterNumExpr is called when entering the numExpr production.
	EnterNumExpr(c *NumExprContext)

	// EnterParenExpr is called when entering the parenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterLetExpr is called when entering the letExpr production.
	EnterLetExpr(c *LetExprContext)

	// EnterIdExpr is called when entering the idExpr production.
	EnterIdExpr(c *IdExprContext)

	// EnterEqualExpr is called when entering the equalExpr production.
	EnterEqualExpr(c *EqualExprContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitApplyExpr is called when exiting the applyExpr production.
	ExitApplyExpr(c *ApplyExprContext)

	// ExitFunExpr is called when exiting the funExpr production.
	ExitFunExpr(c *FunExprContext)

	// ExitAddExpr is called when exiting the addExpr production.
	ExitAddExpr(c *AddExprContext)

	// ExitIfElseExpr is called when exiting the ifElseExpr production.
	ExitIfElseExpr(c *IfElseExprContext)

	// ExitDivExpr is called when exiting the divExpr production.
	ExitDivExpr(c *DivExprContext)

	// ExitSubExpr is called when exiting the subExpr production.
	ExitSubExpr(c *SubExprContext)

	// ExitMultExpr is called when exiting the multExpr production.
	ExitMultExpr(c *MultExprContext)

	// ExitNumExpr is called when exiting the numExpr production.
	ExitNumExpr(c *NumExprContext)

	// ExitParenExpr is called when exiting the parenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitLetExpr is called when exiting the letExpr production.
	ExitLetExpr(c *LetExprContext)

	// ExitIdExpr is called when exiting the idExpr production.
	ExitIdExpr(c *IdExprContext)

	// ExitEqualExpr is called when exiting the equalExpr production.
	ExitEqualExpr(c *EqualExprContext)
}
