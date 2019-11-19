package lexer

import (
	"strings"
	"testing"

	"github.com/mustafa-zidan/interpreter_in_go/token"
	"github.com/stretchr/testify/assert"
)

func TestNextOK(t *testing.T) {
	input := strings.NewReader(`let five = 5;
		let ten = 10

		let add = fn(x, y) {
			x + y;
		}

		let result = add(five, ten)
		`)
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SIMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SIMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "func"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SIMICOLON, ";"},
		{token.RPAREN, "}"},

		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
	}
	l := New("test.moose", input)
	for _, tt := range tests {
		tok := l.Next()
		assert.Equal(t, tok.Type, tt.expectedType)
		assert.Equal(t, tok.Literal, tt.expectedLiteral)
	}
}

func TestNextError(t *testing.T) {
	//TODO test invalid line error
}
