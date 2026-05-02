package parser

import (
	"github.com/AkumaKuro/BlitzBasicTranspiler/src/ast"
	"github.com/AkumaKuro/BlitzBasicTranspiler/src/lexer"
)

type bindingPower int

const (
	defaultBP bindingPower = iota
	comma
	assignment
	logical
	relational
	additive
	multiplicative
	unary
	call
	member
	primary
)

type stmtHandler func(p *parser) ast.Stmt
type nudHandler func(p *parser) ast.Expr
type ledHandler func(p *parser, left ast.Expr, bp bindingPower) ast.Expr

type stmtLookup map[lexer.TokenKind]stmtHandler
type nudLookup map[lexer.TokenKind]nudHandler
type ledLookup map[lexer.TokenKind]ledHandler
type bpLookup map[lexer.TokenKind]bindingPower

var bpLU = bpLookup{}
var nudLU = nudLookup{}
var ledLU = ledLookup{}
var stmtLU = stmtLookup{}

func led(kind lexer.TokenKind, bp bindingPower, ledFn ledHandler) {
	bpLU[kind] = bp
	ledLU[kind] = ledFn
}

func nud(kind lexer.TokenKind, bp bindingPower, nudFn nudHandler) {
	bpLU[kind] = bp
	nudLU[kind] = nudFn
}

func stmt(kind lexer.TokenKind, stmtFn stmtHandler) {
	bpLU[kind] = defaultBP
	stmtLU[kind] = stmtFn
}

func createTokenLookups() {
	led(lexer.PLUS, additive, parseBinaryExpr)

	led(lexer.STAR, multiplicative, parseBinaryExpr)

	nud(lexer.INT_LIT, primary, parsePrimaryExpr)
	nud(lexer.IDENT, primary, parsePrimaryExpr)

}
