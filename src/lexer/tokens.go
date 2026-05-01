package lexer

import "fmt"

type TokenKind int

const (
	EOF TokenKind = iota
	NUMBER

	PLUS
	MINUS
	STAR

	OPEN_PR
	CLOSE_PR
)

type Token struct {
	kind  TokenKind
	value string
}

func NewToken(kind TokenKind, value string) Token {
	return Token{
		kind,
		value,
	}
}

func (token Token) IsOneOfMany(expectedKinds ...TokenKind) bool {
	for _, expected := range expectedKinds {
		if token.kind == expected {
			return true
		}
	}

	return false
}

func (token Token) Print() {
	if token.kind == NUMBER {
		fmt.Printf(
			"Type: %s, Value: %s\n",
			TokenKindString(token.kind),
			token.value,
		)
	} else {
		fmt.Printf("Type: %s\n", TokenKindString(token.kind))
	}
}

func TokenKindString(kind TokenKind) string {
	switch kind {
	case EOF:
		return "eof"
	case PLUS:
		return "+"
	case MINUS:
		return "-"
	case STAR:
		return "*"
	case NUMBER:
		return "Num"
	case OPEN_PR:
		return "("
	case CLOSE_PR:
		return ")"
	default:
		println("Error, token %s not implemented", kind)
	}

	return ""
}
