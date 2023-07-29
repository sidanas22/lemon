package parser

import (
	"lemon/lemon/ast"
	"lemon/lemon/lexer"
	"lemon/lemon/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}
