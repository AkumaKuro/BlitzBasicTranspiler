package parser

import (
	"fmt"
	"strconv"

	"github.com/AkumaKuro/BlitzBasicTranspiler/src/ast"
	"github.com/AkumaKuro/BlitzBasicTranspiler/src/lexer"
)

func parsePrimaryExpr(p *parser) ast.Expr {
	switch p.currentTokenKind() {
	case lexer.INT_LIT:
		number, _ := strconv.ParseInt(p.advance().Value, 10, 32)
		return ast.IntExpr{
			Value: int32(number),
		}
	case lexer.IDENT:
		return ast.IdentExpr{
			Value: p.advance().Value,
		}
	default:
		panic(fmt.Sprintf(
			"Cannot create primary expr from %s\n",
			lexer.TokenKindString(p.currentTokenKind()),
		))
	}
}

func parseBinaryExpr(p *parser, left ast.Expr, bp bindingPower) ast.Expr {
	operatorToken := p.advance()
	right := parseExpr(p, bp)

	return ast.BinaryExpr{
		Left:     left,
		Operator: operatorToken,
		Right:    right,
	}
}

func parseExpr(p *parser, bp bindingPower) ast.Expr {
	tokenKind := p.currentTokenKind()
	nudFn, exists := nudLU[tokenKind]

	if !exists {
		panic(fmt.Sprintf(
			"NUD Handler expected for Token %s\n",
			lexer.TokenKindString(tokenKind),
		))
	}

	left := nudFn(p)
	for bpLU[p.currentTokenKind()] > bp {
		tokenKind = p.currentTokenKind()
		ledFn, exists := ledLU[tokenKind]

		if !exists {
			panic(fmt.Sprintf(
				"LED Handler expected for Token %s\n",
				lexer.TokenKindString(tokenKind),
			))
		}

		left = ledFn(p, left, bp)
	}

	return left
}
