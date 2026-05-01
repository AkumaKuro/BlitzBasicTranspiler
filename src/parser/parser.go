package parser

import (
	"github.com/AkumaKuro/BlitzBasicTranspiler/src/ast"
	"github.com/AkumaKuro/BlitzBasicTranspiler/src/lexer"
)

type parser struct {
	tokens []lexer.Token
	pos    int
}

func Parse(tokens []lexer.Token) ast.BlockStmt {
	Body := make([]ast.Stmt, 0)
	p := createParser(tokens)

	for p.hasTokens() {
		Body = append(Body, parseStmt(p))
	}

	return ast.BlockStmt{
		Body: Body,
	}
}

func createParser(tokens []lexer.Token) *parser {
	return &parser{
		tokens: tokens,
	}
}

func (p *parser) currentToken() lexer.Token {
	return p.tokens[p.pos]
}

func (p *parser) advance() lexer.Token {
	token := p.currentToken()
	p.pos += 1
	return token
}

func (p *parser) hasTokens() bool {
	return p.pos < len(p.tokens) && p.currentTokenKind() != lexer.EOF
}

func (p *parser) currentTokenKind() lexer.TokenKind {
	return p.currentToken().Kind
}
