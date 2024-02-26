package parser

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/liamawhite/schedule/pkg/recurrence"
	"github.com/liamawhite/schedule/pkg/token"
	"github.com/teambition/rrule-go"
)

func (p *Parser) parseEveryStatement(relativeTo time.Time) (*recurrence.Recurrence, error) {
	// Consume the "every" curToken
	p.nextToken()

	// every day/week/month/year
	if p.currentTokenIsOneOf(token.DAY, token.WEEK, token.MONTH, token.YEAR) {
		freq, err := Lyify(p.curToken.Type)
		if err != nil {
			return nil, err
		}
		return recurrence.NewRecurrence(recurrence.Option{Dtstart: relativeTo, Freq: freq})
	}

	// every weekday/weekend
	if p.currentTokenIs(token.SECTIONOFWEEK) {
		if strings.ToLower(p.curToken.Literal) == "weekday" {
			return recurrence.NewRecurrence(recurrence.Option{
				Dtstart:   relativeTo,
				Freq:      rrule.WEEKLY,
				Byweekday: []rrule.Weekday{rrule.MO, rrule.TU, rrule.WE, rrule.TH, rrule.FR},
			})
		}
		return recurrence.NewRecurrence(recurrence.Option{Dtstart: relativeTo, Byweekday: []rrule.Weekday{rrule.SA}})
	}

	// every <day of week> or list of <day of week>
	if p.currentTokenIs(token.DAYOFWEEK) {
		var days []rrule.Weekday
		for !p.currentTokenIs(token.EOF) {
			day, err := DayOfWeek(p.curToken)
			if err != nil {
				return nil, err
			}
			days = append(days, day)
			p.nextToken()
		}
		return recurrence.NewRecurrence(recurrence.Option{
			Dtstart:   relativeTo,
			Freq:      rrule.WEEKLY,
			Byweekday: days,
		})
	}

	// every <number> <days/weeks/months/years>
	if p.currentTokenIs(token.NUMBER) && p.peekTokenIsOneOf(token.DAY, token.WEEK, token.MONTH, token.YEAR) {
		count, err := strconv.Atoi(p.curToken.Literal)
		if err != nil {
			return nil, fmt.Errorf("unable to convert %v to a number", p.curToken.Literal)
		}
		p.nextToken()
		freq, err := Lyify(p.curToken.Type)
		if err != nil {
			return nil, err
		}
		return recurrence.NewRecurrence(recurrence.Option{Dtstart: relativeTo, Freq: freq, Interval: count})
	}

	// every <day of month/month of year> or list of <day of month/month of year>
	// or any combination of the two
	if p.currentTokenIsOneOf(token.MONTHOFYEAR, token.NUMBER) {
		days, months := []int{}, []int{}
		for p.currentTokenIs(token.NUMBER) || p.currentTokenIs(token.MONTHOFYEAR) {
			if p.currentTokenIs(token.NUMBER) {
				num, err := strconv.Atoi(p.curToken.Literal)
				if err != nil {
					return nil, fmt.Errorf("unable to convert %v to a number", p.curToken.Literal)
				}
				days = append(days, num)
			}
			if p.currentTokenIs(token.MONTHOFYEAR) {
				months = append(months, int(token.MonthOfYear(p.curToken.Literal)))
			}
			p.nextToken()
		}
        // No days are specified use the first day of the months
        if len(days) == 0 {
            days = []int{1}
        }
		return recurrence.NewRecurrence(recurrence.Option{
			Dtstart:    relativeTo,
			Freq:       rrule.YEARLY,
			Bymonth:    months,
			Bymonthday: days,
		})
	}

	return nil, fmt.Errorf("unknown token at beginning: %v", p.curToken.Literal)
}

func DayOfWeek(tok token.Token) (rrule.Weekday, error) {
	switch token.DayOfWeek(tok.Literal) {
	case time.Monday:
		return rrule.MO, nil
	case time.Tuesday:
		return rrule.TU, nil
	case time.Wednesday:
		return rrule.WE, nil
	case time.Thursday:
		return rrule.TH, nil
	case time.Friday:
		return rrule.FR, nil
	case time.Saturday:
		return rrule.SA, nil
	case time.Sunday:
		return rrule.SU, nil
	}
	return rrule.MO, fmt.Errorf("unable to convert token to weekday: %v", tok.Literal)
}

func Lyify(tok token.Type) (rrule.Frequency, error) {
	switch tok {
	case token.DAY:
		return rrule.DAILY, nil
	case token.WEEK:
		return rrule.WEEKLY, nil
	case token.MONTH:
		return rrule.MONTHLY, nil
	case token.YEAR:
		return rrule.YEARLY, nil
	}
	return rrule.DAILY, fmt.Errorf("unable to convert token to frequency: %v", tok)
}
