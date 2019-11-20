package token

import (
	"unicode"

	"golang.org/x/text/unicode/rangetable"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//Identifiers + Literals
	IDENT = "IDENT" //add, foobar, x, y, ...
	INT   = "INT"   //123456789

	//Operators
	ASSIGN     = "="
	PLUS       = "+"
	MINUS      = "-"
	MULTIPLY   = "*"
	DEVID      = "/"
	PERCENTAGE = "%"
	GT         = ">"
	GTE        = ">="
	LT         = "<"
	LTE        = "<="
	AND        = "&"
	BAND       = "&&"
	OR         = "|"
	BOR        = "||"
	EQUAL      = "=="
	NEQUAL     = "!="
	BANG       = "!"

	//Delimiters
	COMMA     = ","
	SIMICOLON = ";"
	COLON     = ":"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var (
	keywords = map[string]TokenType{
		"fn":  FUNCTION,
		"let": LET,
		"<=":  LTE,
		">=":  GTE,
		"&&":  BAND,
		"||":  BOR,
		"==":  EQUAL,
		"!=":  NEQUAL,
	}
	//Allowed Identifier CharSet =
	AllowedIdentChars = []*unicode.RangeTable{unicode.Letter, unicode.Number, rangetable.New('.', '_')}
)

func Lookup(ident string) TokenType {
	if t, ok := keywords[ident]; ok {
		return t
	} else {
		return IDENT
	}
}
