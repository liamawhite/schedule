package token

import (
	"strings"
	"time"
)

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

    "monday": DAYOFWEEK,
    "mon": DAYOFWEEK,
    "tuesday": DAYOFWEEK,
    "tue": DAYOFWEEK,
    "tues": DAYOFWEEK,
    "wednesday": DAYOFWEEK,
    "wed": DAYOFWEEK,
    "weds": DAYOFWEEK,
    "thursday": DAYOFWEEK,
    "thu": DAYOFWEEK,
    "thur": DAYOFWEEK,
    "thurs": DAYOFWEEK,
    "friday": DAYOFWEEK,
    "fri": DAYOFWEEK,
    "saturday": DAYOFWEEK,
    "sat": DAYOFWEEK,
    "sunday": DAYOFWEEK,
    "sun": DAYOFWEEK,

    "january": MONTHOFYEAR,
    "jan": MONTHOFYEAR,
    "february": MONTHOFYEAR,
    "feb": MONTHOFYEAR,
    "march": MONTHOFYEAR,
    "mar": MONTHOFYEAR,
    "april": MONTHOFYEAR,
    "apr": MONTHOFYEAR,
    "may": MONTHOFYEAR,
    "june": MONTHOFYEAR,
    "jun": MONTHOFYEAR,
    "july": MONTHOFYEAR,
    "jul": MONTHOFYEAR,
    "august": MONTHOFYEAR,
    "aug": MONTHOFYEAR,
    "september": MONTHOFYEAR,
    "sep": MONTHOFYEAR,
    "sept": MONTHOFYEAR,
    "october": MONTHOFYEAR,
    "oct": MONTHOFYEAR,
    "november": MONTHOFYEAR,
    "nov": MONTHOFYEAR,
    "december": MONTHOFYEAR,
    "dec": MONTHOFYEAR,

    "weekday": SECTIONOFWEEK,
    "weekend": SECTIONOFWEEK,

    "daily": DAILY,
    "weekly": WEEKLY,
    "monthly": MONTHLY,
    "yearly": YEARLY,

    "st": ST,
    "nd": ND,
    "rd": RD,
    "th": TH,
}

func LookupKeyword(word string) Type {
    if tok, ok := keywords[strings.ToLower(word)]; ok {
        return tok
    }
    return ILLEGAL
}

func DayOfWeek(literal string) time.Weekday {
    switch strings.ToLower(literal) {
    case "monday", "mon":
        return time.Monday
    case "tuesday", "tue", "tues":
        return time.Tuesday
    case "wednesday", "wed", "weds":
        return time.Wednesday
    case "thursday", "thu", "thur", "thurs":
        return time.Thursday
    case "friday", "fri":
        return time.Friday
    case "saturday", "sat":
        return time.Saturday
    case "sunday", "sun":
        return time.Sunday
    }
    return time.Sunday
}

func MonthOfYear(literal string) time.Month {
    switch strings.ToLower(literal) {
    case "january", "jan":
        return time.January
    case "february", "feb":
        return time.February
    case "march", "mar":
        return time.March
    case "april", "apr":
        return time.April
    case "may":
        return time.May
    case "june", "jun":
        return time.June
    case "july", "jul":
        return time.July
    case "august", "aug":
        return time.August
    case "september", "sep", "sept":
        return time.September
    case "october", "oct":
        return time.October
    case "november", "nov":
        return time.November
    case "december", "dec":
        return time.December
    }
    return time.January
}
    
