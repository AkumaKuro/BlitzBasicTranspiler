package lexer

import (
	"fmt"
	"regexp"
	"strings"
)

type regexHandler = func(lex *lexer, regex *regexp.Regexp)

type pattern struct {
	regex   *regexp.Regexp
	handler regexHandler
}

type lexer struct {
	tokens   []Token
	source   string
	pos      int
	patterns []pattern
}

func (lex *lexer) advanceN(n int) {
	lex.pos += n
}

func (lex *lexer) push(token Token) {
	lex.tokens = append(lex.tokens, token)
}

func (lex *lexer) at() byte {
	return lex.source[lex.pos]
}

func (lex *lexer) remainder() string {
	return lex.source[lex.pos:]
}

func (lex *lexer) atEof() bool {
	return lex.pos >= len(lex.source)
}

func defaultHandler(kind TokenKind, value string) regexHandler {
	return func(lex *lexer, regex *regexp.Regexp) {
		lex.advanceN(len(value))
		lex.push(NewToken(kind, value))
	}
}

func newLineHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())
	lex.advanceN(len(match))
	lex.push(NewToken(NL, "\n"))
}

func floatHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())
	lex.push(NewToken(FLT_LIT, match))
	lex.advanceN(len(match))
}

func intHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())
	lex.push(NewToken(INT_LIT, match))
	lex.advanceN(len(match))
}

func stringHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	str := lex.remainder()[match[0]:match[1]]

	lex.push(NewToken(STR_LIT, str))
	lex.advanceN(len(str))
}

func commentHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())
	lex.push(NewToken(COMMENT, match))

	lex.advanceN(len(match))
}

func symbolHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())

	match = strings.ToLower(match)

	if kind, exists := reserved_keywords[match]; exists {
		lex.push(NewToken(kind, match))
	} else {
		lex.push(NewToken(IDENT, match))
	}

	lex.advanceN(len(match))
}

func skipHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	lex.advanceN(match[1])
}

func createLexer(source string) *lexer {
	return &lexer{
		tokens: make([]Token, 0),
		source: source,
		pos:    0,
		patterns: []pattern{
			{regexp.MustCompile(`\"[^\"]*\"`), stringHandler},
			{regexp.MustCompile(`\;.*`), commentHandler},
			{regexp.MustCompile(`[a-zA-Z]+`), symbolHandler},
			{regexp.MustCompile(`[0-9]+(\.[0-9]+)?`), floatHandler},
			{regexp.MustCompile(`[0-9]+`), intHandler},
			{regexp.MustCompile(`\(`), defaultHandler(OPEN_PR, "(")},
			{regexp.MustCompile(`\)`), defaultHandler(CLOSE_PR, ")")},
			{regexp.MustCompile(`\+`), defaultHandler(PLUS, "+")},
			{regexp.MustCompile(`\-`), defaultHandler(MINUS, "-")},
			{regexp.MustCompile(`\*`), defaultHandler(STAR, "*")},
			{regexp.MustCompile(`\<`), defaultHandler(LT, "<")},
			{regexp.MustCompile(`\>`), defaultHandler(GT, ">")},
			{regexp.MustCompile(`\=`), defaultHandler(EQ, "=")},
			{regexp.MustCompile(`\,`), defaultHandler(COMMA, ",")},
			{regexp.MustCompile(`\r?\n`), newLineHandler},
			{regexp.MustCompile(`\s+`), skipHandler},
		},
	}
}

func Tokenize(source string) []Token {
	lex := createLexer(source)

	for !lex.atEof() {
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
