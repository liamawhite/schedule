package parser

import (
	"fmt"
	"time"

	"github.com/liamawhite/schedule/pkg/date"
	"github.com/liamawhite/schedule/pkg/lexer"
	"github.com/liamawhite/schedule/pkg/recurrence"
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

	// If we've got an ordinal suffix, just skip it.
	if p.currentTokenIsOneOf(token.ST, token.ND, token.RD, token.TH) {
		p.nextToken()
	}
	// If we've got a delimiter, just skip it.
	if p.currentTokenIsOneOf(token.COMMA, token.SLASH) {
		p.nextToken()
	}
}

// Parse date uses the same underlying logic as Parse, but it returns a Date object instead of a Recurrence object.
// It does this by building a recurrence object with the parsed date and then returning the date from that object.
func (p *Parser) ParseDate(relativeTo time.Time) (*date.Date, error) {
	rule, err := p.Parse(relativeTo)
	if err != nil {
		return nil, err
	}
	instances := rule.All()
	if len(instances) != 1 {
		return nil, fmt.Errorf("unexpected issue processing date")
	}
	return instances[0], nil
}

// This parser returns recurrences. If the input represents a date, then there its a reccurence
// where the first and only repeat is the date described.
func (p *Parser) Parse(relativeTo time.Time) (*recurrence.Recurrence, error) {
	// Normalize the time to midnight and UTC.
	relativeTo = time.Date(relativeTo.Year(), relativeTo.Month(), relativeTo.Day(), 0, 0, 0, 0, time.UTC)

	switch {
	case p.currentTokenIs(token.IN):
		return p.parseInStatement(relativeTo)
	case p.currentTokenIs(token.NEXT):
		return p.parseNextStatement(relativeTo)
	case p.currentTokenIs(token.EVERY):
		return p.parseEveryStatement(relativeTo)
	case p.currentTokenIsOneOf(token.DAY, token.DAYOFWEEK, token.WEEK, token.MONTH, token.MONTHOFYEAR, token.YEAR, token.SECTIONOFWEEK, token.NUMBER, token.DAILY, token.WEEKLY, token.MONTHLY, token.YEARLY):
		return p.parseExpressions(relativeTo)

	}
	return nil, fmt.Errorf("Unknown token at beginning: %v", p.curToken.Literal)
}

func (p Parser) currentTokenIs(t token.Type) bool {
	return p.curToken.Type == t
}

func (p Parser) currentTokenIsOneOf(types ...token.Type) bool {
	for _, t := range types {
		if p.currentTokenIs(t) {
			return true
		}
	}
	return false
}

func (p Parser) peekTokenIs(t token.Type) bool {
	return p.peekToken.Type == t
}

func (p Parser) peekTokenIsOneOf(types ...token.Type) bool {
	for _, t := range types {
		if p.peekTokenIs(t) {
			return true
		}
	}
	return false
}
