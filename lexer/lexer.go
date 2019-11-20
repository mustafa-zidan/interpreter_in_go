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
	case '=':
		t = newToken(token.ASSIGN, l.char)
	case ';':
		t = newToken(token.SIMICOLON, l.char)
	case '(':
		t = newToken(token.LPAREN, l.char)
	case ')':
		t = newToken(token.RPAREN, l.char)
	case ',':
		t = newToken(token.COMMA, l.char)
	case '+':
		t = newToken(token.PLUS, l.char)
	case '{':
		t = newToken(token.LBRACE, l.char)
	case '}':
		t = newToken(token.RBRACE, l.char)
	case rune(0):
		t = newToken(token.EOF, l.char)
	default:
		if unicode.IsLetter(l.char) || l.char == '_' {
			t.Literal = l.readIdentifier()
			t.Type = token.Lookup(t.Literal)
		} else if unicode.IsNumber(l.char) {
			t.Literal = l.readIdentifier()
			t.Type = token.INT
		} else {
			log.Fatalf("Error in %s line %d column: %d : Illegal charachter %v\n", l.file, l.line, l.position, l.char)
			t = newToken(token.ILLEGAL, l.char)
		}
	}
	l.readChar()
	return &t
}

func newToken(tokenType token.TokenType, char rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
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
