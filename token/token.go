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
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	PERCENT  = "%"
	GT       = ">"
	GTE      = ">="
	LT       = "<"
	LTE      = "<="
	AND      = "&"
	BAND     = "&&"
	OR       = "|"
	BOR      = "||"
	EQ       = "=="
	NE       = "!="
	BANG     = "!"
	DPLUS    = "++"

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
	IF       = "IF"
	ELSE     = "ELSE"
	SWITCH   = "SWITCH"
	CASE     = "CASE"
	DEFAULT  = "DEFAULT"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	FOR      = "FOR"
	BREAK    = "BREAK"
)

var (
	keywords = map[string]TokenType{
		"fn":      FUNCTION,
		"let":     LET,
		"if":      IF,
		"else":    ELSE,
		"return":  RETURN,
		"switch":  SWITCH,
		"case":    CASE,
		"default": DEFAULT,
		"break":   BREAK,
		"true":    TRUE,
		"false":   FALSE,
		"for":     FOR,
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
