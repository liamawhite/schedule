package parser

import (
	"fmt"
	"strconv"
	"time"

	"github.com/liamawhite/schedule/pkg/recurrence"
	"github.com/liamawhite/schedule/pkg/token"
)

const day = time.Hour * 24

// In statements are always of the form "in x days", "in x weeks", "in x months", "in x years"
func (p *Parser) parseInStatement(relativeTo time.Time) (*recurrence.Recurrence, error) {

	// Consume the "in" token
	p.nextToken()

	// Ensure the next token is a number
	if !p.currentTokenIs(token.NUMBER) {
		return nil, fmt.Errorf("expected a number after in but got %v", p.curToken.Literal)
	}

	// Ensure the token after that is a unit
	if !p.peekTokenIsOneOf(token.DAY, token.WEEK, token.MONTH, token.YEAR) {
		return nil, fmt.Errorf("expected days/weeks/months/years after in NUMBER but got %v", p.peekToken.Literal)
	}

	// Parse the two tokens
	count, err := strconv.Atoi(p.curToken.Literal)
	if err != nil {
		return nil, fmt.Errorf("unable to convert %v to a number", p.curToken.Literal)
	}
	unit := p.peekToken

	newDate := relativeTo

	switch unit.Type {
	case token.DAY:
		newDate = newDate.AddDate(0, 0, count)
	case token.WEEK:
		newDate = newDate.AddDate(0, 0, count*7)
	case token.MONTH:
		newDate = newDate.AddDate(0, count, 0)
	case token.YEAR:
		newDate = newDate.AddDate(count, 0, 0)
	}

    return recurrence.NewRecurrence(recurrence.Option{
        Count:   1,
        Dtstart: newDate,
    })
}
