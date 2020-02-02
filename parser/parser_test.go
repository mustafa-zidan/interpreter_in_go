package parser

import (
	"strings"
	"testing"

	"moose/ast"
	"moose/lexer"

	"github.com/stretchr/testify/assert"
)

func TestLetStatment(t *testing.T) {
	input := strings.NewReader(`
	let x = 5;

	let y = 10;

	let foobar = 838383;

	`)

	l := lexer.New("test_parser.moose", input)

	p := New(l)

	program := p.ParseProgram()

	assert.NotNil(t, program)

	assert.Equal(t, 3, len(program.Statements))

	tests := []struct {
		expectedIdentifier string
	}{{"x"}, {"y"}, {"foobar"}}

	for i, test := range tests {
		statement := program.Statements[i]
		assert.IsType(t, *ast.LetStatement{}, statement)
		ls := statement.(*ast.LetStatement)
		assert.Equal(t, "let", ls.TokenLiteral())
		assert.Equal(t, test.expectedIdentifier, ls.Name.Value)
		assert.Equal(t, test.expectedIdentifier, ls.Name.TokenLiteral())
	}

}
