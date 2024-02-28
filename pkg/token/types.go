package token

type Type string

const (
    // Special tokens
    ILLEGAL = Type("illegal")
    EOF = Type("eof")

    // Literals
    NUMBER = Type("number")

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

    // Time Literal
    DAYOFWEEK = Type("dayofweek")
    MONTHOFYEAR = Type("monthofyear")
    SECTIONOFWEEK = Type("sectionofweek")

    // Recurring time units
    DAILY = Type("daily")
    WEEKLY = Type("weekly")
    MONTHLY = Type("monthly")
    YEARLY = Type("yearly")

    // Ordinal suffixes
    ST = Type("st")
    ND = Type("nd")
    RD = Type("rd")
    TH = Type("th")
)
