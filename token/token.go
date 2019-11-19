package token

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
	ASSIGN = "="
	PLUS   = "-"

	//Delimiters
	COMMA     = ","
	SIMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func Lookup(ident string) TokenType {
	if t, ok := keywords[ident]; ok {
		return t
	} else {
		return IDENT
	}
}
