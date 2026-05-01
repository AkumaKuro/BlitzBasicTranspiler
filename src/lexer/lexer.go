package lexer

import (
	"fmt"
	"regexp"
)

type RegexHandler = func(lex *Lexer, regex *regexp.Regexp)

type Pattern struct {
	regex   *regexp.Regexp
	handler RegexHandler
}

type Lexer struct {
	tokens   []Token
	source   string
	pos      int
	patterns []Pattern
}

func (lex *Lexer) advanceN(n int) {
	lex.pos += n
}

func (lex *Lexer) push(token Token) {
	lex.tokens = append(lex.tokens, token)
}

func (lex *Lexer) at() byte {
	return lex.source[lex.pos]
}

func (lex *Lexer) remainder() string {
	return lex.source[lex.pos:]
}

func (lex *Lexer) at_eof() bool {
	return lex.pos >= len(lex.source)
}

func DefaultHandler(kind TokenKind, value string) RegexHandler {
	return func(lex *Lexer, regex *regexp.Regexp) {
		lex.advanceN(len(value))
		lex.push(NewToken(kind, value))
	}
}

func NumberHandler(lex *Lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())
	lex.push(NewToken(NUMBER, match))
	lex.advanceN(len(match))
}

func SkipHandler(lex *Lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	lex.advanceN(match[1])
}

func CreateLexer(source string) *Lexer {
	return &Lexer{
		tokens: make([]Token, 0),
		source: source,
		pos:    0,
		patterns: []Pattern{
			{regexp.MustCompile(`\s+`), SkipHandler},
			{regexp.MustCompile(`[0-9]+(\.[0-9]+)?`), NumberHandler},
			{regexp.MustCompile(`\(`), DefaultHandler(OPEN_PR, "(")},
			{regexp.MustCompile(`\)`), DefaultHandler(CLOSE_PR, ")")},
			{regexp.MustCompile(`\+`), DefaultHandler(PLUS, "+")},
			{regexp.MustCompile(`\-`), DefaultHandler(MINUS, "-")},
			{regexp.MustCompile(`\*`), DefaultHandler(STAR, "*")},
		},
	}
}

func Tokenize(source string) []Token {
	lex := CreateLexer(source)

	for !lex.at_eof() {
		matched := false

		for _, pattern := range lex.patterns {
			loc := pattern.regex.FindStringIndex(lex.remainder())

			if loc != nil && loc[0] == 0 {
				pattern.handler(lex, pattern.regex)
				matched = true
				break
			}
		}

		if !matched {
			panic(fmt.Sprintf(
				"Lexer error: Unrecognized token near %s",
				lex.remainder(),
			))
		}
	}

	lex.push(NewToken(EOF, ""))

	return lex.tokens
}
