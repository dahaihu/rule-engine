// Code generated from /home/studing/GolandProjects/src/rule-engine/parser/Rule.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Rule

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRuleListener is a complete listener for a parse tree produced by RuleParser.
type BaseRuleListener struct{}

var _ RuleListener = &BaseRuleListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRuleListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRuleListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRuleListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRuleListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterAttrPath is called when production attrPath is entered.
func (s *BaseRuleListener) EnterAttrPath(ctx *AttrPathContext) {}

// ExitAttrPath is called when production attrPath is exited.
func (s *BaseRuleListener) ExitAttrPath(ctx *AttrPathContext) {}

// EnterSubAttr is called when production subAttr is entered.
func (s *BaseRuleListener) EnterSubAttr(ctx *SubAttrContext) {}

// ExitSubAttr is called when production subAttr is exited.
func (s *BaseRuleListener) ExitSubAttr(ctx *SubAttrContext) {}

// EnterStart is called when production start is entered.
func (s *BaseRuleListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseRuleListener) ExitStart(ctx *StartContext) {}

// EnterFloat is called when production Float is entered.
func (s *BaseRuleListener) EnterFloat(ctx *FloatContext) {}

// ExitFloat is called when production Float is exited.
func (s *BaseRuleListener) ExitFloat(ctx *FloatContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseRuleListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseRuleListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseRuleListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseRuleListener) ExitAddSub(ctx *AddSubContext) {}

// EnterNegativeValue is called when production NegativeValue is entered.
func (s *BaseRuleListener) EnterNegativeValue(ctx *NegativeValueContext) {}

// ExitNegativeValue is called when production NegativeValue is exited.
func (s *BaseRuleListener) ExitNegativeValue(ctx *NegativeValueContext) {}

// EnterArithmeticWithParenthesis is called when production ArithmeticWithParenthesis is entered.
func (s *BaseRuleListener) EnterArithmeticWithParenthesis(ctx *ArithmeticWithParenthesisContext) {}

// ExitArithmeticWithParenthesis is called when production ArithmeticWithParenthesis is exited.
func (s *BaseRuleListener) ExitArithmeticWithParenthesis(ctx *ArithmeticWithParenthesisContext) {}

// EnterInt is called when production Int is entered.
func (s *BaseRuleListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production Int is exited.
func (s *BaseRuleListener) ExitInt(ctx *IntContext) {}

// EnterAttr is called when production Attr is entered.
func (s *BaseRuleListener) EnterAttr(ctx *AttrContext) {}

// ExitAttr is called when production Attr is exited.
func (s *BaseRuleListener) ExitAttr(ctx *AttrContext) {}

// EnterCompareWithParenthesis is called when production CompareWithParenthesis is entered.
func (s *BaseRuleListener) EnterCompareWithParenthesis(ctx *CompareWithParenthesisContext) {}

// ExitCompareWithParenthesis is called when production CompareWithParenthesis is exited.
func (s *BaseRuleListener) ExitCompareWithParenthesis(ctx *CompareWithParenthesisContext) {}

// EnterBasicCompare is called when production BasicCompare is entered.
func (s *BaseRuleListener) EnterBasicCompare(ctx *BasicCompareContext) {}

// ExitBasicCompare is called when production BasicCompare is exited.
func (s *BaseRuleListener) ExitBasicCompare(ctx *BasicCompareContext) {}

// EnterBasicLogical is called when production BasicLogical is entered.
func (s *BaseRuleListener) EnterBasicLogical(ctx *BasicLogicalContext) {}

// ExitBasicLogical is called when production BasicLogical is exited.
func (s *BaseRuleListener) ExitBasicLogical(ctx *BasicLogicalContext) {}

// EnterTwoLogical is called when production TwoLogical is entered.
func (s *BaseRuleListener) EnterTwoLogical(ctx *TwoLogicalContext) {}

// ExitTwoLogical is called when production TwoLogical is exited.
func (s *BaseRuleListener) ExitTwoLogical(ctx *TwoLogicalContext) {}

// EnterLogicalWithParenthesis is called when production LogicalWithParenthesis is entered.
func (s *BaseRuleListener) EnterLogicalWithParenthesis(ctx *LogicalWithParenthesisContext) {}

// ExitLogicalWithParenthesis is called when production LogicalWithParenthesis is exited.
func (s *BaseRuleListener) ExitLogicalWithParenthesis(ctx *LogicalWithParenthesisContext) {}
