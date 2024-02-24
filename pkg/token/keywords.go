package token

import "strings"

var keywords = map[string]Type{
    "every": EVERY,
    "in": IN,
    "next": NEXT,
    "other": OTHER,
    "on": ON,
    "of": OF,
    "the": THE,

    "day": DAY,
    "days": DAY,
    "week": WEEK,
    "weeks": WEEK,
    "month": MONTH,
    "months": MONTH,
    "year": YEAR,
    "years": YEAR,

    "monday": MONDAY,
    "mon": MONDAY,
    "tuesday": TUESDAY,
    "tue": TUESDAY,
    "tues": TUESDAY,
    "wednesday": WEDNESDAY,
    "wed": WEDNESDAY,
    "weds": WEDNESDAY,
    "thursday": THURSDAY,
    "thu": THURSDAY,
    "thur": THURSDAY,
    "thurs": THURSDAY,
    "friday": FRIDAY,
    "fri": FRIDAY,
    "saturday": SATURDAY,
    "sat": SATURDAY,
    "sunday": SUNDAY,
    "sun": SUNDAY,

    "january": JANUARY,
    "jan": JANUARY,
    "february": FEBRUARY,
    "feb": FEBRUARY,
    "march": MARCH,
    "mar": MARCH,
    "april": APRIL,
    "apr": APRIL,
    "may": MAY,
    "june": JUNE,
    "jun": JUNE,
    "july": JULY,
    "jul": JULY,
    "august": AUGUST,
    "aug": AUGUST,
    "september": SEPTEMBER,
    "sep": SEPTEMBER,
    "sept": SEPTEMBER,
    "october": OCTOBER,
    "oct": OCTOBER,
    "november": NOVEMBER,
    "nov": NOVEMBER,
    "december": DECEMBER,
    "dec": DECEMBER,

    "th": TH,
    "st": ST,
    "rd": RD,
    "nd": ND,
}

func LookupKeyword(word string) Type {
    if tok, ok := keywords[strings.ToLower(word)]; ok {
        return tok
    }
    return ILLEGAL
}
    
