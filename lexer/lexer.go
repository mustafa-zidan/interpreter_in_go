package lexer

import (
	"bufio"
	"io"
	"log"
	"os"
	"unicode"

	"github.com/mustafa-zidan/interpreter_in_go/token"
)

type Lexer struct {
	reader   *bufio.Reader // buffer of strings
	file     string        // current file name
	line     int           // Line number
	position int           // current position in the input
	char     rune
}

func New(file string, reader io.Reader) *Lexer {
	l := &Lexer{reader: bufio.NewReader(reader), file: file}
	l.readChar()
	return l
}

// Next get next Token
func (l *Lexer) Next() *token.Token {
	var t token.Token

	l.eatWhiteSpaces()

	switch l.char {
	case '!':
		if l.peekChar() == '=' {
			t = newToken(token.NE, token.NE)
			l.readChar()
		} else {
			t = newToken(token.BANG, token.BANG)
		}
	case '=':
		if l.peekChar() == '=' {
			t = newToken(token.EQ, token.EQ)
			l.readChar()
		} else {
			t = newToken(token.ASSIGN, token.ASSIGN)
		}
	case '>':
		if l.peekChar() == '=' {
			t = newToken(token.GTE, token.GTE)
			l.readChar()
		} else {
			t = newToken(token.GT, token.GT)
		}
	case '<':
		if l.peekChar() == '=' {
			t = newToken(token.LTE, token.LTE)
			l.readChar()
		} else {
			t = newToken(token.LT, token.LT)
		}
	case '&':
		if l.peekChar() == '&' {
			t = newToken(token.BAND, token.BAND)
			l.readChar()
		} else {
			t = newToken(token.AND, token.AND)
		}
	case '|':
		if l.peekChar() == '&' {
			t = newToken(token.BOR, token.BOR)
			l.readChar()
		} else {
			t = newToken(token.OR, token.OR)
		}
	case ';':
		t = newToken(token.SIMICOLON, token.SIMICOLON)
	case ':':
		t = newToken(token.COLON, token.COLON)
	case '(':
		t = newToken(token.LPAREN, token.LPAREN)
	case ')':
		t = newToken(token.RPAREN, token.RPAREN)
	case ',':
		t = newToken(token.COMMA, token.COMMA)
	case '+':
		if l.peekChar() == '+' {
			t = newToken(token.DPLUS, token.DPLUS)
			l.readChar()
		} else {
			t = newToken(token.PLUS, token.PLUS)
		}
	case '-':
		t = newToken(token.MINUS, token.MINUS)
	case '*':
		t = newToken(token.ASTERISK, token.ASTERISK)
	case '/':
		t = newToken(token.SLASH, token.SLASH)
	case '%':
		t = newToken(token.PERCENT, token.PERCENT)
	case '{':
		t = newToken(token.LBRACE, token.LBRACE)
	case '}':
		t = newToken(token.RBRACE, token.RBRACE)
	case rune(0):
		t = newToken(token.EOF, string(l.char))
	default:
		if unicode.IsLetter(l.char) || l.char == '_' {
			t.Literal = l.readIdentifier()
			t.Type = token.Lookup(t.Literal)
		} else if unicode.IsNumber(l.char) {
			t.Literal = l.readIdentifier()
			t.Type = token.INT
		} else {
			log.Fatalf("Error in %s line %d column: %d : Illegal charachter '%s'\n", l.file, l.line, l.position, string(l.char))
			t = newToken(token.ILLEGAL, string(l.char))
		}
	}
	l.readChar()
	return &t
}

func newToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}

func (l *Lexer) readIdentifier() string {
	r := make([]rune, 0)
	for unicode.IsOneOf(token.AllowedIdentChars, l.char) {
		r = append(r, l.char)
		l.readChar()
	}
	// reset the the last character
	err := l.reader.UnreadRune()
	if err != nil {
		log.Fatalf("Error in %s line %d column: %d : unknown error %v\n", l.file, l.line, l.position, err)
		os.Exit(1)
	}
	return string(r)
}

func (l *Lexer) eatWhiteSpaces() {
	for unicode.IsSpace(l.char) {
		l.readChar()
	}
}

// read current character and set the charachter position
func (l *Lexer) readChar() {
	var err error
	var n int
	l.char, n, err = l.reader.ReadRune()
	if err != nil {
		if err == io.EOF {
			l.char = rune(0)
		} else {
			log.Fatalf("Error in %s line %d column: %d :%v\n", l.file, l.line, l.position, err)
			os.Exit(1)
		}
	} else if l.char == '\n' {
		l.line++
		l.position = 0
		l.readChar()
	} else {
		l.position += n
	}
}

func (l *Lexer) peekChar() rune {
	char, _, err := l.reader.ReadRune()
	if err != nil {
		if err == io.EOF {
			char = rune(0)
		} else {
			log.Fatalf("Error in %s line %d column: %d :%v\n", l.file, l.line, l.position, err)
			os.Exit(1)
		}
	}
	err = l.reader.UnreadRune()
	if err != nil {
		log.Fatalf("Error in %s line %d column: %d : unknown error %v\n", l.file, l.line, l.position, err)
		os.Exit(1)
	}
	return char
}
