package main

import "github.com/chettriyuvraj/lox-crafting-interpreters/pkg/scanner"

type ExprVisitor interface {
	// visitAssignExpr(expr *AssignExpr) (any, error)
	visitBinaryExpr(expr *BinaryExpr) (any, error)
	// visitCallExpr(expr *CallExpr) (any, error)
	// visitGetExpr(expr *GetExpr) (any, error)
	// visitGroupingExpr(expr *GroupingExpr) (any, error)
	// visitLiteralExpr(expr *LiteralExpr) (any, error)
	// visitLogicalExpr(expr *LogicalExpr) (any, error)
	// visitSetExpr(expr *SetExpr) (any, error)
	// visitSuperExpr(expr *SuperExpr) (any, error)
	// visitThisExpr(expr *ThisExpr) (any, error)
	// visitUnaryExpr(expr *UnaryExpr) (any, error)
	// visitVariableExpr(expr *VariableExpr) (any, error)
}

type Expr interface {
	accept(visitor ExprVisitor) (any, error)
}

type BinaryExpr struct {
	left     Expr
	operator *scanner.Token
	right    Expr
}

func NewBinaryExpr(left Expr, operator *scanner.Token, right Expr) *BinaryExpr {
	return &BinaryExpr{
		left,
		operator,
		right,
	}
}

func (expr *BinaryExpr) accept(visitor ExprVisitor) (any, error) {
	return visitor.visitBinaryExpr(expr)
}
