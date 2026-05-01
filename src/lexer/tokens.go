package lexer

import "fmt"

type TokenKind int

const (
	EOF TokenKind = iota
	NUMBER

	ADD

	OPEN_BR
	CLOSE_BR
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
			"Type: %s, Value: %s",
			TokenKindString(token.kind),
			token.value,
		)
	}
}

func TokenKindString(kind TokenKind) string {
	switch kind {
	case EOF:
		return "eof"
	case ADD:
		return "+"
	case NUMBER:
		return "Num"
	case OPEN_BR:
		return "("
	case CLOSE_BR:
		return ")"
	default:
		println("Error, token %s not implemented", kind)
	}

	return ""
}
