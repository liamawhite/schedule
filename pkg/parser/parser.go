package parser

import (
	"github.com/liamawhite/schedule/pkg/ast"
	"github.com/liamawhite/schedule/pkg/lexer"
	"github.com/liamawhite/schedule/pkg/token"
)

type Parser struct {
    l *lexer.Lexer

    curToken  token.Token
    peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
    p := &Parser{l: l}
    
    // Read two tokens, so curToken and peekToken are both set
    p.nextToken()
    p.nextToken()

    return p
}

func (p *Parser) nextToken() {
    p.curToken = p.peekToken
    p.peekToken = p.l.NextToken()
}

func (p *Parser) Parse() ast.Statement {
    return nil
}
