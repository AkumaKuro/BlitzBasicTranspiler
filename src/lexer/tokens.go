package lexer

import "fmt"

type TokenKind int

const (
	EOF TokenKind = iota
	COMMENT

	COMMA

	FLT_LIT
	STR_LIT
	INT_LIT

	IDENT

	PLUS
	MINUS
	STAR
	SLASH
	NOT
	MOD

	SAR
	SHL
	SHR

	EQ
	LT
	GT
	GE
	LE
	NE

	OPEN_PR
	CLOSE_PR

	TYPE
	FIELD

	AFTER
	BEFORE
	EACH
	FIRST
	DELETE
	LAST
	INSERT
	NEXT

	AND
	OR
	XOR

	SELECT
	CASE
	DEFAULT

	CONST
	LOCAL
	GLOBAL
	DIM
	DATA
	RESTORE
	READ

	GOSUB
	GOTO

	IF
	ELSE
	ELSEIF
	ENDIF
	THEN

	FOR
	WHILE
	FOREVER
	DO
	REPEAT
	EXIT
	WEND
	TO
	STEP
	UNTIL

	INT
	STR
	FLOAT
	NEW
	NULL
	PI

	NL

	FALSE
	TRUE

	FUNCTION
	RETURN

	INCLUDE
	END

	OPEN_BR
	CLOSE_BR

	INT_SIG
	FLT_SIG
	STR_SIG

	HEX_INT

	COLON
	DOT
	BSLASH
	EXP
)

var reserved_keywords map[string]TokenKind = map[string]TokenKind{
	"after":    AFTER,
	"and":      AND,
	"before":   BEFORE,
	"case":     CASE,
	"const":    CONST,
	"data":     DATA,
	"default":  DEFAULT,
	"delete":   DELETE,
	"dim":      DIM,
	"each":     EACH,
	"else":     ELSE,
	"elseif":   ELSEIF,
	"end":      END,
	"endif":    ENDIF,
	"exit":     EXIT,
	"false":    FALSE,
	"field":    FIELD,
	"first":    FIRST,
	"float":    FLOAT,
	"for":      FOR,
	"forever":  FOREVER,
	"function": FUNCTION,
	"global":   GLOBAL,
	"gosub":    GOSUB,
	"goto":     GOTO,
	"if":       IF,
	"insert":   INSERT,
	"int":      INT,
	"last":     LAST,
	"local":    LOCAL,
	"mod":      MOD,
	"new":      NEW,
	"next":     NEXT,
	"not":      NOT,
	"null":     NULL,
	"or":       OR,
	"pi":       PI,
	"read":     READ,
	"repeat":   REPEAT,
	"restore":  RESTORE,
	"return":   RETURN,
	"sar":      SAR,
	"select":   SELECT,
	"shl":      SHL,
	"shr":      SHR,
	"step":     STEP,
	"str":      STR,
	"then":     THEN,
	"to":       TO,
	"true":     TRUE,
	"type":     TYPE,
	"until":    UNTIL,
	"wend":     WEND,
	"while":    WHILE,
	"xor":      XOR,
	"include":  INCLUDE,
}

type Token struct {
	Kind  TokenKind
	Value string
}

func NewToken(kind TokenKind, value string) Token {
	return Token{
		kind,
		value,
	}
}

func (token Token) IsOneOfMany(expectedKinds ...TokenKind) bool {
	for _, expected := range expectedKinds {
		if token.Kind == expected {
			return true
		}
	}

	return false
}

func (token Token) Print() {
	if token.IsOneOfMany(INT_LIT, FLT_LIT, STR_LIT, IDENT, COMMENT, HEX_INT) {
		fmt.Printf(
			"Type: %s, Value: %s\n",
			TokenKindString(token.Kind),
			token.Value,
		)
	} else {
		fmt.Printf("Type: %s\n", TokenKindString(token.Kind))
	}
}

func TokenKindString(kind TokenKind) string {
	switch kind {
	case EOF:
		return "EOF"
	case COMMENT:
		return "COMMENT"
	case COMMA:
		return "COMMA"
	case FLT_LIT:
		return "FLT_LIT"
	case STR_LIT:
		return "STR_LIT"
	case INT_LIT:
		return "INT_LIT"
	case IDENT:
		return "IDENT"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case STAR:
		return "STAR"
	case SLASH:
		return "DIV"
	case NOT:
		return "NOT"
	case MOD:
		return "MOD"
	case SAR:
		return "SAR"
	case SHL:
		return "SHL"
	case SHR:
		return "SHR"
	case EQ:
		return "EQ"
	case LT:
		return "LT"
	case GT:
		return "GT"
	case GE:
		return "GE"
	case LE:
		return "LE"
	case NE:
		return "NE"
	case OPEN_PR:
		return "OPEN_PR"
	case CLOSE_PR:
		return "CLOSE_PR"
	case TYPE:
		return "TYPE"
	case FIELD:
		return "FIELD"
	case AFTER:
		return "AFTER"
	case BEFORE:
		return "BEFORE"
	case EACH:
		return "EACH"
	case FIRST:
		return "FIRST"
	case DELETE:
		return "DELETE"
	case LAST:
		return "LAST"
	case INSERT:
		return "INSERT"
	case NEXT:
		return "NEXT"
	case AND:
		return "AND"
	case OR:
		return "OR"
	case XOR:
		return "XOR"
	case SELECT:
		return "SELECT"
	case CASE:
		return "CASE"
	case DEFAULT:
		return "DEFAULT"
	case CONST:
		return "CONST"
	case LOCAL:
		return "LOCAL"
	case GLOBAL:
		return "GLOBAL"
	case DIM:
		return "DIM"
	case DATA:
		return "DATA"
	case RESTORE:
		return "RESTORE"
	case READ:
		return "READ"
	case GOSUB:
		return "GOSUB"
	case GOTO:
		return "GOTO"
	case IF:
		return "IF"
	case ELSE:
		return "ELSE"
	case ELSEIF:
		return "ELSEIF"
	case ENDIF:
		return "ENDIF"
	case THEN:
		return "THEN"
	case FOR:
		return "FOR"
	case WHILE:
		return "WHILE"
	case FOREVER:
		return "FOREVER"
	case DO:
		return "DO"
	case REPEAT:
		return "REPEAT"
	case EXIT:
		return "EXIT"
	case WEND:
		return "WEND"
	case TO:
		return "TO"
	case STEP:
		return "STEP"
	case UNTIL:
		return "UNTIL"
	case INT:
		return "INT"
	case STR:
		return "STR"
	case FLOAT:
		return "FLOAT"
	case NEW:
		return "NEW"
	case NULL:
		return "NULL"
	case PI:
		return "PI"
	case FALSE:
		return "FALSE"
	case TRUE:
		return "TRUE"
	case FUNCTION:
		return "FUNCTION"
	case RETURN:
		return "RETURN"
	case INCLUDE:
		return "INCLUDE"
	case END:
		return "END"
	case NL:
		return "NewLine"
	case OPEN_BR:
		return "["
	case CLOSE_BR:
		return "]"
	case INT_SIG:
		return "IntSig"
	case FLT_SIG:
		return "FltSig"
	case STR_SIG:
		return "StrSig"
	case HEX_INT:
		return "Hex"
	case COLON:
		return ":"
	case DOT:
		return "."
	case BSLASH:
		return "\\"
	case EXP:
		return "^"
	default:
		return fmt.Sprintf("Error, token %d not implemented\n", kind)
	}
}
