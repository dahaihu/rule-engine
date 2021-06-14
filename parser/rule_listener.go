// Code generated from /home/studing/GolandProjects/src/rule-engine/parser/Rule.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Rule

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RuleListener is a complete listener for a parse tree produced by RuleParser.
type RuleListener interface {
	antlr.ParseTreeListener

	// EnterAttrPath is called when entering the attrPath production.
	EnterAttrPath(c *AttrPathContext)

	// EnterSubAttr is called when entering the subAttr production.
	EnterSubAttr(c *SubAttrContext)

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterFloat is called when entering the Float production.
	EnterFloat(c *FloatContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterNegativeValue is called when entering the NegativeValue production.
	EnterNegativeValue(c *NegativeValueContext)

	// EnterArithmeticWithParenthesis is called when entering the ArithmeticWithParenthesis production.
	EnterArithmeticWithParenthesis(c *ArithmeticWithParenthesisContext)

	// EnterInt is called when entering the Int production.
	EnterInt(c *IntContext)

	// EnterAttr is called when entering the Attr production.
	EnterAttr(c *AttrContext)

	// EnterCompareWithParenthesis is called when entering the CompareWithParenthesis production.
	EnterCompareWithParenthesis(c *CompareWithParenthesisContext)

	// EnterBasicCompare is called when entering the BasicCompare production.
	EnterBasicCompare(c *BasicCompareContext)

	// EnterBasicLogical is called when entering the BasicLogical production.
	EnterBasicLogical(c *BasicLogicalContext)

	// EnterTwoLogical is called when entering the TwoLogical production.
	EnterTwoLogical(c *TwoLogicalContext)

	// EnterLogicalWithParenthesis is called when entering the LogicalWithParenthesis production.
	EnterLogicalWithParenthesis(c *LogicalWithParenthesisContext)

	// ExitAttrPath is called when exiting the attrPath production.
	ExitAttrPath(c *AttrPathContext)

	// ExitSubAttr is called when exiting the subAttr production.
	ExitSubAttr(c *SubAttrContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitFloat is called when exiting the Float production.
	ExitFloat(c *FloatContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitNegativeValue is called when exiting the NegativeValue production.
	ExitNegativeValue(c *NegativeValueContext)

	// ExitArithmeticWithParenthesis is called when exiting the ArithmeticWithParenthesis production.
	ExitArithmeticWithParenthesis(c *ArithmeticWithParenthesisContext)

	// ExitInt is called when exiting the Int production.
	ExitInt(c *IntContext)

	// ExitAttr is called when exiting the Attr production.
	ExitAttr(c *AttrContext)

	// ExitCompareWithParenthesis is called when exiting the CompareWithParenthesis production.
	ExitCompareWithParenthesis(c *CompareWithParenthesisContext)

	// ExitBasicCompare is called when exiting the BasicCompare production.
	ExitBasicCompare(c *BasicCompareContext)

	// ExitBasicLogical is called when exiting the BasicLogical production.
	ExitBasicLogical(c *BasicLogicalContext)

	// ExitTwoLogical is called when exiting the TwoLogical production.
	ExitTwoLogical(c *TwoLogicalContext)

	// ExitLogicalWithParenthesis is called when exiting the LogicalWithParenthesis production.
	ExitLogicalWithParenthesis(c *LogicalWithParenthesisContext)
}
