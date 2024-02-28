package parser

import (
	"strconv"
	"time"

	"github.com/liamawhite/schedule/pkg/recurrence"
	"github.com/liamawhite/schedule/pkg/token"
)

// Expressions are anything that doesn't start with a keyword like, next, in or every.
// Same behavior as next: name of day, weekend, weekday, month, year, week
// Absolute dates: 1st jan, 1 jan, jan, 1, 1st, 2024, 1st jan 2024, jan 1, jan 1 2024, jan 2024, jan 1st, jan 1st 2024
func (p *Parser) parseExpressions(current time.Time) (*recurrence.Recurrence, error) {

	// Handle things that are unambiguously the same as next without the next keyword
	if p.currentTokenIsOneOf(token.DAY, token.DAYOFWEEK, token.WEEK, token.MONTH, token.YEAR, token.SECTIONOFWEEK) {
		t, err := handleRelative(current, p.curToken)
		if err != nil {
			return nil, err
		}

		return recurrence.NewRecurrence(recurrence.Option{
			Count:   1,
			Dtstart: t,
		})
	}

	// Handle daily, weekly, monthly, yearly
	if p.currentTokenIsOneOf(token.DAILY, token.WEEKLY, token.MONTHLY, token.YEARLY) {
		freq, err := TokenToFrequency(p.curToken)
		if err != nil {
			return nil, err
		}
		return recurrence.NewRecurrence(recurrence.Option{Dtstart: current, Freq: freq})
	}

	// We should have between 1 and 3 tokens left, and we need to figure out what they are.
	// Each of them could be a month day, a month, or a Year.
	day, month, year := -1, time.Month(-1), -1
	for !p.currentTokenIs(token.EOF) {
		if p.currentTokenIs(token.NUMBER) {
			d, y, err := guessDayOrYear(p.curToken)
			if err != nil {
				return nil, err
			}
			if d != -1 {
				day = d
			}
			if y != -1 {
				year = y
			}
		} else if p.currentTokenIs(token.MONTHOFYEAR) {
			month = token.MonthOfYear(p.curToken.Literal)
		}
		p.nextToken()
	}

	// If any of the values are -1, we need to figure out what they are.
	// Day is always the first day of the month if it's not specified.
	if day == -1 {
		day = 1
	}
	// If we have year but not month we default to january.
	if month == -1 && year != -1 {
		return recurrence.NewRecurrence(recurrence.Option{
			Count:   1,
			Dtstart: time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC),
		})
	}
	// If we only have the month, we default to the first day of the next occurence of that month.
	if month != -1 && year == -1 {
		// Check if the date is in the past, if so, we need to add a year.
		proposedDate := time.Date(current.Year(), month, day, 0, 0, 0, 0, time.UTC)
		if !proposedDate.After(current) { // use !after to include exact matches
			proposedDate = proposedDate.AddDate(1, 0, 0)
		}
		return recurrence.NewRecurrence(recurrence.Option{Count: 1, Dtstart: proposedDate})
	}
	// If we have neither month nor year, we default to the next occurence of the day.
	if month == -1 && year == -1 {
		// Check if the date is in the past, if so, we need to add a month.
		proposedDate := time.Date(current.Year(), current.Month(), day, 0, 0, 0, 0, time.UTC)
		if !proposedDate.After(current) { // use !after to include exact matches
			proposedDate = proposedDate.AddDate(0, 1, 0)
		}
		return recurrence.NewRecurrence(recurrence.Option{Count: 1, Dtstart: proposedDate})
	}

	// If we've got all three, we're good to go.
	return recurrence.NewRecurrence(recurrence.Option{Count: 1, Dtstart: time.Date(year, month, day, 0, 0, 0, 0, time.UTC)})
}

func guessDayOrYear(cur token.Token) (day, year int, err error) {
	num, err := strconv.Atoi(cur.Literal)
	if err != nil {
		return -1, -1, err
	}
	if num > 31 {
		return -1, num, nil
	}
	return num, -1, nil
}
