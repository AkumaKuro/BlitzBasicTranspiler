package parser

import (
	"github.com/AkumaKuro/BlitzBasicTranspiler/src/ast"
)

func parseStmt(p *parser) ast.Stmt {
	stmtFn, exists := stmtLU[p.currentTokenKind()]

	if exists {
		return stmtFn(p)
	}

	expression := parseExpr(p, defaultBP)
	//p.expect(lexer.NL)

	return ast.ExprStmt{
		Expr: expression,
	}
}
