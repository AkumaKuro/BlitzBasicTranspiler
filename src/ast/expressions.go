package ast

import "github.com/AkumaKuro/BlitzBasicTranspiler/src/lexer"

type IntExpr struct {
	Value int32
}

func (n IntExpr) expr() {

}

type IdentExpr struct {
	Value string
}

func (n IdentExpr) expr() {}

// COMPLEX EXPR

type BinaryExpr struct {
	Left     Expr
	Right    Expr
	Operator lexer.Token
}

func (n BinaryExpr) expr() {}
