package lexer

import (
	"strings"
	"testing"

	"github.com/mustafa-zidan/interpreter_in_go/token"
	"github.com/stretchr/testify/assert"
)

func TestNextOK(t *testing.T) {
	input := strings.NewReader(`let five = 5;
		let ten = 10;

		let add = fn(x, y) {
			x + y;
		};

		let result = add(five, ten);

		!-/*5

		if ( 5 > 10) {
			return true;
		} else {
			return false;
		}

		10 == 10;
		9 != 10;

		for( i=0; i <= five; i++ ) {
			print(ten);
		}

		switch five {
		  case 10:
            return ten;
		  case 5:
            return five;
			default:
				break;
		}
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
		{token.FUNCTION, "fn"},
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
		{token.RBRACE, "}"},
		{token.SIMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SIMICOLON, ";"},

		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},

		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.GT, ">"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SIMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SIMICOLON, ";"},
		{token.RBRACE, "}"},

		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SIMICOLON, ";"},

		{token.INT, "9"},
		{token.NE, "!="},
		{token.INT, "10"},
		{token.SIMICOLON, ";"},

		{token.FOR, "for"},
		{token.LPAREN, "("},
		{token.IDENT, "i"},
		{token.ASSIGN, "="},
		{token.INT, "0"},
		{token.SIMICOLON, ";"},
		{token.IDENT, "i"},
		{token.LTE, "<="},
		{token.IDENT, "five"},
		{token.SIMICOLON, ";"},
		{token.IDENT, "i"},
		{token.DPLUS, "++"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "print"},
		{token.LPAREN, "("},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SIMICOLON, ";"},
		{token.RBRACE, "}"},

		{token.SWITCH, "switch"},
		{token.IDENT, "five"},
		{token.LBRACE, "{"},
		{token.CASE, "case"},
		{token.INT, "10"},
		{token.COLON, ":"},
		{token.RETURN, "return"},
		{token.IDENT, "ten"},
		{token.SIMICOLON, ";"},
		{token.CASE, "case"},
		{token.INT, "5"},
		{token.COLON, ":"},
		{token.RETURN, "return"},
		{token.IDENT, "five"},
		{token.SIMICOLON, ";"},
		{token.DEFAULT, "default"},
		{token.COLON, ":"},
		{token.BREAK, "break"},
		{token.SIMICOLON, ";"},
		{token.RBRACE, "}"},

		{token.EOF, "\x00"},
	}
	l := New("test.moose", input)
	for _, tt := range tests {
		tok := l.Next()
		assert.Equal(t, tt.expectedType, tok.Type)
		assert.Equal(t, tt.expectedLiteral, tok.Literal)
	}
}

func TestNextError(t *testing.T) {
	//TODO test invalid line error
}
