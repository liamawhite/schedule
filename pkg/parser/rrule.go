package parser

import (
	"fmt"

	"github.com/liamawhite/schedule/pkg/token"
	"github.com/teambition/rrule-go"
)


func TokenToFrequency(tok token.Token) (rrule.Frequency, error) {
    switch tok.Type {
    case token.DAILY:
        return rrule.DAILY, nil
    case token.WEEKLY:
        return rrule.WEEKLY, nil
    case token.MONTHLY:
        return rrule.MONTHLY, nil
    case token.YEARLY:
        return rrule.YEARLY, nil
    default:
        return -1, fmt.Errorf("invalid frequency token: %v", tok.Literal)
    }
}
