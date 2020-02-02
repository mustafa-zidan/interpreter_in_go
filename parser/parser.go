package parser

import (
	"moose/ast"
	"moose/lexer"
	"moose/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  *token.Token
	peekToken *token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	//Read two tokens, so currenToken and peekToken are poth set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.Next()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
