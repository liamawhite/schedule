package token

type Type string

const (
    // Special tokens
    ILLEGAL = Type("illegal")
    EOF = Type("eof")

    // Literals
    INT = Type("int")

    // Delimiters
    COMMA = Type(",")
    SLASH = Type("/")

    // Keywords
    EVERY = Type("every")
    IN = Type("in")
    NEXT = Type("next")
    OTHER = Type("other")
    ON = Type("on")
    OF = Type("of")
    THE = Type("the")

    // Time units
    DAY = Type("day")
    WEEK = Type("week")
    MONTH = Type("month")
    YEAR = Type("year")

    // Time Literals
    MONDAY = Type("monday")
    TUESDAY = Type("tuesday")
    WEDNESDAY = Type("wednesday")
    THURSDAY = Type("thursday")
    FRIDAY = Type("friday")
    SATURDAY = Type("saturday")
    SUNDAY = Type("sunday")
    JANUARY = Type("january")
    FEBRUARY = Type("february")
    MARCH = Type("march")
    APRIL = Type("april")
    MAY = Type("may")
    JUNE = Type("june")
    JULY = Type("july")
    AUGUST = Type("august")
    SEPTEMBER = Type("september")
    OCTOBER = Type("october")
    NOVEMBER = Type("november")
    DECEMBER = Type("december")

    // Ordinal suffixes
    // They don't actually matter, but we need to account for them.
    TH = Type("th")
    ST = Type("st")
    RD = Type("rd")
    ND = Type("nd")
)
