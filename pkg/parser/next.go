package parser

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/liamawhite/schedule/pkg/recurrence"
	"github.com/liamawhite/schedule/pkg/relative"
	"github.com/liamawhite/schedule/pkg/token"
)

func (p *Parser) parseNextStatement(relativeTo time.Time) (*recurrence.Recurrence, error) {
	// next <day/mon/week/month/jan/25/25th/year/weekend/weekday>
	// Some of these aren't something someone would actually use e.g. next 25, but they're valid syntax.

	// Consume the "next" token
	p.nextToken()

	newDate, err := handleRelative(relativeTo, p.curToken)
	if err != nil {
		return nil, fmt.Errorf("expected next to be followed by day, day name, week, month, month name, year, weekend, weekday or a number: %w", err)
	}
	return recurrence.NewRecurrence(recurrence.Option{Count: 1, Dtstart: newDate})
}

func handleRelative(relativeTo time.Time, tok token.Token) (time.Time, error) {
	switch tok.Type {
	case token.DAY: // next day
		return relativeTo.AddDate(0, 0, 1), nil
	case token.DAYOFWEEK: // next mon, tue, etc
		return relative.NextNamedDay(relativeTo, token.DayOfWeek(tok.Literal)), nil
	case token.WEEK: // next week (i.e. monday)
		return relative.NextWeek(relativeTo), nil
	case token.MONTH: // next month (i.e. 1st)
		return relative.NextMonth(relativeTo), nil
	case token.MONTHOFYEAR: // next jan, feb, etc
		return relative.NextNamedMonth(relativeTo, token.MonthOfYear(tok.Literal)), nil
	case token.NUMBER: // next 25, 25th
		num, err := strconv.Atoi(tok.Literal)
		if err != nil {
			return time.Time{}, fmt.Errorf("unable to parse number: %v", err)
		}
		return relative.NextSpecifiedMonthDay(relativeTo, num), nil
	case token.YEAR: // next Year
		return relative.NextYear(relativeTo), nil
	case token.SECTIONOFWEEK:
		if strings.ToLower(tok.Literal) == "weekend" {
			return relative.NextNamedDay(relativeTo, time.Saturday), nil
		}
		return relative.NextWeekday(relativeTo), nil
	}
	return time.Time{}, fmt.Errorf("unknown token: %v", tok.Literal)
}
